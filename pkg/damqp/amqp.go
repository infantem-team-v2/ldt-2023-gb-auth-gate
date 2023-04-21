package damqp

import (
	dconfig "bank_api/pkg/damqp/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(cfg *dconfig.AmqpConfig) (*amqp.Connection, error) {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@rabbitmq-%s:%s/",
		cfg.RabbitMQ.UserName,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		amqp.Logger.Printf("Error with connection to RabbitMQ", err)
		return nil, err
	}
	return conn, nil
}
