package main

import (
	"api/handler"
	pb "api/proto/graphql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.graphql"),
	)

	service.Init()

	pb.RegisterGraphqlHandler(service.Server(), new(handler.Graphql), api.WithEndpoint(
		&api.Endpoint{
			Name:    "Graphql.Call",
			Path:    []string{"^/graphql?$"},
			Method:  []string{"GET", "POST"},
			Handler: "api",
		},
	))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
