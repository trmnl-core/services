package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"test3/handler"
	"test3/subscriber"

	test3 "test3/proto/test3"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test3"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test3.RegisterTest3Handler(service.Server(), new(handler.Test3))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test3", service.Server(), new(subscriber.Test3))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
