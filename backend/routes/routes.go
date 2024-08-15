package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
    ListingController *controllers.ListingController
	UserController *controllers.UserController
	AuthController *controllers.AuthController
}

func SetupRoutes(app *fiber.App, c Controllers) {
	AuthRoutes(app, c.AuthController)      
	UserRoutes(app, c.UserController) 
	ListingRoutes(app, c.ListingController)
}
