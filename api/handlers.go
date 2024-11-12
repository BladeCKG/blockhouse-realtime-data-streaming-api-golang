package api

import (
	"net/http"

	"github.com/BladeCKG/blockhouse-realtime-data-streaming-api-golang/mykafka"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create new Kafka stream with unique topic
func StartStreamHandler(c *gin.Context) {
	streamID := uuid.New().String()

	topic := "stream-" + streamID
	// Create a new Kafka topic for the stream
	err := mykafka.CreateTopic("localhost:9092", topic, 1, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Kafka topic"})
		return
	}

	// Send the stream ID to the client
	c.JSON(http.StatusOK, gin.H{"stream_id": streamID})
}

// Send data to Kafka
func SendDataHandler(c *gin.Context) {
	streamID := c.Param("stream_id")
	data, _ := c.GetRawData()
	if err := mykafka.SendMessage(streamID, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send data to Kafka"})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status": "Data sent successfully"})
	}
}

// Stream results to the client with SSE
func StreamResultsHandler(c *gin.Context) {
	streamID := c.Param("stream_id")
	resultChan := make(chan []byte)

	// Kafka consumer for the given streamID
	consumer, err := mykafka.NewConsumer("localhost:9092", streamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Kafka consumer"})
		return
	}

	go consumer.Consume(streamID, resultChan)

	// Set SSE headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// Stream messages to the client in SSE format
	for msg := range resultChan {
		// fmt.Println("Sending message: ", string(msg))
		c.Writer.Write([]byte("data: " + string(msg) + "\n\n"))
		c.Writer.Flush()
	}
}
