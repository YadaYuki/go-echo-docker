package controller

import (
	"github.com/labstack/echo"
	"go-echo-todo-app/entities"
	"go-echo-todo-app/interface/database"
	"go-echo-todo-app/usecase"
	"strconv"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func New(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) ReadTodoById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.FindById(id)
	if err != nil {
		return err
	}
	c.JSON(200, todo)
	return nil
}

func (controller *TodoController) CreateTodo(c echo.Context) error {
	todo := new(entities.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	insertId, err := controller.Interactor.AddTodo(todo.Title)
	if err != nil {
		return err
	}
	c.String(200, string(insertId))
	return nil
}
