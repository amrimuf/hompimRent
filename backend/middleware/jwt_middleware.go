package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
    secretKey := os.Getenv("JWT_SECRET")
    return func(c *fiber.Ctx) error {
        authHeader := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
        if authHeader == "" {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "error": "Missing Authorization header",
            })
        }

        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(authHeader, &claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(secretKey), nil
        })

        if err != nil || !token.Valid {
            log.Printf("JWT error: %v", err)
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid or expired token",
            })
        }

        // Attach user information to context (optional)
        if userID, ok := claims["user_id"].(string); ok {
            c.Locals("user_id", userID)
        } else {
            return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid token payload",
            })
        }

        return c.Next()
    }
}
