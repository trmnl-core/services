package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/services/distributed/api/handler"
	pb "github.com/micro/services/distributed/api/proto"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.distributed"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	h := handler.NewHandler(service.Client())
	pb.RegisterDistributedHandler(service.Server(), h)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
