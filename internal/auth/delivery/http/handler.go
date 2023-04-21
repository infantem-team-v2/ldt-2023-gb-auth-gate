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

// SignUp godoc
// @Summary Sign up with email
// @Description Sign up with email and password
// @Tags Authorization
// @Accept json
// @Produce json
// @Param email, password, first_name, last_name
// @Success 200
// @Failure 404
// @Router /auth/sign/up [post]
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

// SignIn godoc
// @Summary Sign in
// @Description Authorization and get access token
// @Tags Authorization
// @Accept json
// @Produce json
// @Param email, password
// @Success 200
// @Failure 404
// @Router /auth/sign/in [post]
func (ah *AuthHandler) SignIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.SignInRequest
		if err := ctx.BodyParser(&params); err != nil {
			return err
		}
		response, err := ah.AuthUC.SignIn(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}

// ValidateEmail godoc
// @Summary Validating user's email
// @Description Validating user's email with take message on email and writing code
// @Tags Authorization
// @Accept json
// @Produce json
// @Param code
// @Success 200
// @Failure 404
// @Router /auth/email/validate [post]
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
