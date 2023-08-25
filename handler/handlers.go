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

	return func(c *fiber.Ctx) error {
		var newDuty skeleton.Todo_list

		err := c.BodyParser(&newDuty)
		if err != nil {
			return err
		}

		if err := db.Create(&newDuty).Error; err != nil {
			return err
		}

		return c.JSON(newDuty)

	}

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
