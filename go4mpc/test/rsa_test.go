package test

import (
	"crypto/rand"
	"fmt"
	"s3l/mpcfgo/internal/encrypt/asym/rsa"
	"testing"
)

func TestRSA(t *testing.T) {
	for i := 0; i < 100; i++ {
		pk, sk := rsa.SamplePublicAndPrivateKey()
		fmt.Println("Key bytes", len(pk))
		bytes := make([]byte, 17)
		rand.Read(bytes)
		encryped := rsa.Encrypt(pk, bytes)
		decrypted := rsa.Decrypt(sk, encryped)

		fmt.Println(bytes, "\n", decrypted)
		fmt.Println()
	}

}
