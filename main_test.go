package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filipecsoares/api-go-gin/controllers"
	"github.com/filipecsoares/api-go-gin/database"
	"github.com/filipecsoares/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID uint

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	routes.GET("/healthcheck", controllers.HealthCheck)
	routes.GET("/students", controllers.GetAllStudents)
	routes.GET("/students/email/:email", controllers.GetStudentByEmail)
	return routes
}

func CreateStudentMock() {
	student := models.Student{
		Name:  "Test",
		Email: "test@test.com",
	}
	database.DB.Create(&student)
	ID = student.ID
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
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

func TestListAllStudents(t *testing.T) {
	database.ConnectDataBase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByEmail(t *testing.T) {
	database.ConnectDataBase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	req, _ := http.NewRequest("GET", "/students/email/test@test.com", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}
