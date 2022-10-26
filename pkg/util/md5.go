package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// Check string md5
func CheckMD5(content, encrypted string) bool {
	return strings.EqualFold(EncodeMD5(content), encrypted)
}
