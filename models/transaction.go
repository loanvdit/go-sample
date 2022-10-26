package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	ID           string    `gorm:"primary_key" json:"id"`
	UserId       string    `json:"user_id"`
	PartnerId    string    `json:"partner_id"`
	BankId       string    `json:"bank_id"`
	BankBranchId string    `json:"bank_branch_id"`
	Amount       int64     `json:"amount"`
	Fee          int64     `json:"fee"`
	Message      string    `json:"message"`
	Currency     int       `json:"currency"`
	SendedAt     time.Time `json:"sender_at"`
	ReceivedAt   time.Time `json:"received_at"`
	Status       int       `json:"status"`
	Mode         int       `json:"mode"`
	Model
}

// Get transactions of user
func GetTransactions(user_id string) []*Transaction {
	var transactions []*Transaction
	err := db.Select("*").
		Where("user_id=?", user_id).Find(&transactions).Error
	if err == nil && len(transactions) > 0 {
		return transactions
	}
	return nil
}

// Get transactions of user
func GetTransactionsByDate(start_date string, end_date string) []*Transaction {
	var transactions []*Transaction
	err := db.Select("*").
		Where("created_at BETWEEN  ? AND ?", start_date, end_date).Find(&transactions).Error
	if err == nil && len(transactions) > 0 {
		return transactions
	}
	return nil
}

// Get a transaction
func GetTransaction(id string) (Transaction, error) {
	var transaction Transaction
	err := db.Select("*").Where("id=?", id).Find(&transaction).Error
	if err == nil && err != gorm.ErrRecordNotFound {
		return transaction, nil
	}
	return transaction, err
}

// Update status transaction
func UpdateStatusTransaction(id string, status int) bool {
	var transaction Transaction
	if err := db.Model(&transaction).Where("id=?", id).Update(Transaction{Status: status}).Error; err != nil {
		return false
	}
	return true
}

// Check partner existing
func CheckTransactionExist(id string) bool {
	var t Transaction
	err := db.Select("*").Where("id=?", id).First(&t).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// Add new transaction
func AddTransaction(data map[string]interface{}) error {
	t := Transaction{
		ID:           data["id"].(string),
		UserId:       data["user_id"].(string),
		PartnerId:    data["partner_id"].(string),
		BankId:       data["bank_id"].(string),
		BankBranchId: data["bank_branch_id"].(string),
		Amount:       data["amount"].(int64),
		Fee:          data["fee"].(int64),
		Message:      data["message"].(string),
		Currency:     data["currency"].(int),
		SendedAt:     data["sended_at"].(time.Time),
		ReceivedAt:   data["received_at"].(time.Time),
		Status:       data["status"].(int),
		Mode:         data["mode"].(int),
	}
	t.Model.CreatedAt = time.Now()
	t.Model.UpdatedAt = time.Now()

	if err := db.Create(&t).Error; err != nil {
		return err
	}

	return nil
}

// Edit a transaction
func EditTransaction(id string, data interface{}) error {
	if err := db.Model(&Transaction{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// Remove a transaction (soft delete)
func DeleteTransaction(id string) error {
	if err := db.Where("id = ?", id).Delete(&Transaction{}).Error; err != nil {
		return err
	}

	return nil
}
