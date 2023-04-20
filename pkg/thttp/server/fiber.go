package server

import (
	"bank_api/config"
	"bank_api/pkg/terrors"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
	"time"
)

func BuildFiberApp(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)
	errorHandler := ctn.Get("errorHandler").(*terrors.HttpErrorHandler)
	return fiber.New(fiber.Config{
		AppName:               cfg.BaseConfig.Service.Name,
		DisableStartupMessage: false,
		Prefork:               true,
		ErrorHandler:          errorHandler.Handle,
		WriteTimeout:          time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
		ReadTimeout:           time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
	}), nil
}
