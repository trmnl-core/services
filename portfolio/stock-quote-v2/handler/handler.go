package handler

import (
	"github.com/micro/go-micro/client"
	iex "github.com/micro/services/portfolio/helpers/iex-cloud"
	"github.com/micro/services/portfolio/helpers/worldtradingdata"
	insights "github.com/micro/services/portfolio/insights/proto"
	storage "github.com/micro/services/portfolio/stock-quote/storage"
	stocks "github.com/micro/services/portfolio/stocks/proto"
	trades "github.com/micro/services/portfolio/trades/proto"
)

// New returns an instance of Handler
func New(wtd worldtradingdata.Service, iex iex.Service, db storage.Service, client client.Client) *Handler {
	return &Handler{
		db:       db,
		wtd:      wtd,
		iex:      iex,
		trades:   trades.NewTradesService("kytra-v1-trades:8080", client),
		stocks:   stocks.NewStocksService("kytra-v1-stocks:8080", client),
		insights: insights.NewInsightsService("kytra-v1-insights:8080", client),
	}
}

// Handler is an object can process RPC requests
type Handler struct {
	db       storage.Service
	iex      iex.Service
	wtd      worldtradingdata.Service
	trades   trades.TradesService
	stocks   stocks.StocksService
	insights insights.InsightsService
}
