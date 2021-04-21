package database

import (
	"gorm.io/gorm"
)

type SqlHandler interface {
	First(interface{}, ...interface{}) *gorm.DB
	// Query(string, ...interface{}) (Row, error)
	Create(interface{}) *gorm.DB
}

// type Result interface {
// 	LastInsertId() (int64, error)
// 	RowsAffected() (int64, error)
// }

// type DB struct {
// 	Error        error
// 	RowsAffected int64
// 	clone        int
// }

// type Row interface {
// 	Scan(...interface{}) error
// 	Next() bool
// 	Close() error
// }
