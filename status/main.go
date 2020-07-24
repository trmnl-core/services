package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/m3o/services/status/handler"
	status "github.com/m3o/services/status/proto/status"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.status"),
	)

	// Initialise service
	service.Init()
	// grab services to monitor
	svcs := service.Options().Config.Get("micro", "status", "services").StringSlice(nil)
	log.Infof("Services to monitor %+v", svcs)
	// Register Handler
	status.RegisterStatusHandler(service.Server(), handler.NewStatusHandler(svcs))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
