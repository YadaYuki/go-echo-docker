package main

import (
	"net/http"
	"database/sql"
	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
)


func main(){
	DB_HOST := "database"
	DB_PORT := "3306"
	DB_USER := "user"
	DB_PASS := "password"
	DB_NAME := "sampledb"
	connectionString :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",DB_USER, DB_PASS,DB_HOST, DB_PORT, DB_NAME)
	log.Println(DB_USER+","+DB_PASS)
	_, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
