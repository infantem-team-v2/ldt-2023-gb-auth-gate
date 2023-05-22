package dependency

import (
	"gb-auth-gate/config"
	authRepo "gb-auth-gate/internal/auth/repository"
	authUC "gb-auth-gate/internal/auth/usecase"
	mdwHttp "gb-auth-gate/internal/pkg/middleware/delivery/http"
	"gb-auth-gate/pkg/damqp/kafka"
	"gb-auth-gate/pkg/damqp/rabbit"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp"
	"gb-auth-gate/pkg/thttp/server"
	"gb-auth-gate/pkg/tlogger"
	tstorageCache "gb-auth-gate/pkg/tstorage/cache"
	tstorageRelational "gb-auth-gate/pkg/tstorage/relational"
	"github.com/sarulabs/di"
)

var dependencyMap = map[string]func(ctn di.Container) (interface{}, error){
	"config": config.BuildConfig,

	"postgres": tstorageRelational.BuildPostgres,
	"redis":    tstorageCache.BuildRedis,

	"httpClient": thttp.BuildHttpClient,

	"logger": tlogger.BuildLogger,

	"rabbit": rabbit.BuildRabbitMQ,
	"kafka":  kafka.BuildKafka,

	"middleware":        mdwHttp.BuildMiddlewareManager,
	"errorHandler":      terrors.BuildErrorHandler,
	"stacktraceHandler": terrors.BuildStacktraceHandler,

	"app": server.BuildFiberApp,

	"authUC":   authUC.BuildAuthUsecase,
	"authRepo": authRepo.BuildPostgresRepository,
}

const TagDI = "di"
