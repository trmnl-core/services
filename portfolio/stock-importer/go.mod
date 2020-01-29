module github.com/micro/services/portfolio/stock-importer

go 1.12

require (
	github.com/micro/services/portfolio/helpers/photos v1.0.0
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/marten-seemann/qtls v0.3.1 // indirect
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.1.1
)

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/helpers/photos => ../helpers/photos
