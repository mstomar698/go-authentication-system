package main

import (
	"go-authentication-system/database"
	"go-authentication-system/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Databse Goes here
	database.ConnectDB()

	// Routes Goes here
	routes.UserRoutes(app)

	app.Listen(":8000")
}
