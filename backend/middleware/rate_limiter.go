package middleware

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var (
    redisClient *redis.Client
    ctx         = context.Background()
)

func init() {
    redisClient = redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })

    _, err := redisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
}

func RateLimit(limit int, duration time.Duration) fiber.Handler {
    return func(c *fiber.Ctx) error {
        ip := c.IP()
        key := "rate_limit:" + ip

        rateLimiter := redisClient

        count, err := rateLimiter.Get(ctx, key).Int()
        if err != nil && err != redis.Nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
        }

        if count >= limit {
            return fiber.NewError(fiber.StatusTooManyRequests, "Rate limit exceeded")
        }

        if count == 0 {
            err := rateLimiter.Set(ctx, key, 1, duration).Err()
            if err != nil {
                return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
            }
        } else {
            err := rateLimiter.Incr(ctx, key).Err()
            if err != nil {
                return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
            }
        }

        return c.Next()
    }
}
