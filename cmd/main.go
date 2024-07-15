package main

import (
	"log"

	"github.com/yelnar0112/todo-app"
)

// @title           Todo API
// @version         1.0
// @description     This is a sample server Todo API.
// @termsOfService  http://swagger.io/terms/
func main() {
	handles := new(todo.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handles.InitRoutes()); err != nil {
		log.Fatal("Error while running http server: %s", err.Error())
	}
}
