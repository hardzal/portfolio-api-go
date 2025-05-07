package models

import "time"

type Stack struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string     `json:"name" gorm:"type:varchar(100);not null"`
	Image     string     `json:"image" gorm:"type:text;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type StackResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type StackDTO struct {
	Name string `json:"name" validate:"require"`
}
