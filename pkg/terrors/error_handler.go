package terrors

import (
	"bank_api/pkg/tlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type HttpErrorHandler struct {
	logger tlogger.ILogger `di:"logger"`
}

func BuildErrorHandler(ctn di.Container) (interface{}, error) {
	return &HttpErrorHandler{
		logger: ctn.Get("logger").(tlogger.ILogger),
	}, nil
}

func (heh *HttpErrorHandler) Handle(ctx *fiber.Ctx, err error) error {
	if tErr, ok := err.(*tError); ok {
		heh.logger.ErrorFull(tErr)
		return ctx.Status(tErr.statusCode).JSON(tErr.externalMessage)
	}
	return nil
}
