package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func AesEncrypt(orig string, key string) string {
	origData := []byte(orig)
	k := []byte(key)

	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key The length must be 16/24/32 length: %s", err.Error()))
	}

	blockSize := block.BlockSize()

	origData = PKCS7Padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])

	cryted := make([]byte, len(origData))

	blockMode.CryptBlocks(cryted, origData)

	return base64.RawURLEncoding.EncodeToString(cryted)

}

func AesDecrypt(cryted string, key string) string {
	crytedByte, _ := base64.RawURLEncoding.DecodeString(cryted)
	k := []byte(key)

	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key The length must be 16/24/32 length: %s", err.Error()))
	}

	blockSize := block.BlockSize()

	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])

	orig := make([]byte, len(crytedByte))

	blockMode.CryptBlocks(orig, crytedByte)

	orig = PKCS7UnPadding(orig)

	return string(orig)
}

// Complement
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// DeCode
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
