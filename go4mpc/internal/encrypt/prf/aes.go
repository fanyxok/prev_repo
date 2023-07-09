package prf

import (
	"crypto/aes"
	"crypto/cipher"
)

type fixedKeyAES struct{}

var FixedKeyAES fixedKeyAES

const AesBlockSizeInBytes = 16
const AesBlockSizeInBits = 128

var AES_Key_Fixed, _ = aes.NewCipher([]byte{75, 51, 199, 95, 11, 223, 138, 9, 134, 123, 124, 88, 155, 34, 173, 147})
var IV = []byte{226, 59, 179, 35, 81, 209, 112, 134, 234, 21, 164, 235, 215, 164, 55, 218}

func (ct fixedKeyAES) Hash(raw []byte) []byte {
	ctr := cipher.NewCTR(AES_Key_Fixed, IV)
	ciphertext := make([]byte, len(raw))
	ctr.XORKeyStream(ciphertext, raw)
	return ciphertext
}
