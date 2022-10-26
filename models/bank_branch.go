package models

type BankBranch struct {
	ID      string `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Bank    Bank
	Model
}
