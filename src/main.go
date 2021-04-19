package main

import (
	"fmt"
	"github.com/labstack/echo"
	"go-echo-todo-app/infrastructure/database"
	"go-echo-todo-app/infrastructure/router"
	"go-echo-todo-app/interface/controller"
	"os"
)

func main() {
	sqlHandler := db.New()
	todoController := controller.New(sqlHandler)
	fmt.Println(todoController.Interactor.TodoRepository.FindById(1))
	e := echo.New()
	router.InitRouting(e, todoController)
	e.Logger.Fatal(e.Start(":8081"))
}
