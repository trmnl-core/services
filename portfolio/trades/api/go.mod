module github.com/micro/services/portfolio/trades-api

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/followers v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/authentication v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/iex-cloud v1.0.0
	github.com/micro/services/portfolio/portfolios v1.0.0
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/services/portfolio/trades v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
)

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/trades => ../trades

replace github.com/micro/services/portfolio/portfolios => ../portfolios

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/ledger => ../ledger

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/microtime => ../helpers/microtime

replace github.com/micro/services/portfolio/helpers/authentication => ../helpers/authentication

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
