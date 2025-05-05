package models

import "time"

type Work struct {
	Id          string     `json:"id" gorm:"primaryKey;not null"`
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
