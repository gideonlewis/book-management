package model

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	ID             int64           `json:"id"`
	BorrowerID     int64           `json:"borrower_id"`
	BookID         int64           `json:"book_id"`
	NumOfBorrowed  int64           `json:"num_of_borrowsed"`
	ExpirationDate time.Time       `json:"expiration_date"`
	CreatedBy      int64           `json:"created_by"`
	UpdatedBy      *int64          `json:"updated_by"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      *gorm.DeletedAt `json:"-"`
}
