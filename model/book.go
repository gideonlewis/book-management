package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID                int64           `json:"id"`
	Title             string          `json:"title"`
	Author            string          `json:"author"`
	Price             int64           `json:"price"`
	TotalQuantity     int64           `json:"total_quantity"`
	AvailableQuantity int64           `json:"available_quantity"`
	CreatedBy         int64           `json:"created_by"`
	UpdatedBy         *int64          `json:"updated_by"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         *gorm.DeletedAt `json:"-"`
}
