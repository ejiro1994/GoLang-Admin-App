package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-admin/models"
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

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})

}
