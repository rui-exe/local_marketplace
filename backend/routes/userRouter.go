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
	userRoutes.GET("/selling_items/:username", controllers.GetSellingItems())
	userRoutes.GET("/me", controllers.GetMe())
	logoutRoutes := incomingRoutes.Group("/logout")
	logoutRoutes.Use(middleware.Authenticate())
	logoutRoutes.POST("/", controllers.Logout())
}
