package models

import "time"

type Stack struct {
	ID        string     `json:"id" gorm:"primaryKey;not null"`
	Name      string     `json:"name" gorm:"not null"`
	Image     string     `json:"image" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type StackResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type StackDTO struct {
	Name  string `json:"name" validate:"require"`
	Image string `json:"image" validate:"require"`
}
