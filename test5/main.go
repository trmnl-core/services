package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"test5/handler"
	"test5/subscriber"

	test5 "test5/proto/test5"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test5"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test5.RegisterTest5Handler(service.Server(), new(handler.Test5))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test5", service.Server(), new(subscriber.Test5))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
