package routes

import (
	"github.com/filipecsoares/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}
