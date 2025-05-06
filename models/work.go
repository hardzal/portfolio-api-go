package models

import "time"

type Work struct {
	ID          string     `json:"id" gorm:"primaryKey;not null"`
	Role        string     `json:"role" gorm:"not null"`
	Company     string     `json:"company" gorm:"not null"`
	Description string     `json:"description" gorm:"not null"`
	Stacks      []string   `json:"stack" gorm:"not null;type:text[]"`
	Image       string     `json:"image" gorm:"not null"`
	StartDate   *string    `json:"start_date" form:"start_date"`
	EndDate     *string    `json:"end_date" form:"end_date"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type WorkResponse struct {
	ID          string     `json:"id"`
	Role        string     `json:"role"`
	Company     string     `json:"company"`
	Description string     `json:"description"`
	Stacks      []string   `json:"stack"`
	Image       string     `json:"image"`
	StartDate   *string    `json:"start_date"`
	EndDate     *string    `json:"end_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type WorkDTO struct {
	Role        string   `json:"role" validate:"required"`
	Company     string   `json:"company" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Stacks      []string `json:"stack" validate:"required"`
	Image       string   `json:"image" validate:"required"`
	StartDate   *string  `json:"start_date" validate:"required"`
	EndDate     *string  `json:"end_date" validate:"required"`
}
