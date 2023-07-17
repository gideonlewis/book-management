package model

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	ID         int64           `json:"id"`
	UserID     int64           `json:"user_id"`
	BookID     int64           `json:"book_id"`
	BorrowDate time.Time       `json:"borrow_date"`
	ReturnDate *time.Time      `json:"return_date"`
	CreatedBy  int64           `json:"created_by"`
	UpdatedBy  *int64          `json:"updated_by"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"-"`
}
