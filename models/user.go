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

type UserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserLoginDTO struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
}
