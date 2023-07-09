package baseot

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand"
	"s3l/mpcfgo/internal/encrypt/prf"
	"s3l/mpcfgo/internal/network"
	"sync"
)

// implement asharov-lindell OT protocol of paper <<More Efficient Oblivious Transfer and Extensions for
// Faster Secure Computation>>

var q, _ = new(big.Int).SetString("18446744073709551616", 10)
var g = new(big.Int).SetUint64(4518808235179270133)

// in0 and in1 have the same size
func Send(net network.Network, in0 []byte, in1 []byte) {
	if len(in0) != len(in1) {
		panic(fmt.Sprintf("two bytes have different length %d, %d\n", len(in0), len(in1)))
	}

	r := new(big.Int).SetUint64(rand.Uint64())
	u := new(big.Int).Exp(g, r, q)

	h0 := new(big.Int).SetBytes(net.Recv().Data)
	h1 := new(big.Int).SetBytes(net.Recv().Data)

	k0 := new(big.Int).Exp(h0, r, q)
	k1 := new(big.Int).Exp(h1, r, q)

	prfk0 := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k0.Bytes())) // PRF the key
	prfk1 := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k1.Bytes())) // ...

	ciphertext0 := new(big.Int).Xor(prfk0, new(big.Int).SetBytes(in0)) // Mask plaintext
	ciphertext1 := new(big.Int).Xor(prfk1, new(big.Int).SetBytes(in1)) // ...

	length := len(in0)

	verify := make([]byte, 4)
	binary.LittleEndian.PutUint32(verify, uint32(length))
	net.Send(network.NewMsg(00, network.SimpleOT, verify))
	net.Send(network.NewMsg(00, network.SimpleOT, u.Bytes()))

	net.Send(network.NewMsg(10, network.SimpleOT, ciphertext0.Bytes()))
	net.Send(network.NewMsg(11, network.SimpleOT, ciphertext1.Bytes()))

}

func Recv(net network.Network, choice bool) []byte {
	exp := new(big.Int).SetUint64(rand.Uint64())
	h := new(big.Int).SetUint64(rand.Uint64())
	var h0, h1 *big.Int
	if choice {
		h0 = h
		h1 = new(big.Int).Exp(g, exp, q)
	} else {
		h0 = new(big.Int).Exp(g, exp, q)
		h1 = h
	}

	net.Send(network.NewMsg(0, network.SimpleOT, h0.Bytes()))
	net.Send(network.NewMsg(1, network.SimpleOT, h1.Bytes()))

	length := binary.LittleEndian.Uint32(net.Recv().Data)
	u := new(big.Int).SetBytes(net.Recv().Data)
	k := new(big.Int).Exp(u, exp, q)

	prfk := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k.Bytes())) // PRF the key

	var ciphertext *big.Int

	if choice {
		net.Recv()
		ciphertext = new(big.Int).SetBytes(net.Recv().Data)
	} else {
		ciphertext = new(big.Int).SetBytes(net.Recv().Data)
		net.Recv()
	}

	x := new(big.Int).Xor(ciphertext, prfk)

	final := make([]byte, length)
	x.FillBytes(final)
	return final

}

// func RecvN(net network.Network, b []bool) [][]byte {
// 	num := len(b)
// 	final := make([][]byte, num)
// 	ch := make(chan *big.Int, 256)
// 	done := make(chan bool)

// 	length := binary.LittleEndian.Uint32(net.Recv().Data)
// 	u := new(big.Int).SetBytes(net.Recv().Data)

// 	go func() {
// 		for _, v := range b {
// 			exp := new(big.Int).SetUint64(rand.Uint64())
// 			ch <- exp
// 			h := new(big.Int).SetUint64(rand.Uint64())
// 			var h0 *big.Int
// 			var h1 *big.Int
// 			if v {
// 				h0 = h
// 				h1 = new(big.Int).Exp(g, exp, q)
// 			} else {
// 				h0 = new(big.Int).Exp(g, exp, q)
// 				h1 = h
// 			}
// 			net.Send(network.NewMsg(0, network.SimpleOT, h0.Bytes()))
// 			net.Send(network.NewMsg(1, network.SimpleOT, h1.Bytes()))
// 		}
// 		done <- true
// 	}()

