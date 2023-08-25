package main

import (
	"fmt"
	"todolist/route"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=tor dbname=todos port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println(db)

	app := fiber.New()

	route.Router(app, db)

	app.Listen(":3000")

}
