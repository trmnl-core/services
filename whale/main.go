package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"whale/handler"
	"whale/subscriber"

	whale "whale/proto/whale"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.whale"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	whale.RegisterWhaleHandler(service.Server(), new(handler.Whale))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.whale", service.Server(), new(subscriber.Whale))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
