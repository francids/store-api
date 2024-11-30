package controllers

import (
	"net/http"
	"store-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
