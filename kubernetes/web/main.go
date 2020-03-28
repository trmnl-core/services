package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.kubernetes"),
	)

	service.Init()

	// TODO: start the k8s dashboard
	u, _ := url.Parse("https://kubernetes-dashboard.kubernetes-dashboard.svc.cluster.local")
	service.Handle("/", httputil.NewSingleHostReverseProxy(u))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
