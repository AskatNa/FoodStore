package model

import "time"

type Product struct {
	ID          uint64
	Name        string
	Description string
	Price       float64
	StockCount  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
}

type ProductFilter struct {
	ID        *uint64
	Name      *string
	PriceMin  *float64
	PriceMax  *float64
	InStock   *bool
	IsDeleted *bool
}

type ProductUpdateData struct {
	Name        *string
	Description *string
	Price       *float64
	StockCount  *int
	UpdatedAt   *time.Time
	IsDeleted   *bool
}
