package simpleot

import (
	config "s3l/mpcfgo/config"
	tRSA "s3l/mpcfgo/internal/encrypt/asym/rsa"
	tAES "s3l/mpcfgo/internal/encrypt/sym/aes"
	"s3l/mpcfgo/internal/network"
)

func Send(net network.Network, in0 []byte, in1 []byte) {
	// encrypt data by sym key
	sk0 := tAES.SampleSecretKey(config.SymK)
	sk1 := tAES.SampleSecretKey(config.SymK)
	aes0 := tAES.NewAES(sk0)
	aes1 := tAES.NewAES(sk1)

	ciphertext0 := aes0.Encrypt(in0)
	ciphertext1 := aes1.Encrypt(in1)

	// encrypt sym key by asym pk
	m0 := net.Recv()
	m1 := net.Recv()

	encrypted0 := tRSA.Encrypt(m0.Data, sk0)
	encrypted1 := tRSA.Encrypt(m1.Data, sk1)

	// send both cipher and encrypted sk
	net.Send(network.NewMsg(0000, network.SimpleOT, ciphertext0))
	net.Send(network.NewMsg(0001, network.SimpleOT, ciphertext1))
	net.Send(network.NewMsg(0010, network.SimpleOT, encrypted0))
	net.Send(network.NewMsg(0011, network.SimpleOT, encrypted1))

}

func Recv(net network.Network, choice bool) []byte {
	pk0 := tRSA.SamplePublicKey()
	pk1, sk1 := tRSA.SamplePublicAndPrivateKey()

	// send a true pk, a fake pk
	if choice {
		net.Send(network.NewMsg(0, network.SimpleOT, pk0))
		net.Send(network.NewMsg(1, network.SimpleOT, pk1))
	} else {
		net.Send(network.NewMsg(0, network.SimpleOT, pk1))
		net.Send(network.NewMsg(1, network.SimpleOT, pk0))
	}

	// recv two aes key encrypted by pk, and two ciphertext encrypted by aes
	ciphertext0 := net.Recv()
	ciphertext1 := net.Recv()
	encrypted0 := net.Recv()
	encrypted1 := net.Recv()

	var plaintext []byte

	if choice {
		sk := tRSA.Decrypt(sk1, encrypted1.Data)
		aes := tAES.NewAES(sk)
		plaintext = aes.Decrypt(ciphertext1.Data)
	} else {
		sk := tRSA.Decrypt(sk1, encrypted0.Data)
		aes := tAES.NewAES(sk)
		plaintext = aes.Decrypt(ciphertext0.Data)
	}
	return plaintext
}
