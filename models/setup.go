package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
 func ConnectDatabase()  {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/point_of_sale"))
	if err != nil {
		panic(err)
	}

	errs := database.AutoMigrate(&Users{})
	if errs != nil {
		panic(errs)
	}

	DB =  database
 }