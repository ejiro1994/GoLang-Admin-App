package db

import (
	"fmt"
	"go-admin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:rootroot@/sys"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		fmt.Println("Failed to prepare connection to database. DSN:", dsn)
		panic("could not connect to the database")
	}


	
	DB = db

	db.AutoMigrate(&models.User{}, &models.Role{})

}
