package auth_usecase

import (
	"bank_api/pkg/tlogger"
	"github.com/sarulabs/di"
)

type AuthUS struct {
	logger tlogger.ILogger
}

func BuildAuthUsecase(ctn di.Container) (interface{}, error) {
	return &AuthUS{
		logger: ctn.Get("logger").(tlogger.ILogger),
	}, nil
}
