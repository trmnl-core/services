package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	auth "github.com/micro/services/portfolio/helpers/authentication"
	photos "github.com/micro/services/portfolio/helpers/photos"
	"github.com/micro/services/portfolio/search-api/handler"
	proto "github.com/micro/services/portfolio/search-api/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("kytra-api-v1-search"),
		micro.Version("latest"),
	)
	service.Init()

	pics, err := photos.New(os.Getenv("PHOTOS_ADDRESS"))
	if err != nil {
		fmt.Printf("Could not initiate photos package: %v\n.", err)
		os.Exit(2)
	}

	auth, err := auth.New(os.Getenv("JWT_SIGNING_KEY"))
	if err != nil {
		fmt.Printf("Could not initiate auth package: %v\n", err)
		os.Exit(2)
	}

	handler := handler.New(service.Client(), auth, pics)
	proto.RegisterSearchHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
