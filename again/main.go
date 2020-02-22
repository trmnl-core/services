package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"again/handler"
	"again/subscriber"

	again "again/proto/again"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.again"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	again.RegisterAgainHandler(service.Server(), new(handler.Again))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.again", service.Server(), new(subscriber.Again))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
