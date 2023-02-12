package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// docker run --name mysql-standalone -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -e MYSQL_USER=user -e MYSQL_PASSWORD=password -p 3306:3306 -d mysql:5.6
	database, err := gorm.Open(mysql.Open("user:password@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Note{})
	DB.AutoMigrate(&User{})
}
