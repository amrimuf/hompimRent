package middleware

import (
	"github.com/amrimuf/hompimRent/models"
	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(requiredPermissions []string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        role, ok := c.Locals("role").(string)
		if !ok || role == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: role not found",
			})
		}

        permissions, ok := models.RolePermissions[models.Role(role)]
        if !ok {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Insufficient permissions"})
        }

        for _, required := range requiredPermissions {
            if contains(permissions, required) {
                return c.Next()
            }
        }

        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Insufficient permissions"})
    }
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}
