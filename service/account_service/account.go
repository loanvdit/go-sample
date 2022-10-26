package account_service

import (
	"fujitech/web/web-go/go-challenge/models"

	"github.com/google/uuid"
)

type Account struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Phone     string
	Address   string
	Avatar    string
}

func (a *Account) CheckValid() bool {
	return models.CheckValidUser(a.Phone, a.Email)
}

func (a *Account) ActiveAccount() bool {
	return models.ActiveUser(a.Email)
}

func (a *Account) Add() error {
	user := map[string]interface{}{
		"id":         uuid.New().String(),
		"first_name": a.FirstName,
		"last_name":  a.LastName,
		"phone":      a.Phone,
		"address":    a.Address,
		"avatar":     a.Avatar,
		"email":      a.Email,
		"password":   a.Password,
	}

	if err := models.AddUser(user); err != nil {
		return err
	}

	return nil
}
