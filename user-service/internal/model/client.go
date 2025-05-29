package model

import "time"

const (
	//CustomerRole = "customer"
	//AdminRole    = "admin"
	OwnerRole = "owner"
)

type Customer struct {
	ID              uint64
	Name            string
	Phone           string
	Email           string
	CurrentPassword string
	NewPassword     string
	PasswordHash    string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	IsDeleted bool
}

type CustomerFilter struct {
	ID           *uint64
	Name         *string
	Phone        *string
	Email        *string
	PasswordHash *string
	Role         *string

	IsDeleted *bool
}

type CustomerUpdateData struct {
	ID           *uint64
	Name         *string
	Phone        *string
	Email        *string
	PasswordHash *string
	Role         *string
	UpdatedAt    *time.Time
	IsDeleted    *bool
}
