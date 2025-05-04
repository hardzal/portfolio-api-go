package models

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image_url"`
	Stacks      []string `json:"stacks"`
	Repo        *string  `json:"repo"` // nil
	Demo        *string  `json:"demo"` // nil
}
