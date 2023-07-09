package test

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"s3l/mpcfgo/internal/encrypt/prf"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/ot/baseot"
	"s3l/mpcfgo/internal/ot/ote"
	"s3l/mpcfgo/pkg/fast"
	"testing"
	"unsafe"

	"github.com/forgoer/openssl"
	"github.com/stretchr/testify/assert"
)

func TestFunc_BOT(t *testing.T) {
	n := 128
	bits := 128
	bytes := (bits + 7) / 8
	c := make(chan [][]byte, 1)
	// Sender input && Routine
	x0 := make([][]byte, n)
	x1 := make([][]byte, n)
	for i := 0; i < n; i++ {
		x0[i] = make([]byte, bytes)
		rand.Read(x0[i])
		x1[i] = make([]byte, bytes)
		rand.Read(x1[i])
	}
	//Receiver input && Routine
	b := make([]bool, n)
	for i := 0; i < n; i++ {
		b[i] = misc.Bool()
	}
	ser, cli := Nets()
	Parallel(
		func() {
			baseot.Send128(ser, x0, x1)
		},
		func() {
			c <- baseot.Recv128(cli, b)
		})
	x := <-c
	_ = x
	for i := 0; i < n; i++ {
		if b[i] {
			assert.EqualValues(t, x1[i], x[i])
		} else {
			assert.EqualValues(t, x0[i], x[i])
		}
		if i%100 == 0 {
			fmt.Println("--- --- --- Sample", i/100, "--- --- ---")
			if b[i] {
				fmt.Println(x1[i])
			} else {
				fmt.Println(x0[i])
			}
			fmt.Println(x[i])
		}
	}
}

