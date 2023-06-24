package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	routes.GET("/students/:id", controllers.GetStudentById)
	routes.GET("/students/email/:email", controllers.GetStudentByEmail)
	routes.DELETE("/students/:id", controllers.DeleteStudent)
	routes.PATCH("/students/:id", controllers.EditStudent)
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

func TestGetStudentById(t *testing.T) {
	database.ConnectDataBase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	path := fmt.Sprintf("/students/%d", ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var student models.Student
	json.Unmarshal(response.Body.Bytes(), &student)
	assert.Equal(t, ID, student.ID)
	assert.Equal(t, "Test", student.Name)
	assert.Equal(t, "test@test.com", student.Email)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDataBase()
	CreateStudentMock()
	r := SetupTestRoutes()
	path := fmt.Sprintf("/students/%d", ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditStudent(t *testing.T) {
	database.ConnectDataBase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	studentNewData := models.Student{
		Name:  "Test2",
		Email: "test2@test.com",
	}
	jsonStudent, _ := json.Marshal(studentNewData)
	path := fmt.Sprintf("/students/%d", ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonStudent))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var student models.Student
	json.Unmarshal(response.Body.Bytes(), &student)
	assert.Equal(t, ID, student.ID)
	assert.Equal(t, "Test2", student.Name)
	assert.Equal(t, "test2@test.com", student.Email)
	assert.Equal(t, http.StatusOK, response.Code)
}
