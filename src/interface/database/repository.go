package database

import (
	"go-echo-todo-app/entities"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(identifier int) (todo entities.Todo, err error) {
	// var todo entities.Todo
	repo.First(todo, "id = ?", identifier)
	if err != nil {
		panic(err.Error())
	}
	// var id int
	// var title string
	// row.Next()
	// if err = row.Scan(&id, &title); err != nil {
	// 	panic(err.Error())
	// }
	// todo.ID = id
	// todo.Title = title
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
