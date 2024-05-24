package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rneiva/go-jwt-auth-api/database"
	"github.com/rneiva/go-jwt-auth-api/routes"
)

func main() {
	database.DBconn()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8000",
	}))

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
