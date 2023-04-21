package server

import (
	"bank_api/internal/auth/delivery/http"
	"bank_api/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

var (
	emptyHandlers = map[string]func(app *fiber.App) server.IHandler{
		"auth": http.NewAuthHandler,
	}
)