// 	go func() {
// 		for i, v := range b {
// 			exp := <-ch
// 			k := new(big.Int).Exp(u, exp, q)
// 			prfk := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k.Bytes())) // PRF the key

// 			var ciphertext *big.Int
// 			if v {
// 				net.Recv()
// 				ciphertext = new(big.Int).SetBytes(net.Recv().Data)
// 			} else {
// 				ciphertext = new(big.Int).SetBytes(net.Recv().Data)
// 				net.Recv()
// 			}

// 			x := new(big.Int).Xor(ciphertext, prfk)
// 			final[i] = make([]byte, length)
// 			x.FillBytes(final[i])
// 		}
// 		done <- true
// 	}()
// 	<-done
// 	<-done
// 	return final
// }

func Send128(net network.Network, in0 [][]byte, in1 [][]byte) {
	r := new(big.Int).SetUint64(rand.Uint64())
	u := new(big.Int).Exp(g, r, q)
	net.Send(network.NewMsg(00, network.SimpleOT, u.Bytes()))
	inbuf := net.Recv().Data
	outbuf := make([]byte, 128*32)
	var wg sync.WaitGroup
	for i := 0; i < 128; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			h0 := new(big.Int).SetBytes(inbuf[i*32 : i*32+16])
			h1 := new(big.Int).SetBytes(inbuf[i*32+16 : i*32+32])
			k0 := new(big.Int).Exp(h0, r, q)
			k1 := new(big.Int).Exp(h1, r, q)

			prfk0 := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k0.Bytes())) // PRF the key
			prfk1 := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k1.Bytes())) // ...

			new(big.Int).Xor(prfk0, new(big.Int).SetBytes(in0[i])).FillBytes(outbuf[i*32 : i*32+16])    // Mask plaintext
			new(big.Int).Xor(prfk1, new(big.Int).SetBytes(in1[i])).FillBytes(outbuf[i*32+16 : i*32+32]) // ...
		}(i)
	}
	wg.Wait()
	net.Send(network.NewMsg(10, network.SimpleOT, outbuf))
}

func Recv128(net network.Network, b []bool) [][]byte {
	final := make([][]byte, 128)
	for i := range final {
		final[i] = make([]byte, 16)
	}
	u := new(big.Int).SetBytes(net.Recv().Data)
	buf := make([]byte, 32*128)
	exp := make([]*big.Int, 128)
	var wg sync.WaitGroup
	for i := 0; i < 128; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			exp[i] = new(big.Int).SetUint64(rand.Uint64())
			h := new(big.Int).SetUint64(rand.Uint64())
			var h0 *big.Int
			var h1 *big.Int
			if b[i] {
				h0 = h
				h1 = new(big.Int).Exp(g, exp[i], q)
			} else {
				h0 = new(big.Int).Exp(g, exp[i], q)
				h1 = h
			}
			h0.FillBytes(buf[i*32 : i*32+16])
			h1.FillBytes(buf[i*32+16 : i*32+32])
		}(i)
	}
	wg.Wait()
	net.Send(network.NewMsg(0, network.SimpleOT, buf))
	buf = net.Recv().Data
	for i := 0; i < 128; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			k := new(big.Int).Exp(u, exp[i], q)
			prfk := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k.Bytes())) // PRF the key

			var ciphertext *big.Int
			if b[i] {
				ciphertext = new(big.Int).SetBytes(buf[i*32+16 : i*32+32])
			} else {
				ciphertext = new(big.Int).SetBytes(buf[i*32 : i*32+16])
			}
			new(big.Int).Xor(ciphertext, prfk).FillBytes(final[i])
		}(i)
	}
	wg.Wait()
	return final
}

