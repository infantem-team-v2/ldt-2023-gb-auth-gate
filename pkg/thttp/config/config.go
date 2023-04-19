package config

type ThttpConfig struct {
	TimeOut       uint32 `mapstructure:"time_out"`        // request timeout in seconds
	Accept        string `mapstructure:"accept"`          // what types of content accepts
	DoLogRequests uint8  `mapstructure:"do_log_requests"` // do logging of requests and responses
}

func (thc *ThttpConfig) ParseConfig() error {
	return nil
}
