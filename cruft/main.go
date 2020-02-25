package main

import (
	"cruft/handler"
	"cruft/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	cruft "cruft/proto/cruft"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.cruft"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cruft.RegisterCruftHandler(service.Server(), new(handler.Cruft))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.cruft", service.Server(), new(subscriber.Cruft))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
