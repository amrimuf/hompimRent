package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

func ListingRoutes(app *fiber.App, listingCtrl *controllers.ListingController) {
    listingGroup := app.Group("/listings")

    listingGroup.Post("/", listingCtrl.CreateListing)
    listingGroup.Get("/", listingCtrl.GetAllListings)
    listingGroup.Get("/:id", listingCtrl.GetListingByID)
    listingGroup.Put("/:id", listingCtrl.UpdateListing)
    listingGroup.Delete("/:id", listingCtrl.DeleteListing)
}
