package env

import (
	"os"
)

var DB_USER = os.Getenv("MYSQL_USERS")
var DB_PASSWORD = os.Getenv("MYSQL_PASSWORDS")
var DB_HOST = os.Getenv("MYSQL_HOST")
var DB_NAME = os.Getenv("MYSQL_DATABASES")
var DB_PORT = os.Getenv("DB_PORT")
