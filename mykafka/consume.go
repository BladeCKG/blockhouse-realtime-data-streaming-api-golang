package mykafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	Consumer *kafka.Consumer
}

func NewConsumer(broker, groupID string) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	return &Consumer{Consumer: consumer}, nil
}

func (c *Consumer) Consume(streamID string, resultChan chan<- []byte) {
	topic := "stream-" + streamID
	err := c.Consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic %s: %v", topic, err)
	}

	for {
		msg, err := c.Consumer.ReadMessage(-1)
		if err == nil {
			resultChan <- msg.Value
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
