package models

import "github.com/google/uuid"

type About struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Profession  string    `json:"profession"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"is_available"`
	Handphone   string    `json:"handphone"`
	Email       string    `json:"email"`
	Resume      *string   `json:"resume"`
}
