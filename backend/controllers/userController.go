package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, "password is incorrect"
	}
	return true, ""

}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := c.BindJSON(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate user input
		err = validate.Struct(user)

		//check if user type is admin
		if *user.User_type == "ADMIN" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user type cannot be admin"})
			return
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		countEmail, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking user email"})
		}

		countPhone, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking user phone"})
		}

		if countEmail > 0 || countPhone > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or phone number already exists"})
			return
		}

		hashedPassword, err := HashPassword(*user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while hashing password"})
			return
		}
		user.Password = &hashedPassword

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, err := helpers.GenerateAllTokens(*user.Email, *user.Username, *user.User_type, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while generating token"})
			return
		}
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while inserting user"})
			return
		}

		c.JSON(http.StatusOK, resultInsertionNumber)

	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		err := c.BindJSON(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordMatch, message := VerifyPassword(*foundUser.Password, *user.Password)
		if !passwordMatch {
			c.JSON(http.StatusInternalServerError, gin.H{"error": message})
			return
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		token, refreshToken, err := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.Username, *foundUser.User_type, foundUser.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while generating token"})
			return
		}
		updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_, err = userCollection.UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": bson.M{"token": token, "refresh_token": refreshToken, "updated_at": updated_at}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating user token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
	}

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		// Create a context with a timeout
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel() // Make sure to defer the cancel right after creating the context

		// Initialize a user model to store the result
		var user models.UserDisplay

		// Check if the user exists
		err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// If no error is found, return the user
		c.JSON(http.StatusOK, user)
	}
}
