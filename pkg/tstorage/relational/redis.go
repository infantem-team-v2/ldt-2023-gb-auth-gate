package relational

import (
	"bank_api/pkg/tstorage/config"
	"fmt"
	"github.com/fiorix/go-redis/redis"
)

func InitRedis(cfg *config.TStorageConfig) (*redis.Client, error) {
	url := fmt.Sprintf(
		"%s:%s db=%d passwd=%s",
		cfg.Cache.Redis.Host,
		cfg.Cache.Redis.Port,
		cfg.Cache.Redis.DB,
		cfg.Cache.Redis.Password)
	client, err := redis.NewClient(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}
