package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the database
type User struct {
	ID            primitive.ObjectID   `bson:"_id"`
	Username      *string              `json:"username" validate:"required,min=4,max=25"`
	Password      *string              `json:"password" validate:"required,min=6"`
	Email         *string              `json:"email" validate:"required,email"`
	Phone         *string              `json:"phone" validate:"required"`
	Role          *string              `json:"role" validate:"required,eq=SELLER|eq=BUYER"`
	Picture       *string              `json:"picture"`
	User_type     *string              `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Wishlist      []primitive.ObjectID `json:"wishlist"`
	Selling_items []primitive.ObjectID `json:"selling_items"`
	Created_at    time.Time            `json:"created_at"`
	Updated_at    time.Time            `json:"updated_at"`
	Token         *string              `json:"token"`
	Refresh_token *string              `json:"refresh_token"`
	User_id       string               `json:"user_id"`
}

type UserDisplay struct {
	ID            primitive.ObjectID `bson:"_id"`
	Username      *string            `json:"username"`
	Email         *string            `json:"email"`
	Phone         *string            `json:"phone"`
	Role          *string            `json:"role"`
	Picture       *string            `json:"picture"`
	Created_at    time.Time          `json:"created_at"`
	Wishlist      []primitive.ObjectID
	Selling_items []primitive.ObjectID
}
