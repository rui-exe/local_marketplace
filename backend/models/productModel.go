package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" validate:"required,min=3,max=100"`
	Description string             `json:"description" validate:"required,min=10"`
	Price       float64            `json:"price" validate:"required,gt=0"`
	Category    string             `json:"category" validate:"required"`
	SellerID    primitive.ObjectID `json:"seller_id" validate:"required"`
	Picture     *string            `json:"picture"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Status      string             `json:"status" validate:"required,oneof=available sold"`
}
