package http

import (
	authInterface "bank_api/internal/auth/interface"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	AuthUC authInterface.UseCase `di:"authUC"`
	router fiber.Router
}

func (ah *AuthHandler) GetRouter() fiber.Router {
	return ah.router
}

// VendorAuth godoc
// @Summary Sign in or sign up via Apple or Google
// @Description Accepts token from vendor which we process and returning pair of tokens
// @Tags Authorization
// @Accept json
// @Produce json
// @Param vendor query string true "Vendor which is providing authorization" Enums(apple, google)
// @Success 200 {object} server.CommonResponse
// @Failure 400 {object} server.CommonResponse
// @Router /auth [post]
func (ah *AuthHandler) VendorAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logrus.Info("dsafasjfosdfsda")
		return c.SendString("sdjgajgwejgivqjweogvaoerhvqeirgaliergv")
	}
}
