package usecase

import (
	"go-echo-todo-app/entities"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) TodoById(identifier int) (todo entities.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(identifier)
	return
}
