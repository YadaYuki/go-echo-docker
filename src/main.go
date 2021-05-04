package main

import (
	"go-echo-todo-app/infrastructure/cryptdb"
	"go-echo-todo-app/infrastructure/router"
	"go-echo-todo-app/interface/controller"

	"github.com/labstack/echo"
)

func main() {
	sqlHandler := cryptdb.New()
	todoController := controller.New(sqlHandler)
	// fmt.Println(todoController.Interactor.TodoRepository.FindById(1))
	e := echo.New()
	router.InitRouting(e, todoController)
	e.Logger.Fatal(e.Start(":8081"))
}
