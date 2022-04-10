package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:@/BookStore?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
