package models

type Work struct {
	Id          string   `json:"id"`
	Role        string   `json:"role"`
	Company     string   `json:"company"`
	Description string   `json:"description"`
	Stacks      []string `json:"stack"`
	Image       string   `json:"image"`
	StartDate   *string  `json:"start_date"`
	EndDate     *string  `json:"end_date"`
}
