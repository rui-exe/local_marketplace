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
	productRoutes.GET("/:product_id", controllers.GetProduct())
	productRoutes.PUT("/:product_id", controllers.UpdateProduct())
	productRoutes.POST("/upload/picture/:product_id", controllers.UploadProductPicture())
	productRoutes.POST("/wishlist/:product_id", controllers.AddProductToWishlist())
	productRoutes.DELETE("/wishlist/:product_id", controllers.RemoveProductFromWishlist())
}
