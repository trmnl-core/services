package main

import (
	"github.com/m3o/services/notifications/handler"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"

	notifications "github.com/m3o/services/notifications/proto/notifications"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("go.micro.service.notifications"),
	)

	// Register Handler
	notifications.RegisterNotificationsHandler(new(handler.Notifications))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}