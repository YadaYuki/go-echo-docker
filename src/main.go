package main

import (
	"go-echo-todo-app/infrastructure/gormdb"
	"go-echo-todo-app/infrastructure/router"
	"go-echo-todo-app/interface/controller"

	"github.com/labstack/echo"
)

func main() {
	sqlHandler := gormdb.New()
	todoController := controller.New(sqlHandler)
	e := echo.New()
	router.InitRouting(e, todoController)
	e.Logger.Fatal(e.Start(":8081"))
}
