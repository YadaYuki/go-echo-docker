package cryptdb

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"fmt"
	"go-echo-todo-app/entities"
	"go-echo-todo-app/infrastructure/env"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

var keyText = os.Getenv("KEY_TEXT") // aes key

func New() *SqlHandler {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = connection
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (sql.Result, error) {
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

func (handler *SqlHandler) Query(statement string, args ...interface{}) (*SqlRow, error) {
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
	columns, _ := r.Rows.Columns()
	encryptDestPointerArray := make([]interface{}, len(columns))
	for i := range encryptDestPointerArray {
		encryptDestPointerArray[i] = new([]byte)
	}
	if err := r.Rows.Scan(encryptDestPointerArray...); err != nil {
		return err
	}
	decryptDataArray, err := decryptDataArgs(encryptDestPointerArray...)
	if err != nil {
		panic(err.Error())
	}
	for i := range decryptDataArray {
		// TODO: Correspond to other type(int,float...)
		destItem, ok := dest[i].(*string)
		if !ok {
			panic("failed to cast destItem to *string")
		}
		// overwrite the value pointed to by the pointer
		*destItem = string(decryptDataArray[i])
	}
	return nil
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}

func (r SqlHandler) FindById(identifier int) (todo entities.Todo, err error) {
	row, err := r.Query("SELECT title FROM todos where id=?", identifier)
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

func (r SqlHandler) AddTodo(todo string) (insertId int64, err error) {
	result, err := r.Execute("INSERT INTO todos(title) VALUES (?)", todo)
	if err != nil {
		panic(err.Error())
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return insertId, nil
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

func decryptDataArgs(encryptedDataPointerArray ...interface{}) ([][]byte, error) {
	decryptDataArray := [][]byte{}
	for _, encryptedData := range encryptedDataPointerArray {
		byteEncryptedData, ok := encryptedData.(*[]byte)
		if !ok {
			panic(fmt.Sprintf("Failed to Convert to String. %s", *byteEncryptedData))
		}
		decryptData, err := getDecryptData(*byteEncryptedData)
		if err != nil {
			panic(err)
		}
		decryptDataArray = append(decryptDataArray, decryptData)
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
