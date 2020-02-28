package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/micro/go-micro/v2/web"
	hello "github.com/micro/services/helloworld/proto/helloworld"
)

var (
	head = `<head><style>body {margin: 25px; font-family: sans-serif;}</style></head>`
	html = `<html>` + head + `<body><h1>Enter Name<h1><form method=post><input name=name type=text /></form></body></html>`
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.helloworld"),
	)

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form.Get("name")
			if len(name) == 0 {
				name = "World"
			}

			cli := service.Options().Service.Client()

			cl := hello.NewHelloworldService("go.micro.srv.helloworld", cli)
			rsp, err := cl.Call(context.Background(), &hello.Request{
				Name: name,
			})

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Write([]byte(`<html>` + head + `<body><h1>` + rsp.Msg + `</h1></body></html>`))
			return
		}

		fmt.Fprint(w, html)
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
