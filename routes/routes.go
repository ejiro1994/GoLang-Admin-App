package routes

import (
	"go-admin/controllers"
	"go-admin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// PUBLIC
	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)

	// MIDDLEWARE  CHECK FOR COOKIE
	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/v1/user", controllers.User)
	app.Post("/api/v1/logout", controllers.Logout)

	// USER CRUD METHODS (CREATE, READ, UPDATE & DELETE)
	app.Get("/api/v1/users", controllers.AllUsers)
	app.Post("/api/v1/users", controllers.CreateUser)
	app.Get("/api/v1/users/:id", controllers.GetUser)
	app.Put("/api/v1/users/:id", controllers.UpdateUser)
	app.Delete("/api/v1/users/:id", controllers.DeleteUser)

	// ROLE CRUD METHODS (CREATE, READ, UPDATE & DELETE)
	app.Get("/api/v1/roles", controllers.AllRoles)
	app.Post("/api/v1/roles", controllers.CreateRole)
	app.Get("/api/v1/roles/:id", controllers.GetRole)
	app.Put("/api/v1/roles/:id", controllers.UpdateRole)
	app.Delete("/api/v1/roles/:id", controllers.DeleteRole)

		// PRODUCTS CRUD METHODS (CREATE, READ, UPDATE & DELETE)
		app.Get("/api/v1/products", controllers.AllProducts)
		app.Post("/api/v1/products", controllers.CreateProduct)
		app.Get("/api/v1/products/:id", controllers.GetProduct)
		app.Put("/api/v1/products/:id", controllers.UpdateProduct)
		app.Delete("/api/v1/products/:id", controllers.DeleteProduct)
	
		app.Get("/api/v1/permissions", controllers.AllPermissions)

		app.Post("/api/v1/upload", controllers.Upload)
		app.Static("/api/v1/upload", "./uploads")


		app.Get("/api/v1/orders", controllers.AllOrders)


		app.Post("/api/v1/export", controllers.Export)

		app.Get("/api/v1/chart", controllers.Chart)


}
