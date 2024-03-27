package main

import (
	"log"

	todo "github.com/qasimiss/http-rest-todo"
	"github.com/qasimiss/http-rest-todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}
}
