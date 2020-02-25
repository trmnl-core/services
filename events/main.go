package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"events/handler"
	"events/subscriber"

	events "events/proto/events"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.events"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	events.RegisterEventsHandler(service.Server(), new(handler.Events))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.events", service.Server(), new(subscriber.Events))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
