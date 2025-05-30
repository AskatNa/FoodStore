package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	Stock       int32              `bson:"stock"`
	Category    string             `bson:"category"`
	Images      []string           `bson:"images"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
