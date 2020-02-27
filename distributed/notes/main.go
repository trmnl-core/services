package main

import (
	"notes/handler"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	notes "notes/proto/notes"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.distributed.notes"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	notes.RegisterNotesHandler(service.Server(), handler.NewHandler())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
