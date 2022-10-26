package util

import (
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(id string, email string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		AesEncrypt(id, setting.AppSetting.Key),
		EncodeMD5(email),
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-fuji",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
