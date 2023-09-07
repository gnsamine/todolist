package main

import (
	"fmt"
	"log"
	"os"
	"todolist/route"
	"todolist/storage"

	skeleton "todolist/skeleton"

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

	config := &storage.Config{
		Host:     os.Getenv("DBHost"),
		Port:     os.Getenv("DBPort"),
		Password: os.Getenv("DBPassword"),
		User:     os.Getenv("DBUser"),
		SSLMode:  os.Getenv("DBSSLMode"),
		DBName:   os.Getenv("DBDBName"),
		TimeZone: os.Getenv("DBTimezone"),
	}

	dsn := storage.Informations(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println(db)

	db.AutoMigrate(&skeleton.Todo_table{}) // This creates the table if it doesn't exist

	app := fiber.New()

	//Middlewares
	app.Use(logger.New())  //records the details of incoming requests when any HTTP request is made. This can be used for purposes such as debugging and performance optimization.
	app.Use(recover.New()) //catches any errors that may cause the program to crash or interrupt and keep the server running.
	app.Use(cors.New())    //It helps applications bypass CORS restrictions by providing appropriate responses that allow or deny HTTP requests access to their resources.

	route.Router(app, db)

	app.Listen(":3000")

}
