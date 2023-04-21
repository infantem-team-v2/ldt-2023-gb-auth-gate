package damqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func Consume(ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	log.Panicf("Error with register consumer: %s", err)

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
		}
	}()
}
