module github.com/micro/services/portfolio/notifications-mailer

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/helpers/iex-cloud v1.0.0
	github.com/micro/services/portfolio/helpers/mailer v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/photos v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/unique v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/notifications v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.2.0
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.0
)

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/notifications => ../notifications

replace github.com/micro/services/portfolio/comments => ../comments

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/posts => ../posts

replace github.com/micro/services/portfolio/feed-items => ../feed-items

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/helpers/unique => ../helpers/unique

replace github.com/micro/services/portfolio/helpers/mailer => ../helpers/mailer

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/micro/services/portfolio/helpers/textenhancer => ../helpers/textenhancer

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos
