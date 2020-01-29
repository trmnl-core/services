module github.com/micro/services/portfolio/notifications-api

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/helpers/authentication v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/iex-cloud v1.0.0
	github.com/micro/services/portfolio/helpers/photos v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/notifications v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/push-notifications v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.2.0
)

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/push-notifications => ../push-notifications

replace github.com/micro/services/portfolio/notifications => ../notifications

replace github.com/micro/services/portfolio/posts => ../posts

replace github.com/micro/services/portfolio/feed-items => ../feed-items

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/comments => ../comments

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/helpers/sms => ../helpers/sms

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/textenhancer => ../helpers/textenhancer

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/micro/services/portfolio/helpers/authentication => ../helpers/authentication

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos
