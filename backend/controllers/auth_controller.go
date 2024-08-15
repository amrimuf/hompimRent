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
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    user := &models.User{
        Username: req.Username,
        Email:    req.Email,
    }

    if err := c.AuthService.Register(user, req.Password); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
    }

    return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := ctx.BodyParser(&req); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    token, err := c.AuthService.Login(req.Email, req.Password)
    if err != nil {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
