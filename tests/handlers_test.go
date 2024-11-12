package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BladeCKG/blockhouse-realtime-data-streaming-api-golang/api"
	"github.com/gin-gonic/gin"
)

func TestStartStreamHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/stream/start", api.StartStreamHandler)

	req, _ := http.NewRequest("POST", "/stream/start", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.Code)
	}
}
