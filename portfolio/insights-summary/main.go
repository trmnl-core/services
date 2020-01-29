package main

import (
	"fmt"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/insights-summary/handler"
	proto "github.com/micro/services/portfolio/insights-summary/proto"
)

func main() {
	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-v1-insights-summary"),
		micro.Version("latest"),
	)
	service.Init()

	// Register to Service Discovery
	hander := handler.New(service.Client())
	proto.RegisterInsightsSummaryHandler(service.Server(), hander)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
