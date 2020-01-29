module github.com/micro/services/portfolio/push-notifications

go 1.12

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/services/portfolio/helpers/microgorm v1.0.0
	github.com/micro/services/portfolio/helpers/sms v0.0.0-00010101000000-000000000000
	github.com/lucas-clemente/quic-go v0.11.2 // indirect
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.1.1
	github.com/oliveroneill/exponent-server-sdk-golang v0.0.0-20190619221248-65582a10ed02
	github.com/pkg/errors v0.8.1
)

replace github.com/micro/services/portfolio/helpers/sms => ../helpers/sms

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud
