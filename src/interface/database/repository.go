package database

import (
	"fmt"
	"go-echo-todo-app/entities"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(identifier int) (todo entities.Todo, err error) {
	row, err := repo.Query("SELECT id, title FROM todos WHERE id = ?", identifier)
	defer row.Close()
	fmt.Print(row)
	if err != nil {
		panic(err.Error)
	}
	var id int
	var title string
	row.Next()
	if err = row.Scan(&id, &title); err != nil {
		panic(err.Error())
	}
	todo.ID = id
	todo.Title = title
	return todo, nil
}
