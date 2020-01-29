module github.com/micro/services/portfolio/feed-generator

go 1.12

require (
	github.com/micro/services/portfolio/feed-items v1.0.0
	github.com/micro/services/portfolio/followers v1.0.0
	github.com/micro/services/portfolio/posts v1.0.0
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.2.0
)

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/posts => ../posts

replace github.com/micro/services/portfolio/feed-items => ../feed-items

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
