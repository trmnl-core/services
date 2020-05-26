package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/runtime"
)

const (
	// serviceRootFile is the file which a directory must contain to be considered a service
	serviceRootFile = "main.go"
	// source is the repository containing the source code
	source = "github.com/micro/services"
)

func main() {
	var services []string

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Base(path) != serviceRootFile {
			return nil
		}
		services = append(services, filepath.Dir(path))
		return nil
	}

	if err := filepath.Walk(".", walker); err != nil {
		logger.Fatal(err)
	}

	srv := micro.NewService()
	srv.Init()
	logger.Infof("Using %v runtime", srv.Options().Runtime)

	for _, name := range services {
		logger.Infof("Creating %v", name)

		err := srv.Options().Runtime.Create(
			&runtime.Service{Name: name, Source: source},
			runtime.CreateNamespace("micro"),
			runtime.CreateImage("docker.pkg.github.com/micro/services/"+strings.ReplaceAll(name, "/", "-")),
		)

		if err != nil {
			logger.Fatal(err)
		}
	}
}
