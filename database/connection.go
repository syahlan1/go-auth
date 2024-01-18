package database

import (
	"github.com/syahlan1/go-auth.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(mysql.Open("root:@/go_auth"), &gorm.Config{})

	if err != nil {
		panic("couldn't connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
