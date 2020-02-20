package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"test6/handler"
	"test6/subscriber"

	test6 "test6/proto/test6"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test6"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test6.RegisterTest6Handler(service.Server(), new(handler.Test6))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test6", service.Server(), new(subscriber.Test6))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
