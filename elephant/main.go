package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"elephant/handler"
	"elephant/subscriber"

	elephant "elephant/proto/elephant"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.elephant"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	elephant.RegisterElephantHandler(service.Server(), new(handler.Elephant))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.elephant", service.Server(), new(subscriber.Elephant))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
