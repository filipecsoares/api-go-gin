package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filipecsoares/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	routes.GET("/healthcheck", controllers.HealthCheck)
	return routes
}

func TestStatusCodeHealthCheck(t *testing.T) {
	r := SetupTestRoutes()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", response.Code)
	}
}
