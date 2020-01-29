package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	auth "github.com/micro/services/portfolio/helpers/authentication"
	"github.com/micro/services/portfolio/registration-api/handler"
	proto "github.com/micro/services/portfolio/registration-api/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("kytra-api-v1-registration"),
		micro.Version("latest"),
	)
	service.Init()

	auth, err := auth.New(os.Getenv("JWT_SIGNING_KEY"))
	if err != nil {
		fmt.Printf("Could not initiate auth package: %v\n", err)
		os.Exit(2)
	}

	handler := handler.New(auth, service.Client())
	proto.RegisterRegistrationHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
