package route

import (
	handlers "todolist/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {
	app.Get("/todos", handlers.List(db))
	app.Post("/todos", handlers.Create(db))
	app.Get("/todos/:id", handlers.ListSpecificTodo(db))
	app.Put("/todos/:id", handlers.Update(db))
	app.Delete("/todos/:id", handlers.Delete(db))
}
