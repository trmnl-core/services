package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"appletree/handler"
	"appletree/subscriber"

	appletree "appletree/proto/appletree"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.appletree"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	appletree.RegisterAppletreeHandler(service.Server(), new(handler.Appletree))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.appletree", service.Server(), new(subscriber.Appletree))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
