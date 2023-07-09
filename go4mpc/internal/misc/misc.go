package misc

import (
	"encoding/binary"
	"fmt"
	"s3l/mpcfgo/pkg/always"
	"unsafe"
)

func BoolsToBytes(t []bool) []byte {
	b := make([]byte, (len(t)+7)/8)
	for i, x := range t {
		if x {
			b[i/8] |= 0x1 << uint(i%8)
		}
	}
	return b
}

func BytesToBools(b []byte) []bool {
	t := make([]bool, 8*len(b))
	for i, x := range b {
		for j := 0; j < 8; j++ {
			if (x>>uint(j))&0x1 == 0x1 {
				t[8*i+j] = true
			}
		}
	}
	return t
}

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func bytesXorBytes(dst, a, b []byte) []byte {
	n := len(dst)
	w := n / wordSize
	if w > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		aw := *(*[]uintptr)(unsafe.Pointer(&a))
		bw := *(*[]uintptr)(unsafe.Pointer(&b))
		_ = aw[w-1]
		_ = bw[w-1]
		_ = dw[w-1]
		for i := 0; i < w; i++ {
			dw[i] = aw[i] ^ bw[i]
		}
	}

	for i := (n - n%wordSize); i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}

func BytesXorBytes(dst, a, b []byte) []byte {
	if len(a) != len(b) {
		s := fmt.Sprintln("len() not equal. lhs len", len(a), "rhs len", len(b), "dst len", len(dst))
		panic(s)
	}
	always.Eq(len(a), len(b))
	bytesXorBytes(dst, a, b)
	return dst
}

func BytesXorBytes0(a, b []byte) []byte {
	var short []byte
	var long []byte
	if len(a) < len(b) {
		short = a
		long = b
	} else {
		short = b
		long = a
	}
	ls, ll := len(short), len(long)
	if ls == 0 {
		return nil
	}
	if ls == ll {
		dst := make([]byte, ls)
		bytesXorBytes(dst, short, long)
		return dst
	} else {
		dst := make([]byte, ls, ll)
		bytesXorBytes(dst, short, long[:ls])
		copy(dst[ls:], long[ls:])
		return dst
	}
}

func EncodeInt64[T ~int64 | ~uint64](v T) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func EncodeInt32[T ~int | ~int32 | ~uint32](v T) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(v))
	return b
}

func EncodeInt16[T ~int16 | ~uint16](v T) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v))
	return b
}

func EncodeInt8[T ~int8 | ~uint8](v T) []byte {
	b := make([]byte, 1)
	b[0] = byte(v)
	return b
}

func DecodeInt64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func DecodeInt32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func DecodeInt16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func DecodeInt8(b []byte) uint8 {
	return uint8(b[0])
}

func EncodeInt(x interface{}) {
	switch v := x.(type) {
	case int64:
		EncodeInt64(v)
	case uint64:
		EncodeInt64(v)
	case int32:

	default:

	}
}

// assume n >= 0
func IntPow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := IntPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}

func BigEndian2LittleEndian(b []byte, n int) []byte {
	d := make([]byte, n)

	for i, j := n-len(b), 0; i < n/2; i++ {
		d[i], d[n-i-1] = b[len(b)-j-1], b[j]
		j++
	}
	return d
}
