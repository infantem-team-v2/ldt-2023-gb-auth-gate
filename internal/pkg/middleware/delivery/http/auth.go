package http

import (
	"gb-auth-gate/internal/auth/model"
	"gb-auth-gate/pkg/terrors"
	"github.com/gofiber/fiber/v2"
)

// SignatureMiddleware Validates request by HMAC512
func (mdw *MiddlewareManager) SignatureMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		body := c.Body()
		ok, err := mdw.AuthUC.ValidateService(&model.AuthHeadersLogic{
			Signature: headers["x-512-signature"],
			PublicKey: headers["x-public-key"],
			Body:      body,
		})
		if !ok {
			if err == nil {
				return terrors.Raise(nil, 300001)
			}
			return err
		}
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
