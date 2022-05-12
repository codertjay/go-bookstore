package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// codertjay:thankgod12/go_test_db?charset=utf8&parseTime=True&loc=Local
func Connect() {
	d, err := gorm.Open("mysql",
		"codertjay:thankgod12@tcp(127.0.0.1:3306)/go_test_db?charset=utf8&parseTime=True&loc=Local")
		// "codertjay:thankgod12@tcp(localhost:3306)/go_test_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
