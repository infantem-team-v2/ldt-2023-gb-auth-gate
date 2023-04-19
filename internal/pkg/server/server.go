package server

import (
	"bank_api/config"
	"bank_api/internal/pkg/dependency"
	"bank_api/pkg/terrors"
	"bank_api/pkg/thttp"
	"bank_api/pkg/tlogger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	mwLogger "github.com/gofiber/fiber/v2/middleware/logger"
	mwRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"time"
)

type Server struct {
	Config     *config.Config     `di:"config"`
	Logger     tlogger.ILogger    `di:"logger"`
	HttpClient *thttp.ThttpClient `di:"httpClient"`
	container  *dependency.TDependencyContainer
	app        *fiber.App
}

// NewServer (Fabric) Builds server with DI container which contains main pkg Singletons,
// which we use to build other entities
func NewServer() *Server {
	ctn := dependency.NewDependencyContainer().BuildDependencies().BuildContainer()

	cfg := ctn.ContainerInstance().Get("config").(*config.Config)
	errorHandler := ctn.ContainerInstance().Get("errorHandler").(*terrors.HttpErrorHandler)

	s := &Server{
		container: ctn,
		app: fiber.New(fiber.Config{
			AppName:      cfg.BaseConfig.Service.Name,
			ErrorHandler: errorHandler.Handle,
			WriteTimeout: time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
			ReadTimeout:  time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
		}),
		Config: nil,
		Logger: nil,
	}
	ctn.Inject(s)
	return s
}

// MapHandlers with middlewares and routers
func (s *Server) MapHandlers() *Server {

	sh := s.container.ContainerInstance().Get("stacktraceHandler").(*terrors.StacktraceHandler)
	s.app.Use(mwRecover.New(mwRecover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: sh.Handle,
	}))

	s.app.Use(mwLogger.New(mwLogger.Config{}))

	return s
}

// Run app on tconfig.BaseConfig.System.Port
func (s *Server) Run() error {
	return s.app.Listen(fmt.Sprintf("%s:%s", s.Config.BaseConfig.System.Host, s.Config.BaseConfig.System.Port))
}
