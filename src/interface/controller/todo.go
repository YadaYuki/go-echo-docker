package controller

import (
	"go-echo-todo-app/interface/database"
	"go-echo-todo-app/usecase"
	"github.com/labstack/echo"
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

func (controller *TodoController) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.FindById(id)
	if err != nil {
			return err
	}
	c.JSON(200, todo)
	return nil
}

