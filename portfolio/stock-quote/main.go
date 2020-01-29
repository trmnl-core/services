package main

import (
	"fmt"
	"os"
	"time"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	iex "github.com/micro/services/portfolio/helpers/iex-cloud"
	"github.com/micro/services/portfolio/stock-quote/handler"
	proto "github.com/micro/services/portfolio/stock-quote/proto"
)

func main() {
	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-v1-stock-quote"),
		micro.Version("latest"),
	)
	service.Init()

	// Initialize IEX Package
	iex, err := iex.New(os.Getenv("IEX_TOKEN"))
	if err != nil {
		fmt.Printf("Could not initiate iex package: %v\n.", err)
		os.Exit(2)
	}

	// Register to Service Discovery and setup Preemptive Caching
	hander := handler.New(iex)
	go setupPreemptiveCaching(hander)
	proto.RegisterStockQuoteHandler(service.Server(), hander)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func setupPreemptiveCaching(h *handler.Handler) {
	interval := 10 * time.Minute
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		fmt.Println("Starting Preemptive Cache Refresh...")
		if err := h.RefreshCache(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Ending Preemptive Cache Refresh...")
		}
	}
}
