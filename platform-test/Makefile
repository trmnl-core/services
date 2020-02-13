
GOPATH:=$(shell go env GOPATH)


.PHONY: run
run:
	MICRO_AUTH=service MICRO_AUTH_EXCLUDE=PlatformTest.GetHealth MICRO_BROKER=service MICRO_CONFIG=service MICRO_REGISTRY=service MICRO_RUNTIME=service MICRO_STORE=service go run .
    

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=. --go_out=. proto/platform/platform.proto
    

.PHONY: build
build: proto

	go build -o platform-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t platform-srv:latest
