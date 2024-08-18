package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, db *mongo.Database) {
	// Define the routes for the application
	router.GET("/user/:username", func(c *gin.Context) { controllers.GetUser(c, db) })

}
