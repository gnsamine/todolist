package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	handlers "todolist/handler"
)

func main() {

	dsn := "host=localhost user=postgres password=tor dbname=todos port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println(db)

	app := fiber.New()

	app.Get("/todos", handlers.List(db))
	app.Post("/todos", handlers.Create(db))
	app.Get("/todos/:id", handlers.Get(db))
	app.Put("/todos/:id", handlers.Update(db))
	app.Delete("/todos/:id", handlers.Delete(db))

	app.Listen(":3000")

}