func SendN(net network.Network, in0, in1 [][]byte) {
	if len(in0) != len(in1) {
		panic(fmt.Sprintf("list of in0, and list of in1 have different length %d, %d\n", len(in0), len(in1)))
	}
	// u: pk(activated one) pk_prime: pk'
	pk_prime := new(big.Int).SetUint64(rand.Uint64())
	pk := new(big.Int).Exp(g, pk_prime, q)

	// exchange constant, length and u.

	
	ul := uint32(len(in0[0]))
	l := binary.LittleEndian.AppendUint32(nil, ul)
	// s1
	net.Send(network.NewMsg(00, network.SimpleOT, l))
	net.Send(network.NewMsg(00, network.SimpleOT, pk.Bytes()))
	bytes := len(in0[0])
	size := len(in0)
	outbuf := make([]byte, 2*size*bytes)

	inbuf := net.Recv().Data
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			h0 := new(big.Int).SetBytes(inbuf[2*i*bytes : 2*i*bytes+bytes])
			h1 := new(big.Int).SetBytes(inbuf[2*i*bytes+bytes : 2*(i+1)*bytes])

			k0 := new(big.Int).Exp(h0, pk_prime, q)
			k1 := new(big.Int).Exp(h1, pk_prime, q)

			prfk0 := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k0.Bytes())) // PRF the key
			prfk1 := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k1.Bytes())) // ...

			new(big.Int).Xor(prfk0, new(big.Int).SetBytes(in0[i])).FillBytes(outbuf[2*i*bytes : 2*i*bytes+bytes])     // Mask plaintext
			new(big.Int).Xor(prfk1, new(big.Int).SetBytes(in1[i])).FillBytes(outbuf[2*i*bytes+bytes : 2*(i+1)*bytes]) // Mask plaintext                                         // ...
		}(i)
	}
	wg.Wait()
	net.Send(network.NewMsg(10, network.SimpleOT, outbuf))
}

func RecvN(net network.Network, b []bool) [][]byte {
	size := len(b)
	// c1
	bytes := int(binary.LittleEndian.Uint32(net.Recv().Data))

	final := make([][]byte, size)
	for i:= 0;i<size;i++{
		final[i] = make([]byte, bytes)
	}

	for i := range final {
		final[i] = make([]byte, bytes)
	}
	u := new(big.Int).SetBytes(net.Recv().Data)
	buf := make([]byte, bytes*size*2)
	exp := make([]*big.Int, size)
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			exp[i] = new(big.Int).SetUint64(rand.Uint64())
			h := new(big.Int).SetUint64(rand.Uint64())
			var h0 *big.Int
			var h1 *big.Int
			if b[i] {
				h0 = h
				h1 = new(big.Int).Exp(g, exp[i], q)
			} else {
				h0 = new(big.Int).Exp(g, exp[i], q)
				h1 = h
			}
			h0.FillBytes(buf[2*i*bytes : 2*i*bytes+bytes])
			h1.FillBytes(buf[2*i*bytes+bytes : 2*(i+1)*bytes])
		}(i)
	}
	wg.Wait()
	net.Send(network.NewMsg(0, network.SimpleOT, buf))
	buf = net.Recv().Data
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			k := new(big.Int).Exp(u, exp[i], q)
			prfk := new(big.Int).SetBytes(prf.FixedKeyAES.Hash(k.Bytes())) // PRF the key

			var ciphertext *big.Int
			if b[i] {
				ciphertext = new(big.Int).SetBytes(buf[2*i*bytes+bytes : 2*(i+1)*bytes])
			} else {
				ciphertext = new(big.Int).SetBytes(buf[2*i*bytes : 2*i*bytes+bytes])
			}
			new(big.Int).Xor(ciphertext, prfk).FillBytes(final[i])
		}(i)
	}
	wg.Wait()
	return final
}
