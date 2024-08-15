package controllers

import (
	"github.com/amrimuf/hompimRent/models"
	"github.com/amrimuf/hompimRent/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
    AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
    return &AuthController{AuthService: authService}
}
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
	}

	if err := c.AuthService.Register(user, req.Password); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to register user")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	token, err := c.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}