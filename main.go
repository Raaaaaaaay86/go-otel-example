package main

import (
	"github.com/raaaaaaaay86/go-otel-example/controller"
	"github.com/raaaaaaaay86/go-otel-example/pkg/tracex"
	"github.com/raaaaaaaay86/go-otel-example/pkg/tracex/exporter"
	"github.com/raaaaaaaay86/go-otel-example/repository"
	"github.com/raaaaaaaay86/go-otel-example/service"
	"log"
	"net/http"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	jaegerExporter, err := exporter.NewJaegerExporter("http://localhost:14268/api/traces")
	if err != nil {
		return err
	}

	httpTracerProvider, err := tracex.NewTracerProvider("http", jaegerExporter)
	if err != nil {
		return err
	}
	serviceTracerProvider, err := tracex.NewTracerProvider("service", jaegerExporter)
	if err != nil {
		return err
	}
	repositoryTracerProvider, err := tracex.NewTracerProvider("repository", jaegerExporter)
	if err != nil {
		return err
	}

	userRepository := repository.NewUserRepository(repositoryTracerProvider)
	userService := service.NewUserService(serviceTracerProvider, userRepository)
	userController := controller.NewUserController(httpTracerProvider, userService)

	http.HandleFunc("/findUser", userController.FindUserById)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
