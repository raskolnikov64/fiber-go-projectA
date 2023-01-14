package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() {

	var err error

	const MYSQL = "root:@tcp(127.0.0.1:3306)/go-fiber-p1?charset=utf8mb4&parseTime=True&loc=Local"

	DSN := MYSQL

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Can't connect to the Database")
	}

	fmt.Println("Successfully connected to the Database")
}
