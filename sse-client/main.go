package main

import (
	"fmt"
	"log"
	"time"

	"github.com/r3labs/sse/v2"
)

func main() {
	// Create a new SSE client
	topic := "c807f630-2a27-433a-b2f5-fdc455a9fa77"
	endpoint := "http://localhost:8080/stream/" + topic + "/results"
	client := sse.NewClient(endpoint)

	// Subscribe to messages
	err := client.Subscribe("messages", func(msg *sse.Event) {
		fmt.Printf("%s Received message: %s\n", time.Now(), msg.Data)
		// Handle the received data
	})
	if err != nil {
		log.Fatalf("Error subscribing to SSE: %v", err)
	}
}
