package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	config "s3l/mpcfgo/config"
)

func SamplePublicKey() []byte {
	sk, err := rsa.GenerateKey(rand.Reader, config.PubK)
	letItCrash(err)
	pk := x509.MarshalPKCS1PublicKey(&sk.PublicKey)
	return pk
}

func SamplePublicAndPrivateKey() ([]byte, *rsa.PrivateKey) {
	sk, err := rsa.GenerateKey(rand.Reader, config.PubK)
	letItCrash(err)
	pkBytes := x509.MarshalPKCS1PublicKey(&sk.PublicKey)
	return pkBytes, sk
}

func Encrypt(pkBytes []byte, data []byte) []byte {
	pk, err := x509.ParsePKCS1PublicKey(pkBytes)
	letItCrash(err)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pk, data)
	letItCrash(err)
	return ciphertext
}

func Decrypt(sk *rsa.PrivateKey, ciphertext []byte) []byte {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, sk, ciphertext)
	letItCrash(err)
	return plaintext
}

func letItCrash(err error) {
	if err != nil {
		panic(err)
	}
}
