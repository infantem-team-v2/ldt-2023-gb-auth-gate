package kafka

import (
	dconfig "bank_api/pkg/damqp/config"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Kafka struct {
	cfg    *dconfig.BrokerConfig
	client sarama.Client
}

func (k Kafka) InitKafka() error {
	config := sarama.NewConfig()
	brokers := []string{fmt.Sprintf("%s:%s", k.cfg.Kafka.Host, k.cfg.Kafka.Port)}

	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatalf("Cant't creating kafka client: %s", err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Fatalf("Cant't closing kafka client: %s", err)
		}
	}()
	return nil
}

func (k Kafka) Consume(topic string, groupID string) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumerGroupFromClient(groupID, k.client)
	if err != nil {
		log.Fatalf("Can't create consumer kafka: %s", err)
	}
	var handler sarama.ConsumerGroupHandler
	go func() {
		for {
			if err = consumer.Consume(context.Background(), []string{topic}, handler); err != nil {
				log.Fatalf("Error with runnig kafka consumer: %s", err)
			}
		}
	}()
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	if err = consumer.Close(); err != nil {
		log.Fatalf("Error with closing consumer kafka: %s", err)
	}
}

func (k Kafka) Publish(topic string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	producer, err := sarama.NewSyncProducerFromClient(k.client)
	if err != nil {
		log.Fatalf("Error with creating kafka producer: %s", err)
		return err
	}
	defer func() {
		if err != nil {
			log.Printf("Error with closing SyncProducer: %s", err)
		}
	}()
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error with send message kafka: %s", err)
		return err
	}
	log.Printf("Message sended in topic: %s, partition: %d, offset: %d", topic, partition, offset)
	return nil
}
