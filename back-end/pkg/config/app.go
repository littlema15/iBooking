package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

const (
	dsn = "root:123321@tcp(localhost:3306)/iBooking?charset=utf8mb4&parseTime=true&loc=Local"
)

// Connect to the database
func Connect() {
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB return database
func GetDB() *gorm.DB {
	return db
}
