package route

import (
	handlers "todolist/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {
	// Create a router group for '/todos' routes
	todosGroup := app.Group("/todos")

	// Define the routes within the '/todos' group
	todosGroup.Get("", handlers.List(db))
	todosGroup.Post("", handlers.Create(db))
	todosGroup.Get("/:id", handlers.ListSpecificTodo(db))
	todosGroup.Put("/:id", handlers.Update(db))
	todosGroup.Delete("/:id", handlers.Delete(db))
}
