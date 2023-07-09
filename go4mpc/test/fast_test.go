package test

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"s3l/mpcfgo/pkg/fast"
	"strings"
	"sync"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestShiftLeft128(t *testing.T) {
	high := rand.Uint64()
	low := rand.Uint64()
	for i := 0; i < 128; i++ {
		hi, lo := fast.ShiftLeft128(high, low, i)
		fmt.Println(i)
		fmt.Printf("%064b%064b\n", high, low)
		fmt.Print(strings.Repeat(" ", i))
		fmt.Printf("%064b%064b\n", hi, lo)
		fmt.Println("===")
	}
}

func TestShiftRight128(t *testing.T) {
	high := rand.Uint64()
	low := rand.Uint64()
	for i := 0; i < 128; i++ {
		hi, lo := fast.ShiftRight128(high, low, i)
		fmt.Print(strings.Repeat(" ", i))
		fmt.Printf("%064b%064b\n", high, low)
		fmt.Printf("%064b%064b\n", hi, lo)
		fmt.Println("===")
	}
}
func TestTranspose8(t *testing.T) {
	rows := 8
	cols := 1
	src := make([]byte, rows*cols)
	rand.Read(src)
	expect := make([]byte, rows*cols)
	fast.SimpleBitMatrixTranspose(expect, src, rows, cols)
	fast.Transpose8((*[8]byte)(src))
	for i := 0; i < 8; i++ {
		assert.EqualValues(t, expect[i], src[i])
	}
}

func BenchmarkTranspose8(b *testing.B) {
	rows := 8
	cols := 1
	src := make([]byte, rows*cols)
	expect := make([]byte, rows*cols)
	rand.Read(src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fast.SimpleBitMatrixTranspose(expect, src, rows, cols)
	}
}

func TestTranspose32(t *testing.T) {
	rows := 32
	cols := 4
	src := make([]byte, rows*cols)

	rand.Read(src)

	expect := make([]byte, rows*cols)
	fast.SimpleBitMatrixTranspose(expect, src, rows, cols)

	fast.Transpose32((*[128]byte)(src))

	for i := 0; i < 32; i++ {
		assert.EqualValues(t, expect[i*4:(i+1)*4], src[i*4:(i+1)*4])
	}
}
func BenchmarkTranspose32(b *testing.B) {
	rows := 32
	cols := 4
	src := make([]byte, rows*cols)
	//expect := make([]byte, rows*cols)
	rand.Read(src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fast.Transpose32((*[128]byte)(src))
		//fast.SimpleBitMatrixTranspose(expect, src, rows, cols)
	}
}

func TestTranspose64(t *testing.T) {
	rows := 64
	cols := 8
	src := make([]byte, rows*cols)

	rand.Read(src)

	expect := make([]byte, rows*cols)
	fast.SimpleBitMatrixTranspose(expect, src, rows, cols)

	fast.Transpose64((*[512]byte)(src))

	for i := 0; i < 64; i++ {
		assert.EqualValues(t, expect[i*8:(i+1)*8], src[i*8:(i+1)*8])
	}
}
func BenchmarkTranspose64(b *testing.B) {
	rows := 64
	cols := 8
	src := make([]byte, rows*cols)
	expect := make([]byte, rows*cols)
	rand.Read(src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//fast.Transpose64((*[512]byte)(src))
		fast.SimpleBitMatrixTranspose(expect, src, rows, cols)
	}
}

func TestTranspose128(t *testing.T) {
	rows := 128
	cols := 16
	src := make([]byte, rows*cols)

	rand.Read(src)

	expect := make([]byte, rows*cols)
	fast.SimpleBitMatrixTranspose(expect, src, rows, cols)

	fast.Transpose128((*[2048]byte)(src))

	for i := 0; i < 128; i++ {
		assert.EqualValues(t, expect[i*16:(i+1)*16], src[i*16:(i+1)*16])
	}
}
func BenchmarkTranspose128(b *testing.B) {
	rows := 128
	cols := 16
	src := make([]byte, rows*cols)
	expect := make([]byte, rows*cols)
	rand.Read(src)
	b.Run("Transpose128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fast.Transpose128((*[2048]byte)(src))
		}
	})
	b.Run("SimpleBitMatrixTranspose", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fast.SimpleBitMatrixTranspose(expect, src, rows, cols)
		}
	})
	b.Run("MatricBitColAt8", func(b *testing.B) {
		expect := make([]byte, (rows+7)/8)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for i := 0; i < cols*8; i++ {
				fast.MatricBitColAt8(expect, src, rows, cols, i)
			}
		}
	})

}

