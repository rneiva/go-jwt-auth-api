package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rneiva/go-jwt-auth-api/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
