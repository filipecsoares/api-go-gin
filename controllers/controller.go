package controllers

import (
	"net/http"

	"github.com/filipecsoares/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Students)
}
