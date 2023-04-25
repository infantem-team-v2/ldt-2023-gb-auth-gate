package http

import (
	"bank_api/config"
	"bank_api/internal/pkg/middleware/model"
	"bank_api/pkg/tlogger"
	"bank_api/pkg/tutils/etc"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v2"
	"github.com/sarulabs/di"
	"time"
)

type MiddlewareManager struct {
	sessionStorage *session.Store
	Logger         tlogger.ILogger `di:"logger"`
}

func BuildMiddlewareManager(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)
	logger := ctn.Get("logger").(tlogger.ILogger)
	s := session.New(session.Config{
		CookieHTTPOnly: true,
		Storage: redis.New(redis.Config{
			Reset:    false,
			Host:     cfg.StorageConfig.Cache.Redis.Host,
			Port:     etc.MustParseToInt(cfg.StorageConfig.Cache.Redis.Port),
			Password: cfg.StorageConfig.Cache.Redis.Password,
			Database: etc.MustParseToInt(cfg.StorageConfig.Cache.Redis.DB),
		}),
		Expiration:   time.Duration(cfg.MiddlewareConfig.Auth.ExpirationTime) * time.Minute,
		KeyLookup:    fmt.Sprintf("cookie:%s", model.SessionKey),
		CookieDomain: cfg.BaseConfig.Service.URL,
		CookieSecure: true,
	})

	return &MiddlewareManager{
		sessionStorage: s,
		Logger:         logger,
	}, nil
}
