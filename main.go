package main

import (
	"go-authentication-system/database"
	"go-authentication-system/routes"

	"github.com/gin-gonic/gin"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	// app := fiber.New()
	router := gin.New()
	// Databse Goes here
	database.ConnectDB()

	// Routes Goes here
	// routes.UserRoutes(app)
	routes.AuthRoutes(router)
	routes.HomeRoute(router)

	// app.Listen(":8000")
	router.Run(":8000")
}
