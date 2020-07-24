package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	"github.com/m3o/services/kubernetes/service/handler"
	pb "github.com/m3o/services/kubernetes/service/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.kubernetes"),
		micro.Version("latest"),
	)

	service.Init()

	pb.RegisterKubernetesHandler(service.Server(), handler.New(service))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
