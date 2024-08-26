package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	productRoutes := incomingRoutes.Group("/products")
	productRoutes.Use(middleware.Authenticate())
	productRoutes.POST("/", controllers.CreateProduct())
}
