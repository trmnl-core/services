
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=. --go_out=. proto/johnnytest1/johnnytest1.proto
    

.PHONY: build
build: proto

	go build -o johnnytest1-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t johnnytest1-srv:latest
