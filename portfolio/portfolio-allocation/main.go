package main

import (
	"fmt"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/portfolio-allocation/handler"
	proto "github.com/micro/services/portfolio/portfolio-allocation/proto"
)

func main() {
	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-v1-portfolio-allocation"),
		micro.Version("latest"),
	)
	service.Init()

	// Register to Service Discovery
	hander := handler.New(service.Client())
	proto.RegisterPortfolioAllocationHandler(service.Server(), hander)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
