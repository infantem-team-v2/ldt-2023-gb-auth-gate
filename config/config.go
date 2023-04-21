package config

import (
	"bank_api/pkg/tconfig"
	thttpConfig "bank_api/pkg/thttp/config"
	tloggerConfig "bank_api/pkg/tlogger/config"
	tsecureConfig "bank_api/pkg/tsecure/config"
	"fmt"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	BaseConfig   tconfig.BaseConfig
	HttpConfig   thttpConfig.ThttpConfig
	LoggerConfig tloggerConfig.TLoggerConfig
	SecureConfig tsecureConfig.TSecureConfig
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

//func (c *Config) ParseConfig() error {
//	return nil
//}

func loadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("/home/dima/go/src/bank_api/config") //TEST
	//v.AddConfigPath("config") //PROD
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
