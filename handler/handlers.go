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

func ListSpecificTodo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Params("id")

		var todo skeleton.Todo_list
		err := db.First(&todo, ID).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
			}
			return err
		}

		return c.JSON(todo)
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

func Update(db *gorm.DB) fiber.Handler {
	return nil

}

func Delete(db *gorm.DB) fiber.Handler {
	return nil

}
