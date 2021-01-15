package main

import (
	"github.com/trmnl-core/services/emails/handler"
	pb "github.com/trmnl-core/services/emails/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("emails"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterEmailsHandler(srv.Server(), handler.NewEmailsHandler())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
