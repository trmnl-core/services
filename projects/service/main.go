package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/m3o/services/projects/service/handler"
	pb "github.com/m3o/services/projects/service/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.projects"),
		micro.Version("latest"),
	)

	service.Init()

	pb.RegisterProjectsHandler(service.Server(), handler.New(service))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
