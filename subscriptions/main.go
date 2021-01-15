package main

import (
	paymentsproto "github.com/trmnl-core/services/payments/proto"
	"github.com/trmnl-core/services/subscriptions/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("subscriptions"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(handler.New(
		paymentsproto.NewProviderService("payments", srv.Client()),
	))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
