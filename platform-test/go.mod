module platform-test

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro/v2 v2.0.1-0.20200212105717-d76baf59de2e
	github.com/micro/micro/v2 v2.0.1-0.20200213093446-b8e350745a0e
)

replace github.com/micro/go-micro/v2 => ../../go-micro
