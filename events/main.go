package main

import (
	"github.com/m3o/services/events/handler"
	pb "github.com/m3o/services/events/proto"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
)

func main() {
	srv := service.New(
		service.Name("go.micro.service.events"),
	)

	pb.RegisterEventsHandler(handler.New(srv))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
