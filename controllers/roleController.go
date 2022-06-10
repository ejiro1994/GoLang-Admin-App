package controllers

import (
	db "go-admin/database"
	"go-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	db.DB.Find(&roles)

	return c.JSON(roles)

}

type RoleDTO struct{
	Name string
	Permissions []float64
 }



func CreateRole(c *fiber.Ctx) error {
	var roleDTO RoleDTO
 
	err := c.BodyParser(&roleDTO)
   
	if err != nil {
	  return err
	}
   
	list := roleDTO.Permissions
   
	permissions := make([]models.Permission, len(list))
   
	for i, permissionId := range list {
	  permissions[i] = models.Permission{
		Id: uint(permissionId),
	  }
	}
   
	role := models.Role{
	  Name:        roleDTO.Name,
	  Permissions: permissions,
	}
   
	db.DB.Create(&role)
   
	return c.JSON(role)
  
}

func GetRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	db.DB.Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO RoleDTO

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	list := roleDTO.Permissions
   
	permissions := make([]models.Permission, len(list))
   
	for i, permissionId := range list {
	  permissions[i] = models.Permission{
		Id: uint(permissionId),
	  }
	}

	var result interface{}
	db.DB.Table("role_permissions").Where("role_id", id).Delete(result)
   
	role := models.Role{
		Id: uint(id),
	  Name:        roleDTO.Name,
	  Permissions: permissions,
	}

	db.DB.Model(&role).Updates(role)

	return c.JSON(role)

}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	db.DB.Delete(&role)

	return nil

}
