package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(c *gin.Context, db *mongo.Database) {
	// Get user username from URL
	username := c.Param("username")

	// Get user from database
	user, err := models.GetUser(db, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return user
	c.JSON(http.StatusOK, user)
}