func BenchmarkBOTN(b *testing.B) {
	ser, cli := Nets()
	n := 128
	x0 := make([][]byte, n)
	x1 := make([][]byte, n)
	o := make([][]byte, n)
	for i := 0; i < n; i++ {
		x0[i] = make([]byte, 16)
		rand.Read(x0[i])
		x1[i] = make([]byte, 16)
		rand.Read(x1[i])
		o[i] = make([]byte, 16)
	}
	bo := make([]bool, n)
	for i := 0; i < n; i++ {
		bo[i] = misc.Bool()
	}
	for i := 0; i < b.N; i++ {
		Parallel(
			func() {
				//baseot.Send128(ser, x0, x1)
				baseot.SendN(ser, x0, x1)
			},
			func() {
				//baseot.Recv128(cli, bo)
				baseot.RecvN(cli, bo)
			})
	}
}
func TestBytesBools(t *testing.T) {
	m := make([]byte, 126)
	d := make([]bool, 126*8)
	dd := make([]bool, 126*8)
	mm := make([]byte, 126)
	fast.Bytes2Bools(d, m)
	fast.Bools2Bytes(mm, d)
	assert.EqualValues(t, m, mm)
	fast.Bytes2Bools(dd, mm)
	assert.EqualValues(t, d, dd)
}
func BenchmarkBytesBools(b *testing.B) {
	m := make([]byte, b.N)
	d := make([]bool, b.N*8)
	mm := make([]byte, b.N)
	b.ResetTimer()
	fast.Bytes2Bools(d, m)
	fast.Bools2Bytes(mm, d)
}
func TestECB(t *testing.T) {
	n := 32
	src := make([]byte, n)
	dst := make([]byte, n)
	rand.Read(src)
	ecb := openssl.NewECBEncrypter(prf.AES_Key_Fixed)
	ecb.CryptBlocks(dst[:16], src[:16])
	ecb.CryptBlocks(dst[16:], src[16:])
	fmt.Println(dst)
	ecb.CryptBlocks(dst, src)
	fmt.Println(dst)

}
func TestFixedKeyAES(t *testing.T) {
	n := 128
	src := make([]byte, n)
	dst := make([]byte, n)
	// t.Run("CFB", func(t *testing.T) {
	// 	ctr := cipher.NewCFBEncrypter(prf.AES_Key_Fixed, prf.IV)
	// 	ctr.XORKeyStream(dst, src)
	// 	fmt.Printf("%v", dst)
	// 	ctr.XORKeyStream(dst, src)
	// 	fmt.Printf("\n%v\n", dst)
	// })
	// t.Run("CTR", func(t *testing.T) {
	// 	ctr := cipher.NewCTR(prf.AES_Key_Fixed, prf.IV)
	// 	ctr.XORKeyStream(dst, src)
	// 	fmt.Printf("%v", dst)
	// 	ctr.XORKeyStream(dst, src)
	// 	fmt.Printf("\n%v\n", dst)
	// })
	// t.Run("ECB", func(t *testing.T) {
	// 	ecb := openssl.NewECBEncrypter(prf.AES_Key_Fixed)
	// 	ecb.CryptBlocks(dst, src)
	// 	fmt.Printf("%v", dst)
	// 	ecb.CryptBlocks(dst, src)
	// 	fmt.Printf("\n%v\n", dst)
	// })
	t.Run("GCM", func(t *testing.T) {
		// cgm, _ := cipher.NewGCM(prf.AES_Key_Fixed)
		// fmt.Printf("cgm.NonceSize(): %v\n", cgm.NonceSize())
		// fmt.Printf("cgm.Overhead(): %v\n", cgm.Overhead())
		// src0, src1 := make([]byte, n), make([]byte, n)
		// rand.Read(src0)
		// dst0, dst1 := make([]byte, 2), make([]byte, 2)
		// dst := cgm.Seal(dst0, prf.IV[:12], src0, nil)
		// fmt.Printf("%v", dst)
		// fmt.Printf("len(dst): %v\n", len(dst))
		// dst = cgm.Seal(dst1, prf.IV[:12], src1, nil)
		// fmt.Printf("\n%v\n", dst)
	})
	t.Run("CBC", func(t *testing.T) {
		cbc := cipher.NewCBCEncrypter(prf.AES_Key_Fixed, prf.IV)
		cbc.CryptBlocks(dst, src)
		fmt.Printf("%v", dst)
		cbc.CryptBlocks(dst, src)
		fmt.Printf("\n%v\n", dst)
	})
	t.Run("ECB-Normal", func(t *testing.T) {
		n := 16
		src := make([]byte, n)
		dst := make([]byte, n)
		ecb := openssl.NewECBEncrypter(prf.AES_Key_Fixed)
		ecb.CryptBlocks(dst, src)
		fmt.Printf("%v", dst)
		ecb.CryptBlocks(dst, src)
		fmt.Printf("\n%v\n", dst)
	})
}

func BenchmarkFixedKeyAES(b *testing.B) {
	b.Run("GCM-Normal", func(b *testing.B) {
		n := 16
		src := make([]byte, n)
		dst := make([]byte, n)
		b.ResetTimer()
		cgm, _ := cipher.NewGCM(prf.AES_Key_Fixed)
		for i := 0; i < b.N; i++ {
			cgm.Seal(dst, prf.IV[:12], src, nil)
		}
	})
	// b.Run("GCM-Batch", func(b *testing.B) {
	// 	src := make([]byte, 8*b.N)
	// 	dst := make([]byte, 8*b.N)
	// 	b.ResetTimer()
	// 	cgm, _ := cipher.NewGCM(prf.AES_Key_Fixed)
	// 	cgm.Seal(dst, prf.IV[:12], src, nil)
	// })
	b.Run("ECB-Normal", func(b *testing.B) {
		n := 16
		src := make([]byte, n)
		dst := make([]byte, n)
		b.ResetTimer()
		ecb := openssl.NewECBEncrypter(prf.AES_Key_Fixed)
		for i := 0; i < b.N; i++ {
			ecb.CryptBlocks(dst, src)
		}
	})
	// b.Run("ECB-Batch", func(b *testing.B) {
	// 	src := make([]byte, 8*b.N)
	// 	_, _ = openssl.AesECBEncrypt(src, prf.IV, openssl.PKCS7_PADDING)
	// })
}

