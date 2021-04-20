package router

import (
	"github.com/labstack/echo"
	"go-echo-todo-app/interface/controller"
)

// InitRouting Initialize Router
func InitRouting(e *echo.Echo, todoController *controller.TodoController) {
	e.GET("/todos/:id", todoController.ReadTodoById)
	e.POST("/todos", todoController.CreateTodo)
}
