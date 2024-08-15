package controllers

import (
	"time"

	"github.com/amrimuf/hompimRent/models"
	"github.com/amrimuf/hompimRent/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

type UserController struct {
    Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
    var user models.User
    if err := ctx.BodyParser(&user); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }
    if err := c.Service.CreateUser(&user); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
    }
    return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
    users, err := c.Service.GetAllUsers()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
    }
    return ctx.Status(fiber.StatusOK).JSON(users)
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
    id, err := uuid.FromString(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }
    user, err := c.Service.GetUserByID(id)
    if err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }
    return ctx.Status(fiber.StatusOK).JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
    var user models.User
    if err := ctx.BodyParser(&user); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }
    id, err := uuid.FromString(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
    }
	user.ID = id
    user.UpdatedAt = time.Now()
    if err := c.Service.UpdateUser(&user); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
    }
    return ctx.Status(fiber.StatusOK).JSON(user)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
    id, err := uuid.FromString(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }
    if err := c.Service.DeleteUser(id); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
    }
    return ctx.Status(fiber.StatusNoContent).SendString("")
}