func TestHashValues(t *testing.T) {
	n := 1
	src0 := make([]byte, n)
	src1 := make([]byte, n)
	rand.Read(src0)
	rand.Read(src1)
	fmt.Println(src0, src1)
	dst0 := make([]byte, n)
	dst1 := make([]byte, n)
	ote.HashValuesJ(dst0, src0, 1, n)
	ote.HashValuesJ(dst1, src1, 1, n)
	fmt.Println(dst0)
	fmt.Println(dst1)
	ote.HashValuesJ(dst0, src0, 1, n)
	ote.HashValuesJ(dst1, src1, 1, n)
	fmt.Println(dst0)
	fmt.Println(dst1)
	fmt.Println(16 - 15%16)
}
func TestOTE(t *testing.T) {
	n := 128
	bytes := 1
	x0 := make([]byte, n*bytes)
	x1 := make([]byte, n*bytes)
	x := make([]byte, n*bytes)
	c := make([]bool, n)
	for i := 0; i < n; i++ {
		c[i] = misc.Bool()
	}
	rand.Read(x0)
	rand.Read(x1)
	ser, cli := Nets()
	Parallel(
		func() {
			ote.InitOtSender(ser)
			ote.OTE.SendN(ser, x0, x1, n, bytes)
		},
		func() {
			ote.InitOtReceiver(cli)
			ote.OTE.RecvN(cli, x, c, n, bytes)
		})
	for i := range make([]bool, 128) {
		if ote.OtInit_Sbs[i] {
			assert.EqualValues(t, ote.OtInit_K[i], ote.OtInit_K1[i])
		} else {
			assert.EqualValues(t, ote.OtInit_K[i], ote.OtInit_K0[i])
		}
		if i%100 == 0 {
			fmt.Println("--- --- --- Sample", i/10, "--- --- ---")
			if c[i] {
				fmt.Println(ote.OtInit_K1[i])
			} else {
				fmt.Println(ote.OtInit_K0[i])
			}
			fmt.Println(ote.OtInit_K[i])
		}
	}
	for i := range make([]bool, n) {
		if c[i] {
			//assert.EqualValues(t, ote.OtInit_K[i].Int(), ote.OtInit_K0[i].Int())
			assert.EqualValues(t, x1[i*bytes:(i+1)*bytes], x[i*bytes:(i+1)*bytes])
		} else {
			//assert.EqualValues(t, ote.OtInit_K[i].Int(), ote.OtInit_K1[i].Int())
			assert.EqualValues(t, x0[i*bytes:(i+1)*bytes], x[i*bytes:(i+1)*bytes])
		}
		if i%100 == 0 {
			fmt.Println("--- --- --- Sample", i/10, "--- --- ---")
			if c[i] {
				fmt.Println(x1[i*bytes : (i+1)*bytes])
			} else {
				fmt.Println(x0[i*bytes : (i+1)*bytes])
			}
			fmt.Println(x[i*bytes : (i+1)*bytes])
		}
	}
	_ = ser
	_ = cli
}
func BenchmarkOTE(b *testing.B) {
	if b.N < 128 {
		return
	}
	bytes := 1
	x0 := make([]byte, b.N*bytes)
	x1 := make([]byte, b.N*bytes)
	x := make([]byte, b.N*bytes)
	c := make([]bool, b.N)
	for i := 0; i < b.N; i++ {
		c[i] = misc.Bool()
	}
	rand.Read(x0)
	rand.Read(x1)
	ser, cli := Nets()
	Parallel(
		func() {
			ote.InitOtSender(ser)
		},
		func() {
			ote.InitOtReceiver(cli)
		})
	b.ResetTimer()
	Parallel(
		func() {
			ote.OTE.SendN(ser, x0, x1, b.N, bytes)
		},
		func() {
			ote.OTE.RecvN(cli, x, c, b.N, bytes)
		})
}

