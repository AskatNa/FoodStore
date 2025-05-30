package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ProductID   primitive.ObjectID `bson:"product_id"`
	ProductName string             `bson:"product_name"`
	Quantity    int32              `bson:"quantity"`
	UnitPrice   float64            `bson:"unit_price"`
	Subtotal    float64            `bson:"subtotal"`
}

type Order struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID      primitive.ObjectID `bson:"customer_id"`
	Items           []OrderItem        `bson:"items"`
	TotalAmount     float64            `bson:"total_amount"`
	Status          string             `bson:"status"`
	ShippingAddress string             `bson:"shipping_address"`
	PaymentMethod   string             `bson:"payment_method"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}
