package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

func ListingRoutes(app *fiber.App, listingController *controllers.ListingController) {
    listingGroup := app.Group("/listings")

    listingGroup.Post("/", listingController.CreateListing)
    listingGroup.Get("/", listingController.GetAllListings)
    listingGroup.Get("/:id", listingController.GetListingByID)
    listingGroup.Put("/:id", listingController.UpdateListing)
    listingGroup.Delete("/:id", listingController.DeleteListing)
}