func TestCOT(t *testing.T) {
	n := 5000
	bytes := 2
	x0 := make([]byte, n*bytes)
	x1 := make([]byte, n*bytes)
	x := make([]byte, n*bytes)
	c := make([]bool, n)
	ote.COT.F = func(args ...interface{}) {
		n := args[3].(int)
		bp := (args[0].([]byte))
		dst := unsafe.Slice((*int16)(unsafe.Pointer(&bp[0])), n)
		fmt.Printf("x1: %p %p %p\n", x1, bp, dst)
		bp = (args[1].([]byte))
		src := unsafe.Slice((*int16)(unsafe.Pointer(&bp[0])), n)
		ld := args[2].(int)
		fmt.Printf("ld: %v\n", ld)
		for i := 0; i < n; i++ {
			//fmt.Printf("x1[i]: %p %p %p\n", &x1[i*2], &bp[i*2], &dst[i])
			if uintptr(unsafe.Pointer(&x1[i*2])) != uintptr(unsafe.Pointer(&dst[i])) {
				panic("?")
			}
			dst[i] = 1<<(15-ld%16) + src[i]

			ld++
		}
	}
	for i := 0; i < n; i++ {
		c[i] = misc.Bool()
	}
	ser, cli := Nets()
	Parallel(
		func() {
			ote.InitOtSender(ser)
			ote.COT.SendN(ser, x0, x1, n, bytes)
		},
		func() {
			ote.InitOtReceiver(cli)
			ote.COT.RecvN(cli, x, c, n, bytes)
		})
	for i := range make([]bool, 128) {
		if ote.OtInit_Sbs[i] {
			assert.EqualValues(t, ote.OtInit_K[i], ote.OtInit_K1[i])
		} else {
			assert.EqualValues(t, ote.OtInit_K[i], ote.OtInit_K0[i])
		}
		if i%100 == 0 {
			fmt.Println("--- --- --- Sample", i/10, "--- --- ---")
			if c[i] {
				fmt.Println(ote.OtInit_K1[i])
			} else {
				fmt.Println(ote.OtInit_K0[i])
			}
			fmt.Println(ote.OtInit_K[i])
		}
	}
	for i := range make([]bool, n) {
		if c[i] {
			//assert.EqualValues(t, x1[i*bytes:(i+1)*bytes], x[i*bytes:(i+1)*bytes])
		} else {
			assert.EqualValues(t, x0[i*bytes:(i+1)*bytes], x[i*bytes:(i+1)*bytes])
		}
		if i%100 == 0 {
			fmt.Println("--- --- --- Sample", i/10, "--- --- ---")
			if c[i] {
				fmt.Println(c[i], "\n", x1[i*bytes:(i+1)*bytes])
			} else {
				fmt.Println(c[i], "\n", x0[i*bytes:(i+1)*bytes])
			}
			fmt.Println(x[i*bytes : (i+1)*bytes])
		}
	}
	_ = ser
	_ = cli
}

func BenchmarkCOT(b *testing.B) {
	if b.N < 128 {
		return
	}
	bytes := 1
	x0 := make([]byte, b.N*bytes)
	x1 := make([]byte, b.N*bytes)
	x := make([]byte, b.N*bytes)
	c := make([]bool, b.N)
	for i := 0; i < b.N; i++ {
		c[i] = misc.Bool()
	}
	rand.Read(x0)
	rand.Read(x1)
	ser, cli := Nets()
	b.ResetTimer()
	Parallel(
		func() {
			//baseot.Send128(ser, x0, x1)
			//ote.SendN(ser, x0, x1)
			ote.InitOtSender(ser)
			ote.COT.SendN(ser, x0, x1, b.N, bytes)
		},
		func() {
			//baseot.Recv128(cli, bo)
			//ote.RecvN(cli, c)
			ote.InitOtReceiver(cli)
			ote.COT.RecvN(cli, x, c, b.N, bytes)
		})
}
