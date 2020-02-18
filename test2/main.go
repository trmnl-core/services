package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"test2/handler"
	"test2/subscriber"

	test2 "test2/proto/test2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test2"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test2.RegisterTest2Handler(service.Server(), new(handler.Test2))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test2", service.Server(), new(subscriber.Test2))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
