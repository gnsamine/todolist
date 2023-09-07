package main

import (
	"fmt"
	"log"
	"todolist/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dsn := "host=localhost user=postgres password=tor dbname=todos port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println(db)

	app := fiber.New()

	//Middlewares
	app.Use(logger.New())  //records the details of incoming requests when any HTTP request is made. This can be used for purposes such as debugging and performance optimization.
	app.Use(recover.New()) //catches any errors that may cause the program to crash or interrupt and keep the server running.
	app.Use(cors.New())    //It helps applications bypass CORS restrictions by providing appropriate responses that allow or deny HTTP requests access to their resources.

	route.Router(app, db)

	app.Listen(":3000")

}
