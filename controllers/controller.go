package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "Filipe",
	})
}
