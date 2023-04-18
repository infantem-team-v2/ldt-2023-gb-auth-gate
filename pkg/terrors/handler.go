package terrors

import (
	"bank_api/pkg/tlogger"
	"github.com/sarulabs/di"
)

type HttpErrorHandler struct {
	logger *tlogger.Logger `di:"logger"`
}

func BuildErrorHandler(ctn di.Container) (interface{}, error) {
	return &HttpErrorHandler{}, nil
}
