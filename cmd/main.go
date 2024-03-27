package main

import (
	"log"

	todo "github.com/qasimiss/http-rest-todo"
	"github.com/qasimiss/http-rest-todo/pkg/handler"
	"github.com/qasimiss/http-rest-todo/pkg/repository"
	"github.com/qasimiss/http-rest-todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}
}
