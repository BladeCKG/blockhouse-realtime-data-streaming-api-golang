package mykafka

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// CreateTopic creates a new Kafka topic with the given name and number of partitions.
func CreateTopic(broker, topic string, numPartitions, replicationFactor int) error {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return fmt.Errorf("failed to create Kafka admin client: %w", err)
	}
	defer adminClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Define topic configuration
	specs := []kafka.TopicSpecification{{
		Topic:             topic,
		NumPartitions:     numPartitions,
		ReplicationFactor: replicationFactor,
	}}

	// Attempt to create the topic
	results, err := adminClient.CreateTopics(ctx, specs)
	if err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			return fmt.Errorf("failed to create topic %s: %s", result.Topic, result.Error.String())
		}
	}

	return nil
}
