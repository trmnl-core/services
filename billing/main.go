package main

import (
	"github.com/trmnl-core/services/billing/handler"

	asproto "github.com/trmnl-core/services/alert/proto/alert"
	csproto "github.com/trmnl-core/services/customers/proto"
	nsproto "github.com/trmnl-core/services/namespaces/proto"
	pproto "github.com/trmnl-core/services/payments/proto"
	subproto "github.com/trmnl-core/services/subscriptions/proto"
	uproto "github.com/trmnl-core/services/usage/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("billing"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(handler.NewBilling(
		nsproto.NewNamespacesService("namespaces", srv.Client()),
		pproto.NewProviderService("payments", srv.Client()),
		uproto.NewUsageService("usage", srv.Client()),
		subproto.NewSubscriptionsService("subscriptions", srv.Client()),
		csproto.NewCustomersService("customers", srv.Client()),
		asproto.NewAlertService("alert", srv.Client()),
		nil,
	))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
