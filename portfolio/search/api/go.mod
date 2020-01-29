module github.com/micro/services/portfolio/search-api

go 1.12

require (
	github.com/abbot/go-http-auth v0.4.1-0.20181019201920-860ed7f246ff
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/followers v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/authentication v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/iex-cloud v1.0.0
	github.com/micro/services/portfolio/helpers/photos v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-plugins v1.5.1
)

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/authentication => ../helpers/authentication

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos
