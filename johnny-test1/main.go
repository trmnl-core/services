package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"johnny-test1/handler"
	"johnny-test1/subscriber"

	johnny "johnny-test1/proto/johnny"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.johnny"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	johnny.RegisterJohnnyHandler(service.Server(), new(handler.Johnny))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.johnny", service.Server(), new(subscriber.Johnny))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
