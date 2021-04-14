package main

import (
	"go-echo-todo-app/infrastructure/database"
	"go-echo-todo-app/infrastructure/router"
	"go-echo-todo-app/interface/controller"
	"github.com/labstack/echo"
)

func main() {
	sqlHandler := database.New()
	todoController := controller.New(sqlHandler)
	e := echo.New()
	router.InitRouting(e,todoController)
	e.Logger.Fatal(e.Start(":8081"))
}
