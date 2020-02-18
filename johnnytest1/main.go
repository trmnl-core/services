package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"johnnytest1/handler"
	"johnnytest1/subscriber"

	johnnytest1 "johnnytest1/proto/johnnytest1"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.johnnytest1"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	johnnytest1.RegisterJohnnytest1Handler(service.Server(), new(handler.Johnnytest1))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.johnnytest1", service.Server(), new(subscriber.Johnnytest1))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
