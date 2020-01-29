package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	rabbitmq "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/followers/handler"
	proto "github.com/micro/services/portfolio/followers/proto"
	"github.com/micro/services/portfolio/followers/storage/postgres"
	"github.com/pkg/errors"
)

func main() {
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
	b := rabbitmq.NewBroker(
		broker.Addrs(os.Getenv("MICRO_BROKER_ADDRESS")),
	)
	if err := b.Connect(); err != nil {
		panic(errors.Wrap(err, "Could not connect to the message broker"))
	}

	service := micro.NewService(
		micro.Name("kytra-v1-followers"),
		micro.Version("latest"),
	)
	service.Init()

	proto.RegisterFollowersHandler(service.Server(), handler.New(storageService, b))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
