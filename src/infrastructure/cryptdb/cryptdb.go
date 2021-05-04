package cryptdb

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"fmt"
	"go-echo-todo-app/infrastructure/env"
	"go-echo-todo-app/interface/database"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

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
	cryptArgs, err := encryptDataArgs(args...)
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
	// https://groups.google.com/g/golang-nuts/c/_AAB2x1SZBk?pli=1
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	columns, _ := rows.Columns()
	data := make([]interface{}, len(columns))
	for i := range data {
		data[i] = new(interface{})
	}
	rows.Next()
	if err = rows.Scan(data...); err != nil {
		panic(err.Error())
	}
	row.Rows = rows
	return row, nil
}

func encryptDataArgs(args ...interface{}) ([]interface{}, error) {
	encryptDataArgs := []interface{}{}
	for _, arg := range args {
		str, ok := arg.(string)
		if !ok {
			panic(fmt.Sprintf("Failed to Convert %s to str", str))
		}
		encryptData, err := getEncryptData([]byte(str))
		if err != nil {
			panic(err)
		}
		encryptDataArgs = append(encryptDataArgs, encryptData)
	}
	return encryptDataArgs, nil
}

func getEncryptData(data []byte) (encryptedData []byte, err error) {
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		return nil, err
	}
	encryptedData = make([]byte, aes.BlockSize+len(data)) // BlockSize = 16
	iv := encryptedData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(c, iv)
	cfb.XORKeyStream(encryptedData[aes.BlockSize:], data)
	return encryptedData, nil
}

func decryptDataArgs(encryptedDataArray ...interface{}) ([]interface{}, error) {
	decryptDataArray := [][]byte{}
	for _, encryptedData := range encryptedDataArray {
		str, ok := encryptedData.(*string)
		if !ok {
			panic(fmt.Sprintf("Failed to Convert %s to str", *str))
		}
		decryptData, err := getDecryptData([]byte(*str))
		if err != nil {
			panic(err)
		}
		decryptDataArray = append(decryptDataArray, &decryptData)
	}
	return decryptDataArray, nil
}

func getDecryptData(encryptedData []byte) (decryptedData []byte, err error) {
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		return nil, err
	}
	iv := encryptedData[:aes.BlockSize]
	targetEncryptedData := encryptedData[aes.BlockSize:]
	decryptedData = make([]byte, len(targetEncryptedData))
	cfb := cipher.NewCFBDecrypter(c, iv)
	cfb.XORKeyStream(decryptedData, targetEncryptedData)
	return decryptedData, nil
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
