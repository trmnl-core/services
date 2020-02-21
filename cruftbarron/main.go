package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"cruftbarron/handler"
	"cruftbarron/subscriber"

	cruftbarron "cruftbarron/proto/cruftbarron"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.cruftbarron"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cruftbarron.RegisterCruftbarronHandler(service.Server(), new(handler.Cruftbarron))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.cruftbarron", service.Server(), new(subscriber.Cruftbarron))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
