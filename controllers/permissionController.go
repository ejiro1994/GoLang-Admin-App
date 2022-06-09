package controllers

import (
	"go-admin/models"
	"github.com/gofiber/fiber/v2"
	"go-admin/database"

)


func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	db.DB.Find(&permissions)
 
	return c.JSON(permissions)


}