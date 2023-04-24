package damqp

import (
	"github.com/Shopify/sarama"
)

type Broker interface {
	StartingConsumeMQ(queueName string, consumeFunc func(d []byte) error) error
	PublishMQ(queueName, data, urlMQ string) error
	Publish(client sarama.Client, topic string, message ...string) error
	Consume(client sarama.Client, topic string, groupID ...string)
}
