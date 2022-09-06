package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pingkuan/go-fiber-api/database"
	"github.com/pingkuan/go-fiber-api/routes"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	app.Use(cors.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
