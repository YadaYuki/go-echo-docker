package db

import (
	"database/sql"
	"fmt"
	"go-echo-todo-app/infrastructure/env"
	"go-echo-todo-app/interface/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type SqlHandler struct {
// 	Conn *sql.DB
// }

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

// func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
// 	res := SqlResult{}
// 	result, err := handler.Conn.Exec(statement, args...)
// 	if err != nil {
// 		return res, err
// 	}
// 	res.Result = result
// 	return res, nil
// }

// func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
// 	rows, err := handler.Conn.Query(statement, args...)
// 	if err != nil {
// 		return new(SqlRow), err
// 	}
// 	row := new(SqlRow)
// 	row.Rows = rows
// 	return row, nil
// }

// type SqlResult struct {
// 	Result sql.Result
// }

// func (r SqlResult) LastInsertId() (int64, error) {
// 	return r.Result.LastInsertId()
// }

// func (r SqlResult) RowsAffected() (int64, error) {
// 	return r.Result.RowsAffected()
// }

// type SqlRow struct {
// 	Rows *sql.Rows
// }

// func (r SqlRow) Scan(dest ...interface{}) error {
// 	return r.Rows.Scan(dest...)
// }

// func (r SqlRow) Next() bool {
// 	return r.Rows.Next()
// }

// func (r SqlRow) Close() error {
// 	return r.Rows.Close()
// }
