package routes

import (
	"go-admin/controllers"
	"go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/v1/user", controllers.User)
	app.Post("/api/v1/logout", controllers.Logout)


	app.Get("/api/v1/users", controllers.AllUsers)

	// CRUD METHODS (CREATE, READ, UPDATE & DELETE)
	app.Post("/api/v1/users", controllers.CreateUser)
	app.Get("/api/v1/users/:id", controllers.GetUser)
	app.Put("/api/v1/users/:id", controllers.UpdateUser)
	app.Delete("/api/v1/users/:id", controllers.DeleteUser)



}
