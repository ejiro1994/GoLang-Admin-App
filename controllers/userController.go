package controllers

import(
	"go-admin/models"
	"go-admin/database"
	"github.com/gofiber/fiber/v2"

)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	db.DB.Find(&users)

	return c.JSON(users)


}

