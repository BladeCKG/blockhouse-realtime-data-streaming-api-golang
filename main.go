package main

import (
	"log"

	"github.com/BladeCKG/blockhouse-realtime-data-streaming-api-golang/api"
	"github.com/BladeCKG/blockhouse-realtime-data-streaming-api-golang/config"
	"github.com/BladeCKG/blockhouse-realtime-data-streaming-api-golang/mykafka"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig() // Initialize the global AppConfig

	if err := mykafka.InitProducer(config.AppConfig.Broker); err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}

	router := gin.Default()
	// router.Use(middleware.APIKeyAuthMiddleware(config.AppConfig.ApiKey))
	api.SetupRoutes(router)

	log.Println("Starting server on :8080...")
	router.Run(":8080")
}
