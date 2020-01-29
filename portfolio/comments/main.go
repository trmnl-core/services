package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/comments/handler"
	proto "github.com/micro/services/portfolio/comments/proto"
	"github.com/micro/services/portfolio/comments/storage/postgres"
	"github.com/pkg/errors"
)

func main() {
	// Connect to the database
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
	broker := rabbitmq.NewBroker(
		broker.Addrs(os.Getenv("MICRO_BROKER_ADDRESS")),
	)
	if err := broker.Connect(); err != nil {
		panic(errors.Wrap(err, "Could not connect to the message broker"))
	}

	service := micro.NewService(
		micro.Name("kytra-v1-comments"),
		micro.Version("latest"),
	)
	service.Init()

	handler := handler.New(storageService, broker)
	proto.RegisterCommentsHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