func TestTranspose128n(t *testing.T) {
	rows := 128
	cols := 4500
	src := make([]byte, rows*cols)

	rand.Read(src)

	expect := make([]byte, rows*cols)
	act := make([]byte, rows*cols)
	fast.SimpleBitMatrixTranspose(expect, src, rows, cols)

	fast.Transpose128n(act, src, cols)

	for i := 0; i < cols*8; i++ {
		assert.EqualValues(t, expect[i*16:(i+1)*16], act[i*16:(i+1)*16])
	}
}

func BenchmarkTranspose128n(b *testing.B) {
	b.Run("Transpose128n", func(b *testing.B) {
		rows := 128
		cols := b.N
		if b.N < 128 {
			return
		}
		src := make([]byte, rows*cols)
		act := make([]byte, rows*cols)
		rand.Read(src)
		b.ResetTimer()
		fast.Transpose128n(act, src, cols)
	})
	b.Run("SimpleBitMatrixTranspose", func(b *testing.B) {
		rows := 128
		cols := b.N
		if b.N < 128 {
			return
		}
		src := make([]byte, rows*cols)
		act := make([]byte, rows*cols)
		rand.Read(src)
		b.ResetTimer()
		fast.SimpleBitMatrixTranspose(act, src, rows, cols)
	})
	b.Run("MatricBitColAt8", func(b *testing.B) {
		rows := 128
		cols := b.N
		if b.N < 128 {
			return
		}
		src := make([]byte, rows*cols)
		act := make([]byte, 16)
		rand.Read(src)
		b.ResetTimer()
		for i := 0; i < cols*8; i++ {
			fast.MatricBitColAt8(act, src, rows, cols, i)
		}
	})

}
func BenchmarkSimpleBitMatrixTranspose(b *testing.B) {
	rows := 128
	cols := b.N
	if b.N < 128 {
		return
	}
	src := make([]byte, rows*cols)
	rand.Read(src)
	expect := make([]byte, rows*cols)
	b.ResetTimer()
	fast.SimpleBitMatrixTranspose(expect, src, rows, cols)
}
func BenchmarkMatrixBitColAt(b *testing.B) {
	rows := 128
	cols := b.N
	if b.N < 128 {
		return
	}
	src := make([]byte, rows*cols)
	rand.Read(src)
	b.ResetTimer()
	var wg sync.WaitGroup
	for p := 0; p < 8; p++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			dst := make([]byte, (rows+7)/8)
			for i := p; i < cols*8; i += 8 {
				fast.MatricBitColAt8(dst, src, rows, cols, i)
			}
		}(p)
	}
	wg.Wait()
}
func TestMatrixBitColAt(t *testing.T) {
	rows := 128
	cols := 5600
	src := make([]byte, rows*cols)

	rand.Read(src)
	expect := make([]byte, rows*cols)
	fast.SimpleBitMatrixTranspose(expect, src, rows, cols)
	// for i := 0; i < cols*8; i++ {
	// 	fmt.Println(i, expect[i*((rows+7)/8):(i+1)*((rows+7)/8)])
	// }
	var wg sync.WaitGroup
	for p := 0; p < 4; p++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			dst := make([]byte, (rows+7)/8)
			for i := p; i < cols*8; i += 4 {
				fast.MatricBitColAt8(dst, src, rows, cols, i)
				assert.EqualValues(t, expect[i*((rows+7)/8):(i+1)*((rows+7)/8)], dst)
			}
		}(p)
	}
	wg.Wait()
	// for i := 0; i < cols*8; i++ {
	// 	fast.MatricBitColAt(dst, src, rows, cols, i)
	// 	assert.EqualValues(t, expect[i*((rows+7)/8):(i+1)*((rows+7)/8)], dst)
	// }
	_ = expect
}

func TestUnsafe(t *testing.T) {
	a := make([]byte, 8)
	rand.Read(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%p\n", &a[i])
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%08b", a[i])
	}
	fmt.Println()
	a0 := binary.LittleEndian.Uint64(a[:])
	a1 := binary.BigEndian.Uint64(a[:])
	fmt.Printf("%064b\n%064b\n", a0, a1)
	b := *(*uint64)(unsafe.Pointer((*[8]byte)(a)))
	fmt.Printf("%064b\n", b)
}

func BenchmarkXor(b *testing.B) {
	b.Run("Xor-Loop", func(b *testing.B) {
		bytes := 16
		dst := make([]byte, bytes)
		A := make([]byte, bytes)
		B := make([]byte, bytes)
		for i := 0; i < b.N; i++ {
			fast.Xor(dst, A, B, bytes)

		}
	})
	b.Run("Xor-Batch", func(b *testing.B) {
		bytes := 16
		dst := make([]byte, bytes*b.N)
		A := make([]byte, bytes*b.N)
		B := make([]byte, bytes*b.N)
		fast.Xor(dst, A, B, bytes)
	})
}
