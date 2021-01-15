package main

import (
	"github.com/trmnl-core/services/api/client/handler"
	client "github.com/trmnl-core/services/api/client/proto/client"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/api"
	log "github.com/micro/micro/v3/service/logger"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("client"),
	)

	// Register Handler
	client.RegisterClientHandler(srv.Server(), &handler.Client{srv.Client()}, api.WithEndpoint(
		// TODO: remove when api supports Call method as default for /foo singular paths
		&api.Endpoint{
			Name:    "Client.Call",
			Path:    []string{"^/client?$"},
			Method:  []string{"GET", "POST"},
			Handler: "rpc",
		},
	))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
