package config

type ThttpConfig struct {
	TimeOut       uint32 `json:"timeOut"`       // request timeout in seconds
	Accept        string `json:"accept"`        // what types of content accepts
	DoLogRequests uint8  `json:"doLogRequests"` // do logging of requests and responses
}

func (thc *ThttpConfig) ParseConfig() error {
	return nil
}
