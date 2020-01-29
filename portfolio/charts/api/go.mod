module github.com/micro/services/portfolio/charts-api

go 1.12

require (
	contrib.go.opencensus.io/exporter/ocagent v0.5.1 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.12.4 // indirect
	github.com/Azure/azure-pipeline-go v0.2.2 // indirect
	github.com/Azure/azure-storage-blob-go v0.7.0 // indirect
	github.com/Azure/go-autorest v12.3.0+incompatible // indirect
	github.com/GoogleCloudPlatform/cloudsql-proxy v0.0.0-20190725230627-253d1edd4416 // indirect
	github.com/RoaringBitmap/roaring v0.4.18 // indirect
	github.com/anacrolix/tagflag v1.0.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/glycerine/go-unsnap-stream v0.0.0-20190730064659-98d31706395a // indirect
	github.com/glycerine/goconvey v0.0.0-20190410193231-58a59202ab31 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/micro/services/portfolio/helpers/iex-cloud v1.0.0
	github.com/micro/services/portfolio/portfolio-value-tracking v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/portfolios v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-plugins v1.5.1
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945 // indirect
	golang.org/x/exp v0.0.0-20190718202018-cfdd5522f6f6 // indirect
	golang.org/x/image v0.0.0-20190729225735-1bd0cf576493 // indirect
	golang.org/x/mobile v0.0.0-20190719004257-d2bd2a29d028 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.3.0 // indirect
	honnef.co/go/tools v0.0.1-2019.2.2 // indirect
	pack.ag/amqp v0.12.0 // indirect
)

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/portfolio-value-tracking => ../portfolio-value-tracking

replace github.com/micro/services/portfolio/portfolio-valuation => ../portfolio-valuation

replace github.com/micro/services/portfolio/portfolios => ../portfolios

replace github.com/micro/services/portfolio/trades => ../trades

replace github.com/micro/services/portfolio/ledger => ../ledger

replace github.com/micro/services/portfolio/stock-quote => ../stock-quote

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/microtime => ../helpers/microtime

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
