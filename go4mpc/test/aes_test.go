package test

import (
	"fmt"
	"math/rand"
	tAES "s3l/mpcfgo/internal/encrypt/sym/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	for i := 0; i < 10000; i++ {
		key := tAES.SampleSecretKey(128)
		fmt.Println("Key bytes", len(key), key)
		ase := tAES.NewAES(key)
		bytes := make([]byte, 17)
		rand.Read(bytes)
		encryped := ase.Encrypt(bytes)
		decrypted := ase.Decrypt(encryped)

		for j := 0; j < len(bytes); j++ {
			assert.Equal(nil, bytes, decrypted, "error")
		}
	}

}
