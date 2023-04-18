package config

type TStorageConfig struct {
	Relational struct {
		Postgres struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			User     string `json:"user"`
			Password string `json:"password"`
			DBName   string `json:"db_name"`
			SSLMode  string `json:"ssl_mode"`
			PgDriver string `json:"pg_driver"`
		} `json:"Postgres"`
	} `json:"Relational"`
	Cache struct {
		Redis struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			DB       string `json:"db"`
			Password string `json:"password"`
		} `json:"Redis"`
	} `json:"Cache"`
	NonRelational struct {
	}
}
