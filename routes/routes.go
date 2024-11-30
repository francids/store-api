package routes

import (
	"store-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.GET("/ping", controllers.Ping)
	r.GET("/products", controllers.GetProducts)

	return r
}
