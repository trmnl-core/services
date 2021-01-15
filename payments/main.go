package main

import (
	"github.com/trmnl-core/services/payments/handler"
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
)

func main() {
	// Setup the service
	srv := service.New(
		service.Name("payments"),
	)

	// Register the provider
	srv.Handle(handler.New(srv))

	// Run the service
	if err := srv.Run(); err != nil {
		log.Fatalf("Error running service: %v", err)
	}
}
