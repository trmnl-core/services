module github.com/micro/services/portfolio/posts-api

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/bullbear v1.0.0
	github.com/micro/services/portfolio/comments v1.0.0
	github.com/micro/services/portfolio/helpers/authentication v1.0.0
	github.com/micro/services/portfolio/helpers/photos v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/textenhancer v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/posts v1.0.0
	github.com/micro/services/portfolio/stocks v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.2.0
	github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d // indirect
)

replace github.com/micro/services/portfolio/comments => ../comments

replace github.com/micro/services/portfolio/bullbear => ../bullbear

replace github.com/micro/services/portfolio/helpers/textenhancer => ../helpers/textenhancer

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/posts => ../posts

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/helpers/authentication => ../helpers/authentication

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
