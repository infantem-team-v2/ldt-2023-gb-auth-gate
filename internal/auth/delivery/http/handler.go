package http

import (
	authInterface "bank_api/internal/auth/interface"
	"bank_api/internal/auth/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	AuthUC authInterface.UseCase `di:"authUC"`
	Router fiber.Router
}

func (ah *AuthHandler) GetRouter() fiber.Router {
	return ah.Router
}

// VendorAuth godoc
// @Summary Sign in or sign up via Apple or Google
// @Description Accepts token from vendor which we process and returning pair of tokens
// @Tags Authorization
// @Accept json
// @Produce json
// @Param vendor query string true "Vendor which is providing authorization" Enums(apple, google)
// @Success 200 {object}
// @Failure 400 {object}
// @Router /auth
func (ah *AuthHandler) VendorAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logrus.Info("dsafasjfosdfsda")
		return ctx.SendString("sdjgajgwejgivqjweogvaoerhvqeirgaliergv")
	}
}

func (ah *AuthHandler) SignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.SignUpRequest
		if err := ctx.BodyParser(&params); err != nil {
			return err
		}
		response, err := ah.AuthUC.SignUp(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}

func (ah *AuthHandler) ValidateEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.EmailValidateRequest
		if err := ctx.BodyParser(&params); err != nil {
			return err
		}
		response, err := ah.AuthUC.ValidateEmail(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}
