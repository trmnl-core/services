package main

import (
	"api/client"
	"api/handler"
	graphql "api/proto/graphql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.graphql"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Graphql srv client
		micro.WrapHandler(client.GraphqlWrapper(service)),
	)

	// Register Handler
	graphql.RegisterGraphqlHandler(service.Server(), new(handler.Graphql))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
