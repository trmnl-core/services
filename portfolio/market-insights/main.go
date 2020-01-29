package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	iex "github.com/micro/services/portfolio/helpers/iex-cloud"
	"github.com/micro/services/portfolio/market-insights/generator"
	"github.com/micro/services/portfolio/market-insights/handler"
	proto "github.com/micro/services/portfolio/market-insights/proto"
	"github.com/micro/services/portfolio/market-insights/storage/postgres"
	"github.com/pkg/errors"
)

func main() {
	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-v1-market-insights"),
		micro.Version("latest"),
	)
	service.Init()

	// Initialize IEX Package
	iex, err := iex.New(os.Getenv("IEX_TOKEN"))
	if err != nil {
		fmt.Printf("Could not initiate iex package: %v\n.", err)
		os.Exit(2)
	}

	// Connect to the DB
	storageService, err := postgres.New(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		panic(errors.Wrap(err, "Could not connect to the database"))
	}
	defer storageService.Close()

	// TODO: MOVE TO CRON
	g := generator.New(iex, storageService, service.Client())
	go g.CreateDailyInsights()

	// Register to Service Discovery
	hander := handler.New(storageService)
	proto.RegisterMarketInsightsHandler(service.Server(), hander)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
