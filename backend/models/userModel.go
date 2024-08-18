package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User represents a user in the database
type User struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
}

// GetUser retrieves a user from the database by username
func GetUser(db *mongo.Database, username string) (*User, error) {
	// Get the users collection
	users := db.Collection("users")

	// Find the user by username
	var user User
	err := users.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
