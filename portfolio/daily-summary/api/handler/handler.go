package handler

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	feed "github.com/micro/services/portfolio/feed-items/proto"
	followers "github.com/micro/services/portfolio/followers/proto"
	auth "github.com/micro/services/portfolio/helpers/authentication"
	"github.com/micro/services/portfolio/helpers/microtime"
	valuation "github.com/micro/services/portfolio/portfolio-value-tracking/proto"
	portfolios "github.com/micro/services/portfolio/portfolios/proto"
	posts "github.com/micro/services/portfolio/posts/proto"
	earnings "github.com/micro/services/portfolio/stock-earnings/proto"
	news "github.com/micro/services/portfolio/stock-news/proto"
	quotes "github.com/micro/services/portfolio/stock-quote-v2/proto"
	target "github.com/micro/services/portfolio/stock-target-price/proto"
	stocks "github.com/micro/services/portfolio/stocks/proto"
	trades "github.com/micro/services/portfolio/trades/proto"
	users "github.com/micro/services/portfolio/users/proto"
)

// New returns an instance of Handler
func New(auth auth.Authenticator, client client.Client) *Handler {
	return &Handler{
		auth:          auth,
		news:          news.NewStockNewsService("kytra-v1-stock-news:8080", client),
		feed:          feed.NewFeedItemsService("kytra-v1-feed-items:8080", client),
		posts:         posts.NewPostsService("kytra-v1-posts:8080", client),
		users:         users.NewUsersService("kytra-v1-users:8080", client),
		stocks:        stocks.NewStocksService("kytra-v1-stocks:8080", client),
		trades:        trades.NewTradesService("kytra-v1-trades:8080", client),
		quotes:        quotes.NewStockQuoteService("kytra-v2-stock-quote:8080", client),
		targets:       target.NewStockTargetPriceService("kytra-v1-stock-target-price:8080", client),
		earnings:      earnings.NewStockEarningsService("kytra-v1-stock-earnings:8080", client),
		valuation:     valuation.NewPortfolioValueTrackingService("kytra-v1-portfolio-value-tracking:8080", client),
		followers:     followers.NewFollowersService("kytra-v1-followers:8080", client),
		portfolios:    portfolios.NewPortfoliosService("kytra-v1-portfolios:8080", client),
		luckyDipCache: map[int64]string{},
	}
}

// Handler is an object can process RPC requests
type Handler struct {
	auth       auth.Authenticator
	news       news.StockNewsService
	feed       feed.FeedItemsService
	users      users.UsersService
	posts      posts.PostsService
	stocks     stocks.StocksService
	trades     trades.TradesService
	quotes     quotes.StockQuoteService
	targets    target.StockTargetPriceService
	earnings   earnings.StockEarningsService
	followers  followers.FollowersService
	valuation  valuation.PortfolioValueTrackingService
	portfolios portfolios.PortfoliosService

	luckyDipCache map[int64]string
}

func (h *Handler) summaryForDate(ctx context.Context, t time.Time) (string, error) {
	portfolio, err := h.fetchPortfolioPrice(ctx, t)
	if err != nil {
		return "", err
	}

	index, err := h.fetchIndexPrice(ctx, t)
	if err != nil {
		return "", err
	}

	diff := portfolio - index
	var format string
	if index > 0 && diff > 0 {
		format = "The markets had a good day %v, rising %v. In addition to this, your portfolio outperformed the market by %v 🎉"
	} else if index > 0 && diff < 0 {
		emoji := "😬"
		if diff < 0.5 {
			emoji = "😭"
		}
		format = "The markets had a bullish day %v, gaining %v. Your portfolio lagged behind the market by %v " + emoji
	} else if index < 0 && diff > 0 {
		format = "The markets took a bit of a beating %v, falling %v. However, your portfolio weathered the storm well and beat the market by %v 😄"
	} else {
		format = "The markets had a rough day %v, falling %v. Your portfolio was hit especially hard and fell by an additional %v 😬"
	}

	date := t.Truncate(time.Hour * 24)
	today := time.Now().Truncate(time.Hour * 24)
	yesterday := today.Add(time.Hour * -24)

	var dayStr string
	if date.Unix() >= today.Unix() {
		dayStr = "today"
	} else if date.Unix() >= yesterday.Unix() {
		dayStr = "yesterday"
	} else {
		dayStr = t.Weekday().String()
	}

	formatFloat32 := func(val float32) string {
		abs := math.Abs(float64(val))
		rounded := math.Round(abs*100) / 100

		if val > 0.1 {
			return fmt.Sprintf("%v%%", rounded)
		}
		return "less than 0.1%"
	}

	return fmt.Sprintf(format, dayStr, formatFloat32(index), formatFloat32(diff)), nil
}

// fetchIndexPrice gets the day % chg for the Nasdaq Composite Index
func (h *Handler) fetchIndexPrice(ctx context.Context, t time.Time) (float32, error) {
	t = t.Truncate(time.Hour * 24)

	ctx = microtime.ContextWithTime(ctx, t.Add(time.Hour*24))
	formattedErr := errors.InternalServerError("MISSING_INDEX_DATA", "No index data was found")

	quote, err := h.quotes.GetQuote(ctx, &quotes.Stock{Uuid: "^IXIC"})
	if err != nil || quote == nil {
		return 0, formattedErr
	}
	// Quote wasn't for the day requested (it most likely doesn't exist yet)
	ca := time.Unix(quote.CreatedAt, 0).Truncate(time.Hour * 24)
	if ca != t {
		return 0, formattedErr
	}

	return quote.GetPercentageChange(), nil
}

// fetchPortfolioPrice gets the % chg for the users portfolios
func (h *Handler) fetchPortfolioPrice(ctx context.Context, t time.Time) (float32, error) {
	user, err := h.auth.UserFromContext(ctx)
	if err != nil {
		return 0, err
	}

	portfolio, err := h.portfolios.Get(ctx, &portfolios.Portfolio{UserUuid: user.UUID})
	if err != nil {
		return 0, err
	}

	rsp, err := h.valuation.GetPriceMovement(ctx, &valuation.GetPriceMovementsRequest{
		PortfolioUuid: portfolio.Uuid, StartDate: t.Unix(), EndDate: t.Unix(),
	})
	if err != nil {
		return 0, err
	}

	return rsp.GetPriceMovement().GetPercentageChange(), nil
}

func (h *Handler) followingForResource(ctx context.Context, user *users.User, rType string) ([]string, error) {
	// Step 1. Get the resources the users are following
	fRsp, err := h.followers.Get(ctx, &followers.Resource{Uuid: user.Uuid, Type: "User"})
	if err != nil {
		return nil, err
	}

	// Step 2. Filter down to the requested type
	uuids := []string{}
	for _, r := range fRsp.GetFollowing() {
		if r.Type == rType {
			uuids = append(uuids, r.Uuid)
		}
	}

	return uuids, nil
}
