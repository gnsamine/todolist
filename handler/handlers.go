package handlers

import (
	"todolist/skeleton"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func List(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var todos []skeleton.Todo_list
		return c.JSON(todos)
	}

}

func Create(db *gorm.DB) fiber.Handler {
	return nil

}

func Get(db *gorm.DB) fiber.Handler {
	return nil

}

func Update(db *gorm.DB) fiber.Handler {
	return nil

}

func Delete(db *gorm.DB) fiber.Handler {
	return nil

}
