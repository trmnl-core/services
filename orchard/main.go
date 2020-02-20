package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"orchard/handler"
	"orchard/subscriber"

	orchard "orchard/proto/orchard"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.orchard"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	orchard.RegisterOrchardHandler(service.Server(), new(handler.Orchard))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.orchard", service.Server(), new(subscriber.Orchard))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
