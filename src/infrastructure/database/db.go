package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-echo-todo-app/interface/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

var DB_HOST = "db"
var DB_PORT = "3306"
var DB_USER = "root"
var DB_PASS = "maguroman2"
var DB_NAME = "todo_app"

func New() database.SqlHandler {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	fmt.Println(connectionString)
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error)
	}
	// defer connection.Close()
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = connection
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	// fmt.Println(r.Rows)
	return r.Rows.Close()
}
