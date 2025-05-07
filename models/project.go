package models

import "time"

type Project struct {
	ID          string     `json:"id" gorm:"primaryKey;not null"`
	Title       string     `json:"title" form:"title" validate:"required" gorm:"not null"`
	Description string     `json:"description" form:"description" validate:"required" gorm:"not null"`
	ImageUrl    *string    `json:"image_url" form:"image_url"`
	Stacks      []string   `json:"stacks" form:"stacks" gorm:"not null;type:text[]"`
	Repo        *string    `json:"repo" form:"repo"` // nil
	Demo        *string    `json:"demo" form:"demo"` // nil
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProjectResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageUrl    *string    `json:"image_url"`
	Stacks      []string   `json:"stacks"`
	Repo        *string    `json:"repo"`
	Demo        *string    `json:"demo"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type ProjectDTO struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Stacks      []string `json:"stacks" validate:"required"`
	Repo        *string  `json:"repo" validate:"required"`
	Demo        *string  `json:"demo" validate:"required"`
}
