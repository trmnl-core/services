package main

import (
	"net/http"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.distributed"),
		web.Version("latest"),
	)

	// Todo: Fix file serving
	service.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := "./app/build" + req.URL.Path
		log.Logf(log.InfoLevel, "Serving file: %v", path)
		http.ServeFile(w, req, path)
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
