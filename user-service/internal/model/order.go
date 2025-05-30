package model

import "time"

const (
	OrderStatusPending   = "pending"
	OrderStatusPaid      = "paid"
	OrderStatusDelivered = "delivered"
	OrderStatusCancelled = "cancelled"
)

type Order struct {
	ID          uint64
	CustomerID  uint64
	Status      string
	TotalAmount float64
	Items       []OrderItem
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
}

type OrderItem struct {
	ID        uint64
	OrderID   uint64
	ProductID uint64
	Quantity  int
	Price     float64
	Subtotal  float64
}

type OrderFilter struct {
	ID         *uint64
	CustomerID *uint64
	Status     *string
	FromDate   *time.Time
	ToDate     *time.Time
	IsDeleted  *bool
}

type OrderUpdateData struct {
	Status    *string
	UpdatedAt *time.Time
	IsDeleted *bool
}
