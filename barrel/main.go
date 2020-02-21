package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"barrel/handler"
	"barrel/subscriber"

	barrel "barrel/proto/barrel"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.barrel"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	barrel.RegisterBarrelHandler(service.Server(), new(handler.Barrel))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.barrel", service.Server(), new(subscriber.Barrel))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
