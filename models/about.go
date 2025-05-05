package models

import (
	"time"

	"github.com/google/uuid"
)

type About struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title       string     `json:"title" form:"title" gorm:"not null"`
	Profession  string     `json:"profession" form:"profession" gorm:"not null"`
	Description string     `json:"description" form:"description"  gorm:"not null"`
	Location    string     `json:"location" form:"location" gorm:"not null"`
	IsAvailable bool       `json:"is_available" form:"is_available" gorm:"default:true"`
	Handphone   string     `json:"handphone" form:"handphone" gorm:"not null"`
	Email       string     `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Resume      *string    `json:"resume" form:"resume"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
