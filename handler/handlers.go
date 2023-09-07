package handlers

import (
	"todolist/skeleton"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var newDuty skeleton.Todo_table

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

func List(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var todos []skeleton.Todo_table

		if err := db.Find(&todos).Error; err != nil {
			return err
		}
		return c.JSON(todos)
	}

}

func ListSpecificTodo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		

		return nil

	}
}

func Update(db *gorm.DB) fiber.Handler {

	return nil

}

func Delete(db *gorm.DB) fiber.Handler {
	return nil

}
