package http

import (
	authInterface "bank_api/internal/auth/interface"
	"bank_api/internal/auth/model"
	"bank_api/pkg/terrors"
	"bank_api/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthUC authInterface.UseCase `di:"authUC"`
	prefix string
	router fiber.Router
}

func (ah *AuthHandler) GetRouter() fiber.Router {
	return ah.router
}

func (ah *AuthHandler) GetPrefix() string {
	return ah.prefix
}

func NewAuthHandler(app *fiber.App) server.IHandler {
	return &AuthHandler{
		router: app.Group("auth"),
		prefix: "auth",
	}
}

// VendorAuth godoc
// @Summary Sign in or sign up via Apple or Google
// @Description Accepts token from vendor which we process and returning pair of tokens
// @Tags Authorization
// @Accept json
// @Produce json
// @Param vendor query string true "Vendor which is providing authorization" Enums(apple, google)
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Router /auth [post]
func (ah *AuthHandler) VendorAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return terrors.Raise(nil, 300000)
	}
}

// SignUp godoc
// @Summary Sign up with email
// @Description Sign up with email and password
// @Tags Authorization
// @Accept json
// @Produce json
// @Param data body model.SignUpRequest true "Authorization data from user"
// @Success 200 {object} model.SignUpResponse
// @Failure 404 {object} common.Response
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
// @Param data body model.SignInRequest true "Authorization data from user"
// @Success 200 {object} model.SignInResponse
// @Failure 404 {object} common.Response
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
// @Param code body model.EmailValidateRequest true "Data for validation by email from app"
// @Success 200 {object} model.EmailValidateResponse
// @Failure 404 {object} common.Response
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
