package usecase

import "go-echo-todo-app/entities"

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) FindById(identifier int) (todo entities.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(identifier)
	return todo, err
}

func (interactor *TodoInteractor) AddTodo(todo string) (insertId int, err error) {
	insertId, err = interactor.TodoRepository.AddTodo(todo)
	return insertId, err
}
