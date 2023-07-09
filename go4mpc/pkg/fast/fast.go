package fast

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"sync"
	"time"
	"unsafe"
)

type Eface struct {
	Typ, Val unsafe.Pointer
}

func PointerOf(x any) unsafe.Pointer {
	return (*Eface)(unsafe.Pointer(&x)).Val
}

func ShiftLeft128(high, low uint64, shift int) (uint64, uint64) {
	if shift < 64 {
		return high<<shift | low>>(64-shift), low << shift
	} else {
		return low << (shift - 64), 0
	}
}

func ShiftRight128(high, low uint64, shift int) (uint64, uint64) {
	if shift < 64 {
		hi := high >> shift
		rem := high << (64 - shift)
		lo := low>>shift | rem
		return hi, lo
	} else {
		return 0, high >> (shift - 64)
	}
}

func Transpose8(src *[8]byte) {
	x := binary.BigEndian.Uint64(src[:])
	x = x&0xAA55AA55AA55AA55 |
		(x&0x00AA00AA00AA00AA)<<7 |
		(x>>7)&0x00AA00AA00AA00AA
	x = x&0xCCCC3333CCCC3333 |
		(x&0x0000CCCC0000CCCC)<<14 |
		(x>>14)&0x0000CCCC0000CCCC
	x = x&0xF0F0F0F00F0F0F0F |
		(x&0x00000000F0F0F0F0)<<28 |
		(x>>28)&0x00000000F0F0F0F0
	binary.BigEndian.PutUint64(src[:], x)
}

// 32 * 32 bit matrix, transpose in place
func Transpose32(src *[128]byte) {
	var X [32]uint32
	for i := 0; i < 32; i++ {
		X[i] = binary.BigEndian.Uint32(src[i*4 : (i+1)*4])
	}
	var m uint32 = 0x0000ffff
	for j := 16; j != 0; {
		for k := 0; k < 32; k = (k + j + 1) & (^j) {
			t := (X[k] ^ (X[k+j] >> j)) & m
			X[k] ^= t
			X[k+j] ^= t << j
		}
		j >>= 1
		m = m ^ (m << j)
	}
	for i := 0; i < 32; i++ {
		binary.BigEndian.PutUint32(src[i*4:(i+1)*4], X[i])
	}
}

func Transpose64(src *[512]byte) {
	var X [64]uint64
	for i := 0; i < 64; i++ {
		X[i] = binary.BigEndian.Uint64(src[i*8 : (i+1)*8])
	}
	var m uint64 = 0x00000000ffffffff
	for j := 32; j != 0; {
		for k := 0; k < 64; k = (k + j + 1) & (^j) {
			t := (X[k] ^ (X[k+j] >> j)) & m
			X[k] ^= t
			X[k+j] ^= t << j
		}
		j >>= 1
		m = m ^ (m << j)
	}
	for i := 0; i < 64; i++ {
		binary.BigEndian.PutUint64(src[i*8:(i+1)*8], X[i])
	}
}

func Transpose128(src *[2048]byte) {
	var low [128]uint64
	var high [128]uint64
	for i := 0; i < 128; i++ {
		head := i * 16
		high[i] = binary.BigEndian.Uint64(src[head : head+8])
		low[i] = binary.BigEndian.Uint64(src[head+8 : head+16])
	}
	var m_high uint64 = 0x0
	var m_low uint64 = 0xffffffffffffffff
	for j := 64; j != 0; {
		for k := 0; k < 128; k = (k + j + 1) & (^j) {
			kj_high, kj_low := ShiftRight128(high[k+j], low[k+j], j)
			t_high := (high[k] ^ kj_high) & m_high
			t_low := (low[k] ^ kj_low) & m_low
			high[k] ^= t_high
			low[k] ^= t_low
			t_high, t_low = ShiftLeft128(t_high, t_low, j)
			high[k+j] ^= t_high
			low[k+j] ^= t_low
		}
		j >>= 1
		m_high_, m_low_ := ShiftLeft128(m_high, m_low, j)
		m_high ^= m_high_
		m_low ^= m_low_
	}
	for i := 0; i < 128; i++ {
		head := i * 16
		binary.BigEndian.PutUint64(src[head:head+8], high[i])
		binary.BigEndian.PutUint64(src[head+8:head+16], low[i])
	}
}

