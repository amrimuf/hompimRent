package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/amrimuf/hompimRent/middleware"
	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
    ListingController *controllers.ListingController
	UserController *controllers.UserController
	AuthController *controllers.AuthController
}

func SetupRoutes(app *fiber.App, c Controllers) {
	AuthRoutes(app, c.AuthController)

	// protected
    app.Use(middleware.JWTMiddleware())      

	userGroup := app.Group("/users")
	UserRoutes(userGroup, c.UserController) 
	
	listingGroup := app.Group("/protected/listings")
	ListingRoutes(listingGroup, c.ListingController)
}
