package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"test1/handler"
	"test1/subscriber"

	test1 "test1/proto/test1"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test1"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test1.RegisterTest1Handler(service.Server(), new(handler.Test1))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test1", service.Server(), new(subscriber.Test1))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
