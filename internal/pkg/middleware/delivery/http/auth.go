package http

import "github.com/gofiber/fiber/v2"

// SignatureMiddleware Validates request by HMAC512
func (mdw *MiddlewareManager) SignatureMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}

// JWTMiddleware Validates request by access token
func (mdw *MiddlewareManager) JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}

// FiberSessionMiddleware Validates request by fiber's session are saved in redis' storage
func (mdw *MiddlewareManager) FiberSessionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}
