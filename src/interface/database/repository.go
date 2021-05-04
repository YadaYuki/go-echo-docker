package database

import (
	"go-echo-todo-app/entities"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(identifier int) (todo entities.Todo, err error) {
	row, err := repo.Query("SELECT title FROM todos where id=?", identifier)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()
	var id int
	var title string
	row.Next()
	if err = row.Scan(&title); err != nil {
		panic(err.Error())
	}
	todo.ID = id
	todo.Title = title
	return todo, nil
}

func (repo *TodoRepository) AddTodo(todo string) (insertId int64, err error) {
	result, err := repo.Execute("INSERT INTO todos(title) VALUES (?)", todo)
	if err != nil {
		panic(err.Error())
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return insertId, nil
}
