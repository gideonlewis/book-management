package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	Author    string          `json:"author"`
	Price     int64           `json:"price"`
	Quantity  int64           `json:"quantity"`
	Remaining int64           `json:"remaining"`
	CreatedBy int64           `json:"created_by"`
	UpdatedBy *int64          `json:"updated_by"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}
