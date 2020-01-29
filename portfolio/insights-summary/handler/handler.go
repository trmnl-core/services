package handler

import (
	"context"

	"github.com/micro/go-micro/client"
	followers "github.com/micro/services/portfolio/followers/proto"
	proto "github.com/micro/services/portfolio/insights-summary/proto"
	valuation "github.com/micro/services/portfolio/portfolio-valuation/proto"
	portfolios "github.com/micro/services/portfolio/portfolios/proto"
	earnings "github.com/micro/services/portfolio/stock-earnings/proto"
	quotes "github.com/micro/services/portfolio/stock-quote-v2/proto"
	target "github.com/micro/services/portfolio/stock-target-price/proto"
	stocks "github.com/micro/services/portfolio/stocks/proto"
	trades "github.com/micro/services/portfolio/trades/proto"
	users "github.com/micro/services/portfolio/users/proto"
)

// New returns an instance of Handler
func New(client client.Client) *Handler {
	return &Handler{
		users:      users.NewUsersService("kytra-v1-users:8080", client),
		stocks:     stocks.NewStocksService("kytra-v1-stocks:8080", client),
		trades:     trades.NewTradesService("kytra-v1-trades:8080", client),
		quotes:     quotes.NewStockQuoteService("kytra-v2-stock-quote:8080", client),
		targets:    target.NewStockTargetPriceService("kytra-v1-stock-target-price:8080", client),
		earnings:   earnings.NewStockEarningsService("kytra-v1-stock-earnings:8080", client),
		valuation:  valuation.NewPortfolioValuationService("kytra-v1-portfolio-valuation:8080", client),
		followers:  followers.NewFollowersService("kytra-v1-followers:8080", client),
		portfolios: portfolios.NewPortfoliosService("kytra-v1-portfolios:8080", client),
	}
}

// Handler is an object can process RPC requests
type Handler struct {
	users      users.UsersService
	stocks     stocks.StocksService
	trades     trades.TradesService
	quotes     quotes.StockQuoteService
	targets    target.StockTargetPriceService
	earnings   earnings.StockEarningsService
	valuation  valuation.PortfolioValuationService
	followers  followers.FollowersService
	portfolios portfolios.PortfoliosService
}

// List returns a summary for the insights on a given day
func (h *Handler) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) (err error) {
	data := summaryData{
		Handler:    h,
		Context:    ctx,
		UserUUID:   req.GetUserUuid(),
		StockUUIDs: req.GetAssetUuids(),
	}
	if err := data.Fetch(); err != nil {
		return err
	}

	rsp.Assets = make([]*proto.Asset, len(req.AssetUuids))
	for i, uuid := range req.AssetUuids {
		rsp.Assets[i] = &proto.Asset{
			Uuid: uuid, Type: req.AssetType, Summary: data.Stringify(uuid),
		}
	}

	return nil
}

// Get returns the summary for a single stock on a given day
func (h *Handler) Get(ctx context.Context, req *proto.GetRequest, rsp *proto.Asset) error {
	data := summaryData{
		Handler:    h,
		Context:    ctx,
		UserUUID:   req.GetUserUuid(),
		StockUUIDs: []string{req.GetAssetUuid()},
	}
	if err := data.Fetch(); err != nil {
		return err
	}

	*rsp = proto.Asset{
		Uuid:    req.GetAssetUuid(),
		Type:    req.AssetType,
		Summary: data.Stringify(req.GetAssetUuid()),
	}

	return nil
}

type summaryData struct {
	// Input
	Handler    *Handler
	Context    context.Context
	UserUUID   string
	StockUUIDs []string

	// Processing
	AllUsers            []*users.User
	AllPositions        []*trades.Position
	AllPortfolios       []*portfolios.Portfolio
	TotalPortfolioValue int64
	SectorAllocations   map[string]int64

	// Ready for use (grouped by stock UUIID)
	StkQuote             map[string]*quotes.Quote
	StkMetadata          map[string]*stocks.Stock
	StkHasEarning        map[string]bool
	StkPriceTarget       map[string]*target.Stock
	StkUsersWithPosition map[string][]*users.User
}
