package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/feeditems/handler"
	proto "github.com/micro/services/portfolio/feeditems/proto"
	"github.com/micro/services/portfolio/feeditems/storage/postgres"
	"github.com/pkg/errors"
)

func main() {
	// Connect to the Database (Postgres)
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

	// Connect to Message Broker (RabbitMQ)
	if err := broker.Init(); err != nil {
		panic(errors.Wrap(err, "Could not connect to the message broker"))
	}
	if err := broker.Connect(); err != nil {
		panic(errors.Wrap(err, "Could not connect to the message broker"))
	}

	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-v1-feed-items"),
		micro.Version("latest"),
	)
	service.Init()

	// Register to Service Discovery (Consul)
	hander := handler.New(storageService)
	proto.RegisterFeedItemsHandler(service.Server(), hander)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
