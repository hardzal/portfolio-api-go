package models

import (
	"time"

	"github.com/lib/pq"
)

type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title       string         `json:"title" form:"title" gorm:"type:varchar(100);not null"`
	Description string         `json:"description" form:"description" gorm:"type:text;not null"`
	ImageUrl    *string        `json:"image_url" form:"image_url" gorm:"type:text"`
	Stacks      pq.StringArray `json:"stacks" form:"stacks" gorm:"not null;type:text[]"`
	Repo        *string        `json:"repo" form:"repo" gorm:"type:text"` // nil
	Demo        *string        `json:"demo" form:"demo" gorm:"type:text"` // nil
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProjectResponse struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	ImageUrl    *string        `json:"image_url"`
	Stacks      pq.StringArray `json:"stacks"`
	Repo        *string        `json:"repo"`
	Demo        *string        `json:"demo"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
}

type ProjectDTO struct {
	Title       string         `json:"title" validate:"required" form:"title"`
	Description string         `json:"description" validate:"required" form:"description"`
	Stacks      pq.StringArray `json:"stacks" validate:"required" form:"stacks[]"`
	Repo        *string        `json:"repo" validate:"required" form:"repo"`
	Demo        *string        `json:"demo" validate:"required" form:"demo"`
}
