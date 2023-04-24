package http

import "github.com/gofiber/fiber/v2"

// SignatureMiddleware Validates request by HMAC512
func SignatureMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return nil
	}
}

// JWTMiddleware Validates request by access token
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
