package routes

import (
	"go-admin/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/v1/register", controllers.Register)
	// app.Get("/other", controllers.Other)

}
