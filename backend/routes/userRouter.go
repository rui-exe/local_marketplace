package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/users")
	userRoutes.Use(middleware.Authenticate())
	userRoutes.GET("/:username", controllers.GetUser())
	userRoutes.POST("/upload/picture", controllers.UploadPicture())
	userRoutes.GET("/wishlist/:username", controllers.GetWishlist())
}
