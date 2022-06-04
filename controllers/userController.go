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

func CreateUser(c *fiber.Ctx) error {
	 var user models.User

	 if err := c.BodyParser(&user); err != nil {
		 return err
	 }


	 
	user.SetPassword("1234")

	db.DB.Create(&user)

	return c.JSON(user)

}