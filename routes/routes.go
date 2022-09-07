package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pingkuan/go-fiber-api/handlers"
	"github.com/pingkuan/go-fiber-api/middlewares"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	api.Get("/", handlers.Running)

	user := api.Group("/users")
	user.Get("/profile", middlewares.Protect, handlers.GetUser)
	user.Put("/profile", middlewares.Protect, handlers.UpdateUser)
	user.Delete("/profile", middlewares.Protect, handlers.DeleteUser)
	user.Post("/", handlers.Register)
	user.Post("/login", handlers.AuthUser)

}
