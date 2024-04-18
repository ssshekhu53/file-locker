package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"ssshekhu53/file-locker/services"
)

const secret = `qwertyuiopasdfghjklzxcvb`

type crypt struct {
	block cipher.Block
}

func New() (services.Crypt, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &crypt{block: block}, nil
}

// Encrypt method is to encrypt crypt
func (c *crypt) Encrypt(creds []byte) []byte {
	ciphertext := make([]byte, aes.BlockSize+len(creds))
	iv := ciphertext[:aes.BlockSize]

	cfb := cipher.NewCFBEncrypter(c.block, iv)
	cipherText := make([]byte, len(creds))
	cfb.XORKeyStream(cipherText, creds)

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(cipherText)))

	base64.StdEncoding.Encode(dst, cipherText)

	return dst
}

func (c *crypt) Decrypt(cred string) ([]byte, error) {
	ciphertext := make([]byte, aes.BlockSize+len(cred))

	iv := ciphertext[:aes.BlockSize]
	cfb := cipher.NewCFBDecrypter(c.block, iv)

	cipherText, err := base64.StdEncoding.DecodeString(cred)
	if err != nil {
		return []byte{}, err
	}

	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return plainText, nil
}
