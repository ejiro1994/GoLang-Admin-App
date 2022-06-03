package main

import (
	"go-admin/database"

	"go-admin/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
	
}