func Transpose128n(dst []byte, src []byte, cols int) {
	block128 := cols / 16
	// for j := 0; j < block128; j++ {
	// 	for k := 0; k < 128; k++ {
	// 		copy(dst[2048*j+k*16:2048*j+(k+1)*16], src[k*cols+j*16:k*cols+(j+1)*16])
	// 	}
	// 	Transpose128((*[2048]byte)(dst[2048*j : 2048*(j+1)]))
	// }
	const workers = 4
	var wg sync.WaitGroup
	wg.Add(workers)
	for p := 0; p < workers; p++ {
		go func(p int) {
			defer wg.Done()
			for j := p; j < block128; j += workers {
				head := 2048 * j
				for k := 0; k < 128; k++ {
					copy(dst[head+k*16:head+(k+1)*16], src[k*cols+j*16:k*cols+(j+1)*16])
				}
				Transpose128((*[2048]byte)(dst[head : head+2048]))
			}
		}(p)
	}
	remind := cols - block128*16
	head := block128 * 2048
	rhead := block128 * 16
	//fmt.Println(remind, head, rhead)
	// 64 * 64
	if remind >= 8 {
		var t, u [512]byte
		for k := 0; k < 64; k++ {
			copy(t[k*8:(k+1)*8], src[k*cols+rhead:k*cols+rhead+8])
			copy(u[k*8:(k+1)*8], src[(k+64)*cols+rhead:(k+64)*cols+rhead+8])
		}
		Transpose64(&t)
		Transpose64(&u)
		for k := 0; k < 64; k++ {
			copy(dst[head+16*k:head+16*k+8], t[k*8:(k+1)*8])
			copy(dst[head+16*k+8:head+16*k+16], u[k*8:(k+1)*8])
		}
		remind = remind - 8
		head = head + 1024
		rhead = rhead + 8
	}
	// 32 * 32
	if remind >= 4 {
		var q, w, e, r [128]byte
		for k := 0; k < 32; k++ {
			l := 4 * k
			d := l + 4
			copy(q[l:d], src[k*cols+rhead:k*cols+rhead+4])
			copy(w[l:d], src[(k+32)*cols+rhead:(k+32)*cols+rhead+4])
			copy(e[l:d], src[(k+64)*cols+rhead:(k+64)*cols+rhead+4])
			copy(r[l:d], src[(k+96)*cols+rhead:(k+96)*cols+rhead+4])
		}
		Transpose32(&q)
		Transpose32(&w)
		Transpose32(&e)
		Transpose32(&r)
		for k := 0; k < 32; k++ {
			l := 4 * k
			d := l + 4
			copy(dst[head:head+4], q[l:d])
			copy(dst[head+4:head+8], w[l:d])
			copy(dst[head+8:head+12], e[l:d])
			copy(dst[head+12:head+16], r[l:d])
			head += 16
		}
		remind = remind - 4
		rhead = rhead + 4
		//fmt.Println(remind, head, rhead)
	}
	// 8 * 8
	for remind > 0 {
		var q [128]byte
		for k := 0; k < 8; k++ {
			q[k] = src[k*cols+rhead]
			q[k+8] = src[(k+8)*cols+rhead]
			q[k+16] = src[(k+16)*cols+rhead]
			q[k+24] = src[(k+24)*cols+rhead]
			q[k+32] = src[(k+32)*cols+rhead]
			q[k+40] = src[(k+40)*cols+rhead]
			q[k+48] = src[(k+48)*cols+rhead]
			q[k+56] = src[(k+56)*cols+rhead]
			q[k+64] = src[(k+64)*cols+rhead]
			q[k+72] = src[(k+72)*cols+rhead]
			q[k+80] = src[(k+80)*cols+rhead]
			q[k+88] = src[(k+88)*cols+rhead]
			q[k+96] = src[(k+96)*cols+rhead]
			q[k+104] = src[(k+104)*cols+rhead]
			q[k+112] = src[(k+112)*cols+rhead]
			q[k+120] = src[(k+120)*cols+rhead]
		}
		for k := 0; k < 16; k++ {
			Transpose8((*[8]byte)(q[k*8 : (k+1)*8]))
		}
		for k := 0; k < 8; k++ {
			dst[head] = q[k]
			dst[head+1] = q[k+8]
			dst[head+2] = q[k+16]
			dst[head+3] = q[k+24]
			dst[head+4] = q[k+32]
			dst[head+5] = q[k+40]
			dst[head+6] = q[k+48]
			dst[head+7] = q[k+56]
			dst[head+8] = q[k+64]
			dst[head+9] = q[k+72]
			dst[head+10] = q[k+80]
			dst[head+11] = q[k+88]
			dst[head+12] = q[k+96]
			dst[head+13] = q[k+104]
			dst[head+14] = q[k+112]
			dst[head+15] = q[k+120]
			head += 16
		}
		remind -= 1
		rhead += 1
	}
	wg.Wait()
}

