package main

import (
	"crypto/tls"
	"net/http"
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

	// setup the proxy
	u, _ := url.Parse("https://kubernetes-dashboard.kubernetes-dashboard.svc.cluster.local")
	p := httputil.NewSingleHostReverseProxy(u)
	p.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	service.Handle("/", p)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
