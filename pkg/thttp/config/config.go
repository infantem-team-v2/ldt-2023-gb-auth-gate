package config

type ThttpConfig struct {
	TimeOut       uint32 `json:"time_out"`        // request timeout in seconds
	Accept        string `json:"accept"`          // what types of content accepts
	DoLogRequests uint8  `json:"do_log_requests"` // do logging of requests and responses
}

func (thc *ThttpConfig) ParseConfig() error {
	return nil
}
