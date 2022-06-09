package models

import (
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 15
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)

	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": float64(int(total) / limit),
		},
	}
}
