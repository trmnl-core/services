package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func kubeURL() string {
//	host := "https://" + os.Getenv("KUBERNETES_SERVICE_HOST") + ":" + os.Getenv("KUBERNETES_SERVICE_PORT")
//	path := "/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/"
	host := "https://kubernetes-dashboard.kubernetes-dashboard.svc.cluster.local"
	path := "/"
	return host + path
}

func main() {
	service := web.NewService(
		web.Name("go.micro.web.kubernetes"),
	)

	service.Init()

	// TODO: start the k8s dashboard

	// setup the proxy
	u, _ := url.Parse(kubeURL())
	p := httputil.NewSingleHostReverseProxy(u)
	p.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	service.Handle("/", p)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
