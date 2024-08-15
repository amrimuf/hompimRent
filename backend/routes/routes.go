package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
    ListingController *controllers.ListingController
}

func SetupRoutes(app *fiber.App, c Controllers) {
	AuthRoutes(app)      
	UserRoutes(app) 
	ListingRoutes(app, c.ListingController)
}
