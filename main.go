package main

import (
	handler "github.com/YakovAkk/test_project_go/app/handler"
	"github.com/YakovAkk/test_project_go/app/service"
	"github.com/YakovAkk/test_project_go/app/todo"
	"github.com/sirupsen/logrus"
)

func main() {

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(":8080", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

}
