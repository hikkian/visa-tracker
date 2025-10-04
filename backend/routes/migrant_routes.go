package routes

import (
	"visa-tracker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/migrants", controllers.GetAllMigrants)
	api.Get("/expired", controllers.GetExpiredMigrants)
	api.Post("/migrants", controllers.CreateMigrant)
	api.Put("/migrants/:id", controllers.UpdateMigrant)
	api.Delete("/migrants/:id", controllers.DeleteMigrant)
}
