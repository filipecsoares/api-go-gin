package routes

import (
	"github.com/filipecsoares/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/healthcheck", controllers.HealthCheck)
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("/students/email/:email", controllers.GetStudentByEmail)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.Run()
}
