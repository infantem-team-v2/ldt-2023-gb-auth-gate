package dependency

import (
	"bank_api/config"
	"bank_api/internal/auth/usecase"
	"bank_api/pkg/damqp/rabbit"
	"bank_api/pkg/terrors"
	"bank_api/pkg/thttp"
	"bank_api/pkg/thttp/server"
	"bank_api/pkg/tlogger"
	tstorageRelational "bank_api/pkg/tstorage/relational"
	"github.com/sarulabs/di"
)

var dependencyMap = map[string]func(ctn di.Container) (interface{}, error){
	"config":            config.BuildConfig,
	"postgres":          tstorageRelational.BuildPostgres,
	"httpClient":        thttp.BuildHttpClient,
	"logger":            tlogger.BuildLogger,
	"errorHandler":      terrors.BuildErrorHandler,
	"stacktraceHandler": terrors.BuildStacktraceHandler,
	"app":               server.BuildFiberApp,
	"authUC":            usecase.BuildAuthUsecase,
	"rabbit":            rabbit.BuildRabbitMQ,
}

const TagDI = "di"
