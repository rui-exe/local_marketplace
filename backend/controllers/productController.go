package controllers

import (
	"backend/database"
	"backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var productCollection = database.OpenCollection(database.Client, "product")
var productValidate = validator.New()

func isUserSeller(uid string) bool {
	var user models.User
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, bson.M{"user_id": uid}).Decode(&user)
	if err != nil {
		return false
	}
	if *user.Role == "SELLER" {
		return true
	}
	return false
}

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var product models.Product
		err := c.BindJSON(&product)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = productValidate.Struct(product)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if product.Seller_id.Hex() != c.GetString("uid") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to access this resource"})
			return
		}

		if !isUserSeller(product.Seller_id.Hex()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not a seller"})
			return
		}

		product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.ID = primitive.NewObjectID()

		result, insertErr := productCollection.InsertOne(ctx, product)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product created", "id": result.InsertedID})

	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var product models.Product
		id, err := primitive.ObjectIDFromHex(c.Param("product_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
			return
		}

		err = productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching product"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}
