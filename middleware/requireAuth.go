package middleware

import "github.com/gofiber/fiber/v2"

func RequireAuth(c *fiber.Ctx) error {
	// Set a custom header on all responses:
	c.Set("X-Custom-Header", "Hello, World")

	// Go to next middleware:
	return c.Next()
}
