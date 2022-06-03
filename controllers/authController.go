package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "ejiro",
	}

	user.LastName = "Akpoduado"

	return c.JSON(user)
}
