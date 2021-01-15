package main

import (
	"github.com/trmnl-core/services/platform/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	srv := service.New(
		service.Name("platform"),
	)

	srv.Handle(handler.New(srv))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
