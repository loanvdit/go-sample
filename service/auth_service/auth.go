package auth_service

import "fujitech/web/web-go/go-challenge/models"

type Auth struct {
	Id        string
	Email     string
	Password  string
	VerifyKey string
}

func (a *Auth) Check() (bool, string, string, error) {
	return models.CheckAuth(a.Email, a.Password)
}

func (a *Auth) ValidKey() (bool, error) {
	return models.CheckValidKey(a.Id, a.VerifyKey)
}
