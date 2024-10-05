package main

import (
	"errors"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
)

type AsyncProducer kafka.Producer

func NewProducer(addr string) *AsyncProducer {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": addr,
		},
	)

	if err != nil {
		log.Fatalln("Failed to start Kafka producer: ", err)
	}

	if *verbose {
		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Error:
					log.Printf("Kafka error: %v\n", ev)
					if ev.IsFatal() {
						log.Fatalf("Fatal Kafka error: %v\n", ev)
					}
				default:
					log.Printf("Kafka event: %v\n", ev)
				}
			}
		}()
	}

	return (*AsyncProducer)(p)
}

func (p *AsyncProducer) SendMessage(payload []byte, topic string) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(uuid.New().String()),
		Value: payload,
	}

	return (*kafka.Producer)(p).Produce(msg, nil)
}

func (p *AsyncProducer) Shutdown() error {
	var pp = (*kafka.Producer)(p)
	pp.Close()

	if !pp.IsClosed() {
		return errors.New("Failed to close the producer cleanly")
	}

	return nil
}
