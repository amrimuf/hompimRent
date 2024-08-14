package main

import (
	"github.com/amrimuf/hompimRent/database"
	"github.com/amrimuf/hompimRent/routes"
	"github.com/gofiber/fiber/v2"
    
    "log"
	"os"
)

func main() {
    app := fiber.New()

    database.ConnectDB()

    routes.SetupRoutes(app)

    port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
    app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

    log.Printf("Server starting on port %s\n", port)
    err := app.Listen(":" + port)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
