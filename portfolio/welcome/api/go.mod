module github.com/micro/services/portfolio/welcome-api

go 1.12

require (
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/dongri/phonenumber v0.0.0-20191029000444-38c6f1b163b5
	github.com/golang/groupcache v0.0.0-20191002201903-404acd9df4cc // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/services/portfolio/helpers/authentication v1.0.0
	github.com/micro/services/portfolio/helpers/microgorm v1.0.0
	github.com/micro/services/portfolio/helpers/photos v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/helpers/sms v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/ledger v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/portfolios v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/sms-verification v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/users v1.0.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/pkg/errors v0.8.1
	go.etcd.io/etcd v3.3.13+incompatible
)

replace github.com/micro/services/portfolio/helpers/authentication => ../helpers/authentication

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/microtime => ../helpers/microtime

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/helpers/sms => ../helpers/sms

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos

replace github.com/micro/services/portfolio/users => ../users

replace github.com/micro/services/portfolio/portfolios => ../portfolios

replace github.com/micro/services/portfolio/sms-verification => ../sms-verification

replace github.com/micro/services/portfolio/ledger => ../ledger

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
