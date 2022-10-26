package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Partner struct {
	ID          string `gorm:"primary_key" json:"id"`
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	BankAccount string `json:"bank_account"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Model
}

// Get partners of user
func GetPartners(user_id string) []*Partner {
	var partners []*Partner
	err := db.Select("*").Where("user_id=?", user_id).Find(&partners).Error
	if err == nil && len(partners) > 0 {
		return partners
	}

	return nil
}

// Check existing partner (bank_account)
func CheckValidPartner(bank_account string) bool {
	var partners []Partner
	err := db.Select("*").Where("bank_account = ?", bank_account).First(&partners).Error
	if err != nil || len(partners) > 0 {
		return false
	}

	return true
}

// Check partner existing
func CheckPartnerExist(id string) bool {
	var partner Partner
	err := db.Select("*").Where("id = ?", id).First(&partner).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// Add new partner
func AddPartner(data map[string]interface{}) error {
	partner := Partner{
		ID:          data["id"].(string),
		UserId:      data["user_id"].(string),
		Name:        data["name"].(string),
		BankAccount: data["bank_account"].(string),
		Phone:       data["phone"].(string),
		Address:     data["address"].(string),
		Email:       data["email"].(string),
	}
	partner.Model.CreatedAt = time.Now()
	partner.Model.UpdatedAt = time.Now()

	if err := db.Create(&partner).Error; err != nil {
		return err
	}

	return nil
}

// Edit a partner
func EditPartner(id string, data interface{}) error {
	if err := db.Model(&Partner{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// Delete a partner
func DeletePartner(id string) error {
	if err := db.Where("id = ?", id).Delete(&Partner{}).Error; err != nil {
		return err
	}

	return nil
}

// Delete all partner
func CleanAllPartner() (bool, error) {
	if err := db.Model(&Partner{}).Where("id != ? ", 0).Delete(&Partner{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
