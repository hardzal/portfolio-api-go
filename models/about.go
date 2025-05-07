package models

import (
	"time"
)

type About struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title" form:"title" gorm:"type:varchar(100);not null"`
	Profession  string     `json:"profession" form:"profession" gorm:"type:varchar(100);not null"`
	Description string     `json:"description" form:"description"  gorm:"type:text;not null"`
	Location    string     `json:"location" form:"location" gorm:"type:varchar(100);not null"`
	IsAvailable bool       `json:"is_available" form:"is_available" gorm:"default:true"`
	ImageUrl    string     `json:"image" form:"image" gorm:"type:text;not null"`
	Handphone   string     `json:"handphone" form:"handphone" gorm:"type:varchar(100);not null"`
	Email       string     `json:"email" form:"email" gorm:"type:varchar(200);not null"`
	Resume      *string    `json:"resume" form:"resume" gorm:"type:text"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type AboutResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Profession  string     `json:"profession"`
	Description string     `json:"description"`
	Location    string     `json:"location"`
	IsAvailable bool       `json:"is_available"`
	ImageUrl    string     `json:"image"`
	Handphone   string     `json:"handphone"`
	Email       string     `json:"email"`
	Resume      string     `json:"resume"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type AboutDTO struct {
	Title       string `json:"title" validate:"required"`
	Profession  string `json:"profession" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	IsAvailable bool   `json:"is_available" validate:"required"`
	Handphone   string `json:"handphone" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Resume      string `json:"resume" validate:"required"`
}
