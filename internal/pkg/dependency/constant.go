package dependency

import (
	"bank_api/config"
	auth_usecase "bank_api/internal/middleware/usecase"
	"bank_api/pkg/terrors"
	"bank_api/pkg/thttp"
	"bank_api/pkg/thttp/server"
	"bank_api/pkg/tlogger"
	"github.com/sarulabs/di"
)

var dependencyMap = map[string]func(ctn di.Container) (interface{}, error){
	"config":            config.BuildConfig,
	"httpClient":        thttp.BuildHttpClient,
	"logger":            tlogger.BuildLogger,
	"errorHandler":      terrors.BuildErrorHandler,
	"stacktraceHandler": terrors.BuildStacktraceHandler,
	"app":               server.BuildFiberApp,
	"authUC":            auth_usecase.BuildAuthUsecase,
}

const TagDI = "di"
