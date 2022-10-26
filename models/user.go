package models

import (
	"fujitech/web/web-go/go-challenge/pkg/util"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID                 string    `gorm:"primary_key" json:"id"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	VerifyKey          string    `json:"verify_key"`
	VerifyKeyCreatedAt time.Time `gorm:"column:verify_key_created_at"`
	IsActive           int       `json:"is_active"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Phone              string    `json:"phone"`
	Address            string    `json:"address"`
	Avatar             string    `json:"avatar"`
	Model
}

// CheckAuth checks if authentication information
func CheckAuth(email, password string) (bool, string, string, error) {
	var user User
	err := db.Select("*").Where(User{Email: email, Password: password, IsActive: 1}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, "", "", err
	}

	if len(user.ID) > 0 {
		verifyKey := util.RandString(5)
		if err := db.Model(&user).Where("id = ?", user.ID).Update(User{VerifyKey: verifyKey, VerifyKeyCreatedAt: time.Now()}).Error; err != nil {
			return false, "", "", err
		} else {
			return true, user.ID, verifyKey, nil
		}
	}

	return false, "", "", nil
}

// Check valid key
func CheckValidKey(id string, verifyKey string) (bool, error) {
	var user User
	if len(id) > 0 {
		err := db.Select("id").Where(User{ID: id, VerifyKey: verifyKey}).Where("verify_key_created_at BETWEEN ? AND ?", time.Now().Add(-time.Minute*30), time.Now()).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return false, err
		}
	}

	if len(user.ID) > 0 {
		return true, nil
	}

	return false, nil
}

// Check existing user (phone or email)
func CheckValidUser(phone string, email string) bool {
	var users []User
	err := db.Select("id").Where("phone = ? OR email = ?", phone, email).First(&users).Error
	if err != nil || len(users) > 0 {
		return false
	}

	return true
}

// Active an user
func ActiveUser(email string) bool {
	var user User
	if err := db.Model(&user).Where("email = ?", email).Update(User{IsActive: 1}).Error; err != nil {
		return false
	}

	return true
}

// Add new user
func AddUser(data map[string]interface{}) error {
	user := User{
		ID:        data["id"].(string),
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Address:   data["address"].(string),
		Avatar:    data["avatar"].(string),
		Email:     data["email"].(string),
		Password:  data["password"].(string),
	}
	user.Model.CreatedAt = time.Now()
	user.Model.UpdatedAt = time.Now()

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
