package config

import (
	"bank_api/pkg/damqp/config"
	"bank_api/pkg/tconfig"
	thttpConfig "bank_api/pkg/thttp/config"
	tloggerConfig "bank_api/pkg/tlogger/config"
	tsecureConfig "bank_api/pkg/tsecure/config"
	tstorageConfig "bank_api/pkg/tstorage/config"
	"fmt"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	BaseConfig    tconfig.BaseConfig
	HttpConfig    thttpConfig.ThttpConfig
	LoggerConfig  tloggerConfig.TLoggerConfig
	SecureConfig  tsecureConfig.TSecureConfig
	StorageConfig tstorageConfig.TStorageConfig
	AmqpConfig    dconfig.BrokerConfig
}

func NewConfig() *Config {
	v, err := loadConfig()
	if err != nil {
		panic(fmt.Errorf("can't parse config: %s", err.Error()))
	}
	config, err := parseConfig(v)
	if err != nil {
		panic(fmt.Errorf("can't parse config: %s", err.Error()))
	}

	return config
}

func BuildConfig(ctn di.Container) (interface{}, error) {
	return NewConfig(), nil
}

func loadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
