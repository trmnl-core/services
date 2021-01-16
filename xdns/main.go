package main

import (
	"github.com/trmnl-core/services/xdns/handler"
	pb "github.com/trmnl-core/services/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("xdns"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterXdnsHandler(srv.Server(), new(handler.Xdns))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
