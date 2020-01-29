module github.com/micro/services/portfolio/investors-api

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/followers v1.0.0
	github.com/micro/services/portfolio/helpers/authentication v1.0.0
	github.com/micro/services/portfolio/helpers/iex-cloud v1.0.0
	github.com/micro/services/portfolio/helpers/photos v1.0.0
	github.com/micro/services/portfolio/helpers/unique v1.0.0
	github.com/micro/services/portfolio/portfolio-allocation v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/portfolio-value-tracking v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/portfolios v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/post-enhancer v1.0.0
	github.com/micro/services/portfolio/posts v1.0.0
	github.com/micro/services/portfolio/stock-quote-v2 v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/services/portfolio/trades v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1

)

replace github.com/micro/services/portfolio/portfolio-allocation => ../portfolio-allocation

replace github.com/micro/services/portfolio/posts => ../posts

replace github.com/micro/services/portfolio/comments => ../comments

replace github.com/micro/services/portfolio/bullbear => ../bullbear

replace github.com/micro/services/portfolio/ledger => ../ledger

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/stock-target-price => ../stock-target-price

replace github.com/micro/services/portfolio/helpers/unique => ../helpers/unique

replace github.com/micro/services/portfolio/helpers/microtime => ../helpers/microtime

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/portfolios => ../portfolios

replace github.com/micro/services/portfolio/trades => ../trades

replace github.com/micro/services/portfolio/insights => ../insights

replace github.com/micro/services/portfolio/stock-earnings => ../stock-earnings

replace github.com/micro/services/portfolio/portfolio-valuation => ../portfolio-valuation

replace github.com/micro/services/portfolio/stock-quote-v2 => ../stock-quote-v2

replace github.com/micro/services/portfolio/stock-quote => ../stock-quote

replace github.com/micro/services/portfolio/portfolio-value-tracking => ../portfolio-value-tracking

replace github.com/micro/services/portfolio/post-enhancer => ../post-enhancer

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/helpers/authentication => ../helpers/authentication

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/worldtradingdata => ../helpers/worldtradingdata

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
