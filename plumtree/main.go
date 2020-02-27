package main

import (
	"fmt"
	"plumtree/handler"
	"plumtree/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	plumtree "plumtree/proto/plumtree"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.plumtree"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	plumtree.RegisterPlumtreeHandler(service.Server(), new(handler.Plumtree))

	fmt.Println("TEST")

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.plumtree", service.Server(), new(subscriber.Plumtree))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
