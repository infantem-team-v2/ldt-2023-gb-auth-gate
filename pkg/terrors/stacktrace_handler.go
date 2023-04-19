package terrors

import (
	"bank_api/pkg/tlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type StacktraceHandler struct {
	logger tlogger.ILogger
}

func BuildStacktraceHandler(ctn di.Container) (interface{}, error) {
	return &StacktraceHandler{
		logger: ctn.Get("logger").(tlogger.ILogger),
	}, nil
}

func (sh *StacktraceHandler) Handle(c *fiber.Ctx, e interface{}) {

	return
}
