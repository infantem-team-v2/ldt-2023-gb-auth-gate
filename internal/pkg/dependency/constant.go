package dependency

import (
	"bank_api/config"
	"bank_api/internal/auth/usecase"
	mdwHttp "bank_api/internal/pkg/middleware/delivery/http"
	"bank_api/pkg/damqp/kafka"
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
	"middleware":        mdwHttp.BuildMiddlewareManager,
	"errorHandler":      terrors.BuildErrorHandler,
	"stacktraceHandler": terrors.BuildStacktraceHandler,
	"app":               server.BuildFiberApp,
	"authUC":            usecase.BuildAuthUsecase,
	"rabbit":            rabbit.BuildRabbitMQ,
	"kafka":             kafka.BuildKafka,
}

const TagDI = "di"
