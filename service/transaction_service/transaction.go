package transaction_service

import (
	"fujitech/web/web-go/go-challenge/models"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	UserId       string
	PartnerId    string
	BankId       string
	BankBranchId string
	Amount       int64
	Fee          int64
	Message      string
	Currency     int
	SendedAt     time.Time
	ReceivedAt   time.Time
	Status       int
	Mode         int
}

func (a *Transaction) GetCurrentTransaction() *Transaction {
	return a
}

func (a *Transaction) GetAll() []*models.Transaction {
	return models.GetTransactions(a.UserId)
}

func (a *Transaction) GetId() models.Transaction {
	t, err := models.GetTransaction(a.ID)
	if err != nil {
		return models.Transaction{}
	} else {
		return t
	}
}

func GetTransactionByDate(startDate string, endDate string) []*models.Transaction {
	return models.GetTransactionsByDate(startDate, endDate)
}

func (a *Transaction) CheckExist() bool {
	return models.CheckTransactionExist(a.ID)
}

func (a *Transaction) Add() error {
	partner := map[string]interface{}{
		"id":             uuid.New().String(),
		"user_id":        a.UserId,
		"partner_id":     a.PartnerId,
		"bank_id":        a.BankId,
		"bank_branch_id": a.BankBranchId,
		"amount":         a.Amount,
		"fee":            a.Fee,
		"message":        a.Message,
		"currency":       a.Currency,
		"sended_at":      time.Now(),
		"received_at":    time.Now(),
		"status":         0,
		"mode":           a.Mode,
	}

	if err := models.AddTransaction(partner); err != nil {
		return err
	}

	return nil
}

func (a *Transaction) Edit() error {
	partner := map[string]interface{}{
		"id":             a.ID,
		"user_id":        a.UserId,
		"partner_id":     a.PartnerId,
		"bank_id":        a.BankId,
		"bank_branch_id": a.BankBranchId,
		"amount":         a.Amount,
		"fee":            a.Fee,
		"message":        a.Message,
		"currency":       a.Currency,
		"sended_at":      a.SendedAt,
		"received_at":    a.ReceivedAt,
		"status":         a.Status,
		"mode":           a.Mode,
	}

	if err := models.EditTransaction(a.ID, partner); err != nil {
		return err
	}

	return nil
}

func (a *Transaction) Delete() error {
	if err := models.DeleteTransaction(a.ID); err != nil {
		return err
	}

	return nil
}

func UpdateStatus(id string, status int) bool {
	return models.UpdateStatusTransaction(id, status)
}
