module github.com/micro/services/portfolio/feed-api

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/feed-items v1.0.0
	github.com/micro/services/portfolio/helpers/authentication v1.0.0
	github.com/micro/services/portfolio/helpers/photos v1.0.0
	github.com/micro/services/portfolio/helpers/textenhancer v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/post-enhancer v1.0.0
	github.com/micro/services/portfolio/posts v1.0.0
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.2.0
)

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/comments => ../comments

replace github.com/micro/services/portfolio/bullbear => ../bullbear

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/post-enhancer => ../post-enhancer

replace github.com/micro/services/portfolio/feed-items => ../feed-items

replace github.com/micro/services/portfolio/posts => ../posts

replace github.com/micro/services/portfolio/helpers/unique => ../helpers/unique

replace github.com/micro/services/portfolio/helpers/authentication => ../helpers/authentication

replace github.com/micro/services/portfolio/helpers/textenhancer => ../helpers/textenhancer

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/microtime => ../helpers/microtime

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
