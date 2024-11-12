package main

import (
	"log"

	"github.com/BladeCKG/blockhouse-realtime-data-streaming-api-golang/api"
	"github.com/BladeCKG/blockhouse-realtime-data-streaming-api-golang/mykafka"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := mykafka.InitProducer("localhost:9092"); err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}

	router := gin.Default()
	api.SetupRoutes(router)

	log.Println("Starting server on :8080...")
	router.Run(":8080")
}
