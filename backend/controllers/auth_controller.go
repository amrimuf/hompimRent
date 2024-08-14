package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
    return c.SendString("User registered")
}

func Login(c *fiber.Ctx) error {
    return c.SendString("User logged in")
}
