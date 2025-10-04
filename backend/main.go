package main

import (
	"visa-tracker/database"
	"visa-tracker/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
