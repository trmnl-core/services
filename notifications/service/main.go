package main

import (
	"github.com/m3o/services/notifications/service/dao"
	"github.com/m3o/services/notifications/service/handler"
	"github.com/m3o/services/notifications/service/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	notifications "github.com/m3o/services/notifications/service/proto/notifications"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.notifications"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()
	dao.Init(service.Options().Store)
	// Register Handler
	notifications.RegisterNotificationsHandler(service.Server(), new(handler.Notifications))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.notifications", service.Server(), new(subscriber.Subscriber))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
