package config

type TLoggerConfig struct {
	Loki struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"Loki"`
}

func (tlc *TLoggerConfig) ParseConfig() error {
	return nil
}
