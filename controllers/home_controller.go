package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"greeting":    "Welcome to store-api",
		"description": "This is a simple store-api built with Go and Gin",
	})
}
