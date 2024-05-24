package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rneiva/go-jwt-auth-api/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/user")

	api.Get("/get-user", controllers.User)

	api.Post("/register", controllers.Register)

	api.Post("/login", controllers.Login)

	api.Post("logout", controllers.Logout)
}
