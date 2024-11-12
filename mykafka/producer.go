package mykafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func InitProducer(broker string) error {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return err
	}
	return nil
}

func SendMessage(streamID string, data []byte) error {
	topic := "stream-" + streamID
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}

	if err := producer.Produce(message, nil); err != nil {
		log.Printf("Failed to send message to %s: %v", topic, err)
		return err
	}
	return nil
}
