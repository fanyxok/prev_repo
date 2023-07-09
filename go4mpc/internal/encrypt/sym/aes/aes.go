package ase

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
)

var IV []byte = []byte("1234567812345678")

type AES struct {
	cipher cipher.Block
	mode   string
}

func NewAES(secretKey []byte) AES {
	instance, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(err.Error())
	}
	return AES{instance, "CTR"}
}

func (ct AES) Encrypt(plaintext []byte) []byte {
	ctr := cipher.NewCTR(ct.cipher, IV)
	ciphertext := make([]byte, len(plaintext))
	ctr.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}

func (ct AES) Decrypt(ciphertext []byte) []byte {
	ctr := cipher.NewCTR(ct.cipher, IV)
	plaintext := make([]byte, len(ciphertext))
	ctr.XORKeyStream(plaintext, ciphertext)
	return plaintext
}

func SampleSecretKey(size int) []byte {
	var bytes int = size / 8
	token := make([]byte, bytes)
	rand.Read(token)
	return token
}
