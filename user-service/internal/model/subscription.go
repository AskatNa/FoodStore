package model

import "time"

const (
	SubscriptionStatusActive   = "active"
	SubscriptionStatusInactive = "inactive"
	SubscriptionStatusExpired  = "expired"
)

type Subscription struct {
	ID         uint64
	CustomerID uint64
	ProductID  uint64
	Status     string
	StartDate  time.Time
	EndDate    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsDeleted  bool
}

type SubscriptionFilter struct {
	ID         *uint64
	CustomerID *uint64
	ProductID  *uint64
	Status     *string
	FromDate   *time.Time
	ToDate     *time.Time
	IsDeleted  *bool
}

type SubscriptionUpdateData struct {
	Status    *string
	EndDate   *time.Time
	UpdatedAt *time.Time
	IsDeleted *bool
}
