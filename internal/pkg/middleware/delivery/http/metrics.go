package http

import "github.com/gofiber/fiber/v2"

// время ответа запроса

func (mdw *MiddlewareManager) MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}
