package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filipecsoares/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, http.StatusOK, response.Code)
	responseMock := `{"message":"OK"}`
	assert.Equal(t, responseMock, response.Body.String())
}
