package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"barfoo/handler"
	"barfoo/subscriber"

	barfoo "barfoo/proto/barfoo"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.barfoo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	barfoo.RegisterBarfooHandler(service.Server(), new(handler.Barfoo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.barfoo", service.Server(), new(subscriber.Barfoo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
