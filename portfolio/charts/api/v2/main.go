package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/services/portfolio/charts-api/handler"
	proto "github.com/micro/services/portfolio/charts-api/proto"
	auth "github.com/micro/services/portfolio/helpers/authentication"
	iex "github.com/micro/services/portfolio/helpers/iex-cloud"
)

func main() {
	service := micro.NewService(
		micro.Name("kytra-api-v2-charts"),
		micro.Version("latest"),
	)
	service.Init()

	auth, err := auth.New(os.Getenv("JWT_SIGNING_KEY"))
	if err != nil {
		fmt.Printf("Could not initiate auth package: %v\n", err)
		os.Exit(2)
	}

	iex, err := iex.New(os.Getenv("IEX_TOKEN"))
	if err != nil {
		fmt.Printf("Could not initiate iex package: %v\n.", err)
		os.Exit(2)
	}

	handler := handler.New(auth, iex, service.Client())
	proto.RegisterChartsHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