func ZerosPadding(src []byte, padded int) []byte {
	return append(src, bytes.Repeat([]byte{byte(0)}, padded)...)
}

// n bools
func Bytes2Bools(dst []bool, src []byte) {
	fmt.Println(len(dst), len(src))
	for i := range dst {
		dst[i] = (src[i/8]>>(7-i%8))&0x1 != 0
	}
}

func Bytes2BoolsN(dst []bool, src []byte, nbits int) {

}
func Bools2Bytes(dst []byte, src []bool) {
	for i := range src {
		b := *(*byte)(unsafe.Pointer(&src[i]))
		dst[i/8] |= (b) << (7 - i%8)
	}
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Suppose rows is 128 or 256, or 8n
func MatricBitColAt8(dst []byte, src []byte, rows, cols, col int) {
	// dst contains (rows+7)/8 bytes
	cIdx := col / 8
	cBit := 7 - col%8 // low to high
	switch cBit {
	case 0:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := src[(i+7)*cols+cIdx] & 0x1
			t |= src[(i+6)*cols+cIdx] & 0x1 << 1
			t |= src[(i+5)*cols+cIdx] & 0x1 << 2
			t |= src[(i+4)*cols+cIdx] & 0x1 << 3
			t |= src[(i+3)*cols+cIdx] & 0x1 << 4
			t |= src[(i+2)*cols+cIdx] & 0x1 << 5
			t |= src[(i+1)*cols+cIdx] & 0x1 << 6
			t |= src[i*cols+cIdx] & 0x1 << 7
			dst[i/8] = t
		}
	case 1:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := src[i*cols+cIdx] << 6
			t |= src[(i+7)*cols+cIdx] >> 1
			t &= 0x81
			t |= src[(i+1)*cols+cIdx] & 0x2 << 5
			t |= src[(i+2)*cols+cIdx] & 0x2 << 4
			t |= src[(i+3)*cols+cIdx] & 0x2 << 3
			t |= src[(i+4)*cols+cIdx] & 0x2 << 2
			t |= src[(i+5)*cols+cIdx] & 0x2 << 1
			t |= src[(i+6)*cols+cIdx] & 0x2
			dst[i/8] = t
		}
	case 2:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := (src[i*cols+cIdx]<<5 | src[(i+6)*cols+cIdx]>>1) & 0x82
			t |= (src[(i+1)*cols+cIdx]<<4 | src[(i+7)*cols+cIdx]>>2) & 0x41
			t |= src[(i+2)*cols+cIdx] & 0x4 << 3
			t |= src[(i+3)*cols+cIdx] & 0x4 << 2
			t |= src[(i+4)*cols+cIdx] & 0x4 << 1
			t |= src[(i+5)*cols+cIdx] & 0x4
			t |= src[(i+7)*cols+cIdx] & 0x4 >> 2
			dst[i/8] = t
		}
	case 3:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := src[i*cols+cIdx] & 0x8 << 4
			t |= src[(i+1)*cols+cIdx] & 0x8 << 3
			t |= src[(i+2)*cols+cIdx] & 0x8 << 2
			t |= src[(i+3)*cols+cIdx] & 0x8 << 1
			t |= src[(i+4)*cols+cIdx] & 0x8
			t |= src[(i+5)*cols+cIdx] & 0x8 >> 1
			t |= src[(i+6)*cols+cIdx] & 0x8 >> 2
			t |= src[(i+7)*cols+cIdx] & 0x8 >> 3
			dst[i/8] = t
		}
	case 4:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := src[i*cols+cIdx] & 0x10 << 3
			t |= src[(i+1)*cols+cIdx] & 0x10 << 2
			t |= src[(i+2)*cols+cIdx] & 0x10 << 1
			t |= src[(i+3)*cols+cIdx] & 0x10
			t |= src[(i+4)*cols+cIdx] & 0x10 >> 1
			t |= src[(i+5)*cols+cIdx] & 0x10 >> 2
			t |= src[(i+6)*cols+cIdx] & 0x10 >> 3
			t |= src[(i+7)*cols+cIdx] & 0x10 >> 4
			dst[i/8] = t
		}
	case 5:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := src[i*cols+cIdx] & 0x20 << 2
			t |= src[(i+1)*cols+cIdx] & 0x20 << 1
			t |= src[(i+2)*cols+cIdx] & 0x20
			t |= src[(i+3)*cols+cIdx] & 0x20 >> 1
			t |= src[(i+4)*cols+cIdx] & 0x20 >> 2
			t |= src[(i+5)*cols+cIdx] & 0x20 >> 3
			t |= src[(i+6)*cols+cIdx] & 0x20 >> 4
			t |= src[(i+7)*cols+cIdx] & 0x20 >> 5
			dst[i/8] = t
		}
	case 6:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := src[i*cols+cIdx] & 0x40 << 1
			t |= src[(i+1)*cols+cIdx] & 0x40
			t |= src[(i+2)*cols+cIdx] & 0x40 >> 1
			t |= src[(i+3)*cols+cIdx] & 0x40 >> 2
			t |= src[(i+4)*cols+cIdx] & 0x40 >> 3
			t |= src[(i+5)*cols+cIdx] & 0x40 >> 4
			t |= src[(i+6)*cols+cIdx] & 0x40 >> 5
			t |= src[(i+7)*cols+cIdx] & 0x40 >> 6
			dst[i/8] = t
		}
	case 7:
		for i := 0; (i + 8) <= rows; i += 8 {
			t := src[i*cols+cIdx] & 0x80
			t |= src[(i+1)*cols+cIdx] & 0x80 >> 1
			t |= src[(i+2)*cols+cIdx] & 0x80 >> 2
			t |= src[(i+3)*cols+cIdx] & 0x80 >> 3
			t |= src[(i+4)*cols+cIdx] & 0x80 >> 4
			t |= src[(i+5)*cols+cIdx] & 0x80 >> 5
			t |= src[(i+6)*cols+cIdx] & 0x80 >> 6
			t |= src[(i+7)*cols+cIdx] & 0x80 >> 7
			dst[i/8] = t
		}
	}
}

