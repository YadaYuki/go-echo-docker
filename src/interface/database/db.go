package database

import (
	"gorm.io/gorm"
)

type SqlHandler interface {
	First(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
}
