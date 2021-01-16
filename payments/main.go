package main

import (
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/trmnl-core/services/payments/handler"
)

func main() {
	srv := service.New(
		service.Name("payments"),
	)

	srv.Handle(handler.New(srv))

	if err := srv.Run(); err != nil {
		log.Fatalf("error running service: %v", err)
	}
}
