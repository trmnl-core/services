package main

import (
	"fmt"
	"os"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	news "github.com/micro/services/portfolio/helpers/news"
	"github.com/micro/services/portfolio/stock-news/handler"
	proto "github.com/micro/services/portfolio/stock-news/proto"
	"github.com/micro/services/portfolio/stock-news/storage/postgres"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

func main() {
	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-v1-stock-news"),
		micro.Version("latest"),
	)
	service.Init()

	// Initialize News Package
	news, err := news.New(os.Getenv("STOCK_NEWS_TOKEN"))
	if err != nil {
		fmt.Printf("Could not initiate news package: %v\n.", err)
		os.Exit(2)
	}

	// Connect to the Database (Postgres)
	db, err := postgres.New(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		fmt.Printf("Could not connect to DB: %v\n.", err)
		os.Exit(2)
	}
	defer db.Close()

	// Connect to Message Broker (RabbitMQ)
	broker := rabbitmq.NewBroker(
		broker.Addrs(os.Getenv("MICRO_BROKER_ADDRESS")),
	)
	if err := broker.Connect(); err != nil {
		panic(errors.Wrap(err, "Could not connect to the message broker"))
	}

	// Register to Service Discovery
	h := handler.New(news, db, service.Client(), broker)

	c := cron.New(cron.WithLocation(time.UTC))
	c.AddFunc("0 * * * *", h.FetchStockNews)
	c.AddFunc("0 6 * * *", h.FetchMarketNews)
	c.Start()
	defer c.Stop()

	go h.FetchMarketNews()

	proto.RegisterStockNewsHandler(service.Server(), h)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
