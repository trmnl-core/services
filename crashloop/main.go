package main

import (
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.crashloop"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	time.Sleep(5 * time.Second)
	log.Fatal("Crash... ")
}
