package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/notifications/api"
	"github.com/micro/services/portfolio/notifications/consumer"
	proto "github.com/micro/services/portfolio/notifications/proto"
	"github.com/micro/services/portfolio/notifications/storage/postgres"
)

func main() {
	cmd.Init()

	// Setup the service
	service := micro.NewService(
		micro.Name("kytra-v1-notifications"),
		micro.Version("latest"),
	)
	service.Init()

	// Connect to the Database (Postgres)
	db, err := postgres.New(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		panic(errors.New("Could not connect to the database"))
	}
	defer db.Close()

	// Setup the broker (RabbitMQ)
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	// Create a Handler and subscribe to events
	c := consumer.New(service.Client(), db)
	broker.Subscribe("kytra-v1-posts-post-created", c.ConsumeNewPost, broker.Queue("notifications-post-created"))
	broker.Subscribe("kytra-v1-followers-new-follow", c.ConsumeNewFollow, broker.Queue("notifications-new-follow"))
	broker.Subscribe("kytra-v1-comments-comment-created", c.ConsumeNewComment, broker.Queue("notifications-comment-created"))

	// Run the service
	proto.RegisterNotificationsHandler(service.Server(), api.New(service.Client(), db))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
