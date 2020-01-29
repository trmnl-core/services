module github.com/micro/services/portfolio/helpers/textenhancer

replace github.com/micro/services/portfolio/users => ../../users

replace github.com/micro/services/portfolio/stocks => ../../stocks

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../passwordhasher

replace github.com/micro/services/portfolio/helpers/microgorm => ../microgorm

go 1.12

require (
	github.com/micro/services/portfolio/stocks v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/users v0.0.0-00010101000000-000000000000
	github.com/marten-seemann/qtls v0.3.1 // indirect
)
