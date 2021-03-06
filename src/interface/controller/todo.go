package controller

import (
	"go-echo-todo-app/entities"
	"go-echo-todo-app/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func New(repo usecase.TodoRepository) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: repo,
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
	c.String(200, string(rune(insertId)))
	return nil
}
