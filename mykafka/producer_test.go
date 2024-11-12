package mykafka

import (
	"testing"
)

func TestSendMessage(t *testing.T) {
	err := InitProducer("localhost:9092")
	if err != nil {
		t.Fatalf("Failed to initialize Kafka producer: %v", err)
	}

	streamID := "c807f630-2a27-433a-b2f5-fdc455a9fa77"
	data := []byte("test message")

	if err := SendMessage(streamID, data); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}
}
