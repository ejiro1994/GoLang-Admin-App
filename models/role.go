package models

type Role struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Permissions []Permission `json:"permisson" gorm:"many2many:role_permissions"`
}