package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/api/users", controllers.GetUsers)
}
