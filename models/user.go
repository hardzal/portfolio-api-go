package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Username  string     `json:"username" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null"`
	Password  string     `json:"password" gorm:"not null;column:password"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
