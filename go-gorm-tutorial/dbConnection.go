package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "somnath:somnath@tcp(localhost:3306)/go_db"

func dbConnection() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("Something went wrong while connecting to DB: ")

		panic(err.Error())
	}
	fmt.Println(" DB connection was successfully done ")
}
