package gormdb

import (
	"database/sql"
	"fmt"
	"go-echo-todo-app/entities"
	"go-echo-todo-app/infrastructure/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	*gorm.DB
}

func New() *SqlHandler {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	sqlDb, _ := sql.Open("mysql", connectionString)
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn: sqlDb,
		}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.DB = db
	return sqlHandler
}

func (r SqlHandler) FindById(identifier int) (todo entities.Todo, err error) {
	r.First(&todo)
	if r.Error != nil {
		return entities.Todo{}, r.Error
	}
	return
}

func (r SqlHandler) AddTodo(title string) (insertId int, err error) {
	todo := entities.Todo{Title: title}
	result := r.Create(&todo)
	if result.Error != nil {
		return -1, result.Error
	}
	return todo.ID, nil
}
