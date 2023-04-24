package http

import "github.com/gofiber/fiber/v2"

func StartLoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func EndLoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
