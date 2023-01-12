package routes

import (
	"go-authentication-system/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
    app.Post("/user", controllers.CreateUser)
    app.Get("/user/:userId", controllers.GetUser)
    app.Put("/user/:userId", controllers.UpdateUser)
    app.Delete("/user/:userId", controllers.DeleteUser)
    app.Get("/users", controllers.GetAllUsers)
}