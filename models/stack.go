package models

import "time"

type Stack struct {
	ID        string     `json:"id" gorm:"primaryKey;not null"`
	Name      string     `json:"name" validate:"require" gorm:"not null"`
	Image     string     `json:"image" validate:"require" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
