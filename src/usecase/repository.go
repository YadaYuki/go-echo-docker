package usecase

import (
	"go-echo-todo-app/entities"
)

type TodoRepository interface {
	FindById(int) (entities.Todo, error)
}
