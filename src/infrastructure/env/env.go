package env

import (
	"os"
)

var DB_USER = "root"
var DB_PASSWORD = os.Getenv("MYSQL_ROOT_PASSWORD")
var DB_HOST = os.Getenv("MYSQL_HOST")
var DB_NAME = os.Getenv("MYSQL_DATABASE")
var DB_PORT = os.Getenv("DB_PORT")
