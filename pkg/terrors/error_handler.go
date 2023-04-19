package terrors

import (
	"bank_api/pkg/tlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type HttpErrorHandler struct {
	logger *tlogger.Logger `di:"logger"`
}

func BuildErrorHandler(ctn di.Container) (interface{}, error) {
	return &HttpErrorHandler{}, nil
}

func (heh *HttpErrorHandler) Handle(ctx *fiber.Ctx, err error) error {
	return nil
}
