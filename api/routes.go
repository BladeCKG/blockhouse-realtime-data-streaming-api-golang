package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.POST("/stream/start", StartStreamHandler)
	router.POST("/stream/:stream_id/send", SendDataHandler)
	router.GET("/stream/:stream_id/results", StreamResultsHandler)
}
