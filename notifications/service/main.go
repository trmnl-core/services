package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/services/notifications/service/dao"
	"github.com/micro/services/notifications/service/handler"
	"github.com/micro/services/notifications/service/subscriber"

	notifications "github.com/micro/services/notifications/service/proto/notifications"
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
