module github.com/micro/services/portfolio/notifications

go 1.12

require (
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/services/portfolio/comments v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/feed-items v1.0.0
	github.com/micro/services/portfolio/followers v1.0.0
	github.com/micro/services/portfolio/helpers/microgorm v1.0.0
	github.com/micro/services/portfolio/helpers/textenhancer v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/posts v1.0.0
	github.com/micro/services/portfolio/push-notifications v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.2.0
	github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	go.etcd.io/etcd v3.3.13+incompatible
)

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/micro/services/portfolio/helpers/sms => ../helpers/sms

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/textenhancer => ../helpers/textenhancer

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/followers => ../followers

replace github.com/micro/services/portfolio/push-notifications => ../push-notifications

replace github.com/micro/services/portfolio/posts => ../posts

replace github.com/micro/services/portfolio/comments => ../comments

replace github.com/micro/services/portfolio/feed-items => ../feed-items

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
