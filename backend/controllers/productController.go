package controllers

import (
	"backend/database"
	"backend/firebase"
	"backend/models"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func isUserOwnerOfProduct(uid string, product_id primitive.ObjectID) bool {
	var product models.Product
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := productCollection.FindOne(ctx, bson.M{"_id": product_id}).Decode(&product)
	if err != nil {
		return false
	}
	if product.Seller_id.Hex() == uid {
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

func UploadProductPicture() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var product models.Product
		id, err := primitive.ObjectIDFromHex(c.Param("product_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id format"})
			return
		}

		err = productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		uid := c.GetString("uid")

		if !isUserOwnerOfProduct(uid, id) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to access this resource, you are not the owner of this product"})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while uploading file"})
			return
		}

		fileContent, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while opening file"})
			return
		}
		defer fileContent.Close()
		fileName := fmt.Sprintf("product_pictures/%s", id.Hex())

		client, err := firebase.App.Storage(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while connecting to firebase"})
			return
		}

		bucket, err := client.Bucket("local-marketplace-fde45.appspot.com")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while initializing Firebase bucket"})
			return
		}

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, fileContent); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while copying the file"})
			return
		}

		object := bucket.Object(fileName)
		wc := object.NewWriter(ctx)
		if _, err = wc.Write(buf.Bytes()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while copying the file"})
			return
		}

		if err := wc.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while closing the writer"})
			return
		}

		imageURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/local-marketplace-fde45.appspot.com/o/%s?alt=media", url.PathEscape(fileName))

		_, err = productCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"picture": imageURL}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating product picture"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	}

}
