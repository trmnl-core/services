
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=. --go_out=. proto/cruftbarron/cruftbarron.proto
    

.PHONY: build
build: proto

	go build -o cruftbarron-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cruftbarron-srv:latest
