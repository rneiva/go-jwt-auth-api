package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/user")

	api.Get("/get-user", func(c *fiber.Ctx) error {
		return c.SendString("Get all users")
	})
}
