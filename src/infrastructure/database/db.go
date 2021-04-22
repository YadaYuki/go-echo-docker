package db

import (
	"database/sql"
	"fmt"
	"go-echo-todo-app/infrastructure/env"
	"go-echo-todo-app/interface/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() database.SqlHandler {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	sqlDb, err := sql.Open("mysql", connectionString)
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn: sqlDb,
		}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
