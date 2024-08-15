package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var errorLog *log.Logger

func init() {
    if err := os.MkdirAll("logs", os.ModePerm); err != nil {
        log.Fatalf("Error creating logs directory: %v", err)
    }
    
    file, err := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }
    errorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func ErrorLogger() fiber.Handler {
    return func(c *fiber.Ctx) error {
        err := c.Next()
        
        if err != nil {
            errorLog.Printf("Error: %v | Method: %s | URL: %s", err, c.Method(), c.OriginalURL())

            if fiberErr, ok := err.(*fiber.Error); ok {
                return c.Status(fiberErr.Code).JSON(fiber.Map{
                    "error": fiberErr.Message,
                })
            }

            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "An internal server error occurred",
            })
        }
        
        return nil
    }
}