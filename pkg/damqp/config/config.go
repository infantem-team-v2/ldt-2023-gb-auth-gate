package dconfig

type AmqpConfig struct {
	RabbitMQ struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	} `json:"RabbitMQ"`
}
