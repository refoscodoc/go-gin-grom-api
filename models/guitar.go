package models

type Guitar struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
}
