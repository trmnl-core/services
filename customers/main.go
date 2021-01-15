package main

import (
	"github.com/trmnl-core/services/customers/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("customers"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(handler.New(srv))
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
