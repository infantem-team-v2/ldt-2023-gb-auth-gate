package terrors

import (
	"bank_api/internal/pkg/common"
	"bank_api/pkg/tlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
	"strings"
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
		return ctx.
			Status(tErr.statusCode).
			JSON(tErr.externalMessage)
	}
	if strings.Contains(err.Error(), "Cannot") {
		return ctx.
			Status(404).
			JSON(common.Response{
				InternalCode: 404,
				Message:      err.Error(),
			})
	}
	return ctx.
		Status(500).
		JSON(common.Response{
			InternalCode: 500,
			Message:      err.Error(),
		})
}
