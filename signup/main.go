package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth/client"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/trmnl-core/services/signup/handler"
)

func main() {
	srv := service.New(
		service.Name("signup"),
	)

	auth := client.NewAuth()

	if err := srv.Handle(handler.NewSignup(srv, auth)); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
