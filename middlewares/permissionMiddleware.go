package middlewares

// import (
// 	"go-admin/database"
// 	"go-admin/models"
// 	"go-admin/util"
// 	"strconv"
// 	"errors"

// 	"github.com/gofiber/fiber/v2"
// )
//  func IsAuthorized(c *fiber.Ctx, page string) error {
// 	cookie := c.Cookies("jwt")
	
// 	 Id, err := util.ParseJwt(cookie)
	 
// 	 if err != nil {
// 		return err
// 	}
// 	userId, _ := strconv.Atoi(Id)

// user := models.User {
// 		Id: uint(userId),
// 	}


// 	db.DB.Preload("Role").Find(&user)



// 	role := models.Role{
// 		Id: user.RoleId,
// 	}

// 	db.DB.Preload("Permissions").Find(&role)

// 	if c.Method() == "GET" {
// 		for _, permission := range role.Permissions {
// 			if permission.Name == "view" + page || permission.Name == "edit_"+page {
// 				return nil
// 			}
// 		} 
// 	} else {
// 		for _, permission := range role.Permissions {
// 			if permission.Name == "edit"+page {
// 				return nil
// 			}
// 		}
// 	}
// 	c.Status(fiber.StatusUnauthorized)
// 	return errors.New("unauthorized")

//  }