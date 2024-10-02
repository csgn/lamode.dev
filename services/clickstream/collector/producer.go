package collector

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	logger *log.Logger
	writer *kafka.Producer
	topic  string
}

func NewProducer(
	logger *log.Logger,
	host string,
	port string,
	topic string,
) (*Producer, error) {
	bootstrapServers := fmt.Sprintf("%s:%s", host, port)

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": bootstrapServers,
		},
	)

	if err != nil {
		return nil, err
	}

	newProducer := &Producer{
		logger: logger,
		writer: p,
		topic:  topic,
	}

	return newProducer, nil
}

func (self *Producer) Send(payload []byte) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &self.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	}

	return self.writer.Produce(message, nil)
}
