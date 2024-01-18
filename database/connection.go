package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:@/go_auth"), &gorm.Config{})

	if err != nil {
		panic("couldn't connect to database")
	}

}
