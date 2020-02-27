module github.com/micro/services/distributed/api

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro/v2 v2.1.2
	github.com/micro/services/notes v0.0.0-00010101000000-000000000000
)

replace github.com/micro/services/notes => ../../notes
