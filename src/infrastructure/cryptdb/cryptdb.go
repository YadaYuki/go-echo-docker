package cryptdb

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-echo-todo-app/infrastructure/env"
	"go-echo-todo-app/interface/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var keyText = "astaxie12798akljzmknm.ahkjkljl;k" // aes key

func New() database.SqlHandler {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = connection
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	cryptArgs, err := cryptDataArgs(args...)
	if err != nil {
		return nil, err
	}
	result, err := handler.Conn.Exec(statement, cryptArgs...)
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

func cryptDataArgs(args ...interface{}) ([]interface{}, error) {
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		panic("Failed to Instantiate Cipher")
	}
	cryptDataArray := []interface{}{}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	for _, arg := range args {
		str, ok := arg.(string)
		if !ok {
			panic(fmt.Sprintf("Failed to Convert %s to string", arg))
		}
		cipherText := make([]byte, len([]byte(str)))
		cfb.XORKeyStream(cipherText, []byte(str))
		cryptDataArray = append(cryptDataArray, fmt.Sprintf("%x", cipherText))
	}
	return cryptDataArray, nil
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
	return r.Rows.Close()
}
