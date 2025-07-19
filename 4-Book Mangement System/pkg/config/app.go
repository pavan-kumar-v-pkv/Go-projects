package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL driver
)

var (
	// DB is the global database connection
	db *gorm.DB
)

// ConnectDB initializes the database connection
func ConnectDB() {
	d, err := gorm.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	if db == nil {
		ConnectDB()
	}
	return db
}
