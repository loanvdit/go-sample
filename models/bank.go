package models

type Bank struct {
	ID   string `gorm:"primary_key" json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Model
}
