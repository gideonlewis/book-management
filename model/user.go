package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	UserName  string          `json:"user_name"`
	Email     string          `json:"email"`
	Gender    string          `json:"gender"`
	Team      string          `json:"team"`
	JoinDate  time.Time       `json:"join_date"`
	CreatedBy int64           `json:"created_by"`
	UpdatedBy *int64          `json:"updated_by"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}
