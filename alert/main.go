package main

import (
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/trmnl-core/services/alert/handler"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("alert"),
	)

	// Register Handler
	srv.Handle(handler.NewAlert())

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
