package database

import (
	"go-echo-todo-app/entities"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(identifier int) (todo entities.Todo, err error) {
	repo.First(&todo, identifier)
	if err != nil {
		panic(err.Error())
	}
	return todo, nil
}

func (repo *TodoRepository) AddTodo(title string) (insertId int64, err error) {
	todo := entities.Todo{Title: title}
	db := repo.Create(&todo)
	if db.Error != nil {
		panic(db.Error)
	}
	return int64(todo.ID), nil
}
