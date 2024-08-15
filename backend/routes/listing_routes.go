package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/amrimuf/hompimRent/middleware"
	"github.com/gofiber/fiber/v2"
)

func ListingRoutes(router fiber.Router, listingCtrl *controllers.ListingController) {
    router.Post("/", middleware.RoleMiddleware([]string{"create_listing"}), listingCtrl.CreateListing)
    router.Get("/", listingCtrl.GetAllListings)
    router.Get("/:id", listingCtrl.GetListingByID)
    router.Put("/:id", listingCtrl.UpdateListing)
    router.Delete("/:id", listingCtrl.DeleteListing)
}
