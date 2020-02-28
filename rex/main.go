package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"rex-srv/handler"
	"rex-srv/subscriber"

	rex "rex-srv/proto/rex"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.rex"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	rex.RegisterRexHandler(service.Server(), new(handler.Rex))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.rex", service.Server(), new(subscriber.Rex))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
