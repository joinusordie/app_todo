package main

import (
	"log"

	todo "github.com/joinusordie/app_todo"
	"github.com/joinusordie/app_todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}