package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/users")
	userRoutes.Use(middleware.Authenticate())
	userRoutes.GET("/:user_id", controllers.GetUser())
}
