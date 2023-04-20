package server

import (
	"bank_api/config"
	"bank_api/internal/auth/delivery/http"
	"bank_api/internal/pkg/dependency"
	"bank_api/pkg/terrors"
	"bank_api/pkg/tlogger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	mwLogger "github.com/gofiber/fiber/v2/middleware/logger"
	mwRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

type Server struct {
	container *dependency.TDependencyContainer
	App       *fiber.App      `di:"app"`
	Config    *config.Config  `di:"config"`
	Logger    tlogger.ILogger `di:"logger"`
}

// NewServer (Fabric) Builds server with DI container which contains main pkg Singletons,
// which we use to build other entities
func NewServer() *Server {
	ctn := dependency.NewDependencyContainer().BuildDependencies().BuildContainer()
	s := &Server{
		container: ctn,
	}
	ctn.Inject(s)
	return s
}

// MapHandlers with middlewares and routers
func (s *Server) MapHandlers() *Server {
	sh := s.container.ContainerInstance().Get("stacktraceHandler").(*terrors.StacktraceHandler)
	s.App.Use(mwRecover.New(mwRecover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: sh.Handle,
	}))
	s.App.Use(mwLogger.New(mwLogger.Config{
		Output: os.Stdout,
	}))
	router := s.App.Group("auth")
	ah := &http.AuthHandler{
		Router: router,
	}
	s.container.Inject(ah)
	ah.SetRoutes()
	return s
}

// Run app on tconfig.BaseConfig.System.Port
func (s *Server) Run() error {
	s.Logger.Infof("STARTED SERVER")
	return s.App.Listen(fmt.Sprintf("%s:%s", s.Config.BaseConfig.System.Host, s.Config.BaseConfig.System.Port))
}