// cols in bytes, col in bits
// dst has (rows+7)/8 bytes
func MatricBitColAt(dst []byte, src []byte, rows, cols, col int) {
	// dst contains (rows+7)/8 bytes
	cIdx := col / 8
	cBit := 7 - col%8
	for i := 0; (i + 8) <= rows; i += 8 {
		var t byte = 0
		for k := 0; k < 8; k++ {
			t |= (src[(i+k)*cols+cIdx] >> cBit) & 0x1 << (7 - k)
		}
		dst[i/8] = t
	}
	if i := rows / 8 * 8; i < rows {
		var t byte = 0
		for i < rows {
			t |= src[i*cols+cIdx] >> cBit << (7 - i%8)
			i++
		}
		dst[rows/8] = t
	}
}

// Deprecated: use MatricBitColAt or SimpleBitMatrixTranspose instead
// Eklundh Prim Algorithm. For my use case, suppose row = 128
func FineBitMatrixTranspose(dst []byte, src []byte, rows int, cols int) {
	if rows != 128 || cols < 128 {
		log.Panicf("BitMatrix has %v rows %v cols\n", rows, cols)
	}

	tmp := make([]byte, rows*cols)
	// transpose by bytes
	col2n := 1 << int(math.Log2(float64(cols)))
	t := time.Now()
	var wg sync.WaitGroup
	workers := 4
	wload := rows / workers
	for p := 0; p < workers; p++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			for i := p * wload; i < (p+1)*wload; i++ {
				for j := col2n; j < cols; j++ {
					tmp[j*rows+i] = src[i*cols+j]
				}
			}
		}(p)
	}
	wg.Wait()

	workload := rows / 2 / workers
	for p := 0; p < workers; p++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			start := p * workload
			end := MinInt(start+workload, rows/2)
			for i := start; i < end; i++ {
				up := i * 2
				for j := 0; j < col2n/2; j++ {
					left := j * 2
					dst[up*cols+left] = src[up*cols+left]
					dst[up*cols+left+1] = src[(up+1)*cols+left]
					dst[(up+1)*cols+left] = src[up*cols+left+1]
					dst[(up+1)*cols+left+1] = src[(up+1)*cols+left+1]
				}
			}
		}(p)
	}
	wg.Wait()

	unitWidth := []int{2, 4, 8, 16, 32, 64}
	blockWidth := []int{4, 8, 16, 32, 64, 128}
	for idx, bWidth := range blockWidth {
		uWidth := unitWidth[idx]
		rowN := rows / bWidth
		colN := col2n / bWidth
		t := make([]byte, uWidth)
		for i := 0; i < rowN; i++ {
			up := i * bWidth
			for j := 0; j < colN; j++ {
				left := j * bWidth
				for k := 0; k < uWidth; k++ {
					copy(t, dst[(up+k)*cols+left+uWidth:(up+k)*cols+left+uWidth*2])
					copy(dst[(up+k)*cols+left+uWidth:(up+k)*cols+left+uWidth*2], dst[(up+uWidth+k)*cols+left:(up+uWidth+k)*cols+left+uWidth])
					copy(dst[(up+uWidth+k)*cols+left:(up+uWidth+k)*cols+left+uWidth], t)
				}
			}
		}
	}

	// // DEBUG
	// for i := 0; i < rows; i++ {
	// 	fmt.Println("TMP", i, tmp[i*cols:i*cols+cols])
	// }
	colN := col2n / 128
	for i := 0; i < colN; i++ {
		for j := 0; j < 128; j++ {
			copy(tmp[i*16384+j*128:i*16384+j*128+128], dst[j*cols+i*128:j*cols+i*128+128])
		}
	}
	tt := time.Since(t)
	fmt.Println("Compute", tt)
	t = time.Now()
	for i := 0; i < cols*rows; i++ {
		dst[i] = 0
	}
	ts := time.Since(t)
	fmt.Println("Clean", ts)
	t = time.Now()

	wload = rows / workers
	for p := 0; p < workers; p++ {
		wg.Add(1)
		go func(p int) {
			st := p * wload
			ed := MinInt(st, rows)
			for j := st; j < ed; j++ {
				for i := 0; i < cols; i++ {
					for k := 0; k < 8; k++ {
						dst[(i*8+k)*16+j/8] |= ((tmp[i*rows+j] >> (7 - k)) & 0x1) << (7 - j%8)
					}
				}
			}
		}(p)
	}

	te := time.Since(t)
	fmt.Println("Compact", te)
}
func SimpleBitMatrixTranspose(dst []byte, src []byte, rows int, cols int) {
	rdiv := rows / 8
	for i := 0; i < rows; i++ {
		idiv := i / 8
		imod := i % 8
		for j := 0; j < cols; j++ {
			for k := 0; k < 8; k++ {
				dst[(j*8+k)*rdiv+idiv] |= ((src[i*cols+j] >> (7 - k)) & 1) << (7 - imod)
			}
		}
	}
}
