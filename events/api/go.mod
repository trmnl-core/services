module github.com/micro/services/events/api

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.5.1-0.20200428112352-414b2ec5f87a
	github.com/micro/services/events/service v0.0.0-00010101000000-000000000000
	github.com/micro/services/project/service v0.0.0-20200427143115-d065db30e6e8
	google.golang.org/appengine v1.6.1
)

replace github.com/micro/services/events/service => ../service
