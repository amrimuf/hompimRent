package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authCtrl *controllers.AuthController) {
	auth := app.Group("/auth")

    auth.Post("/register", authCtrl.Register)
    auth.Post("/login", authCtrl.Login)
}
