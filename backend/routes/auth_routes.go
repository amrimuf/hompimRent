package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
    app.Post("/api/register", controllers.Register)
    app.Post("/api/login", controllers.Login)
}
