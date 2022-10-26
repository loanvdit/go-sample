package util

import (
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"math/rand"
)

const letterBytes = "1234567890"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

// Random string
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
