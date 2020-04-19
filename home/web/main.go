package main

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.home"),
	)

	service.Init()

	// we have to proxy a number of routes
	rp := new(httputil.ReverseProxy)

	// using web.micro.mu internally
	rp.Director = func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = "web.micro.mu"
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/home")
	}

	// the following paths are served as web apps
	// when micro.mu becomes home this should not be needed
	service.Handle("/blog", rp)
	service.Handle("/docs", rp)
	service.Handle("/projects", rp)
	service.Handle("/usage", rp)
	service.Handle("/update", rp)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
