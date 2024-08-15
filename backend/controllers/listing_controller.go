package controllers

import (
	"strconv"

	"github.com/amrimuf/hompimRent/models"
	"github.com/amrimuf/hompimRent/services"
	"github.com/gofiber/fiber/v2"
)

type ListingController struct {
	ListingService *services.ListingService
}

func NewListingController(listingService *services.ListingService) *ListingController {
	return &ListingController{ListingService: listingService}
}

func (lc *ListingController) CreateListing(c *fiber.Ctx) error {
	var listing models.Listing
	if err := c.BodyParser(&listing); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	if err := lc.ListingService.CreateListing(&listing); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create listing")
	}
	return c.Status(fiber.StatusCreated).JSON(listing)
}

func (lc *ListingController) GetAllListings(c *fiber.Ctx) error {
	listings, err := lc.ListingService.GetAllListings()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve listings")
	}
	return c.JSON(listings)
}

func (lc *ListingController) GetListingByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid listing ID")
	}
	listing, err := lc.ListingService.GetListingByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Listing not found")
	}
	return c.JSON(listing)
}

func (lc *ListingController) UpdateListing(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid listing ID")
	}
	var listing models.Listing
	if err := c.BodyParser(&listing); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	listing.ID = id
	if err := lc.ListingService.UpdateListing(&listing); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update listing")
	}
	return c.JSON(listing)
}

func (lc *ListingController) DeleteListing(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid listing ID")
	}
	if err := lc.ListingService.DeleteListing(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete listing")
	}
	return c.SendStatus(fiber.StatusNoContent)
}