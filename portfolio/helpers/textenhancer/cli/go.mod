module github.com/micro/services/portfolio/helpers/textenhancer/cli

replace github.com/micro/services/portfolio/helpers/textenhancer => ../

replace github.com/micro/services/portfolio/helpers/microgorm => ../../microgorm

replace github.com/micro/services/portfolio/helpers/passwordhasher => ../../passwordhasher

replace github.com/micro/services/portfolio/users => ../../../users

replace github.com/micro/services/portfolio/stocks => ../../../stocks

go 1.12

require github.com/micro/services/portfolio/helpers/textenhancer v0.0.0-00010101000000-000000000000
