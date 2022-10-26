package partner_service

import (
	"fujitech/web/web-go/go-challenge/models"

	"github.com/google/uuid"
)

type Partner struct {
	Id          string
	UserId      string
	Name        string
	BankAccount string
	Phone       string
	Address     string
	Email       string
}

func (a *Partner) GetCurrentPartner() *Partner {
	return a
}

func (a *Partner) GetAll() []*models.Partner {
	return models.GetPartners(a.UserId)
}

func (a *Partner) CheckValid() bool {
	return models.CheckValidPartner(a.BankAccount)
}

func (a *Partner) CheckExist() bool {
	return models.CheckPartnerExist(a.Id)
}

func (a *Partner) Add() error {
	partner := map[string]interface{}{
		"id":           uuid.New().String(),
		"user_id":      a.UserId,
		"name":         a.Name,
		"bank_account": a.BankAccount,
		"phone":        a.Phone,
		"address":      a.Address,
		"email":        a.Email,
	}

	if err := models.AddPartner(partner); err != nil {
		return err
	}

	return nil
}

func (a *Partner) Edit() error {
	partner := map[string]interface{}{
		"id":           a.Id,
		"user_id":      a.UserId,
		"name":         a.Name,
		"bank_account": a.BankAccount,
		"phone":        a.Phone,
		"address":      a.Address,
		"email":        a.Email,
	}

	if err := models.EditPartner(a.Id, partner); err != nil {
		return err
	}

	return nil
}

func (a *Partner) Delete() error {
	if err := models.DeletePartner(a.Id); err != nil {
		return err
	}
	return nil
}
