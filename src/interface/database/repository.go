package database

import (
	"go-echo-todo-app/entities"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(identifier int) (todo entities.Todo, err error) {
	row, err := repo.Query("SELECT id, title, created_at FROM todos WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var title string
	row.Next()
	if err = row.Scan(&id, &title); err != nil {
		return
	}
	todo.ID = id
	todo.Title = title
	return
}
