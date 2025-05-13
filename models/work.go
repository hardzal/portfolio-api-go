package models

import (
	"time"

	"github.com/lib/pq"
)

type Work struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Role        string         `json:"role" gorm:"type:varchar(100);not null"`
	Company     string         `json:"company" gorm:"type:varchar(100);not null"`
	Description pq.StringArray `json:"description" gorm:"type:text[];not null"`
	Stacks      pq.StringArray `json:"stack" gorm:"type:text[];not null"`
	Image       string         `json:"image" gorm:"type:text;not null"`
	StartDate   *string        `json:"start_date" form:"start_date" gorm:"type:text"`
	EndDate     *string        `json:"end_date" form:"end_date" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

type WorkResponse struct {
	ID          uint           `json:"id"`
	Role        string         `json:"role"`
	Company     string         `json:"company"`
	Description pq.StringArray `json:"description"`
	Stacks      pq.StringArray `json:"stack"`
	Image       string         `json:"image"`
	StartDate   *string        `json:"start_date"`
	EndDate     *string        `json:"end_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
}

type WorkDTO struct {
	Role        string         `json:"role" validate:"required"`
	Company     string         `json:"company" validate:"required"`
	Description pq.StringArray `json:"description" validate:"required"`
	Stacks      pq.StringArray `json:"stack" validate:"required"`
	Image       string         `json:"image" validate:"required"`
	StartDate   *string        `json:"start_date" validate:"required"`
	EndDate     *string        `json:"end_date" validate:"required"`
}
