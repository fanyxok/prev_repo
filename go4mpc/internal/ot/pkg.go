package ot

import (
	"fmt"
	"math/rand"
	"s3l/mpcfgo/pkg/always"
	"unsafe"
)

// get the ith col of mat
// unsafe style using of unsafe will cause unpredictable data change due to GC
// ensure using the safe style of unsafe
func Transpose0(mat [][]byte, idx int) []byte {
	nrow, ncol := len(mat), len(mat[0])
	if ncol*8 <= idx {
		panic(fmt.Sprintln("idx out of range, should within", ncol*8, "find", idx))
	}
	ret := make([]byte, (nrow+7)/8)
	for i := range mat {
		fmt.Println("Before", ret, "\nwill", mat[i][idx/8]>>(idx%8))
		rptr := (*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(&ret[0])) + uintptr(i)))
		*rptr = *(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat[i][0])) + uintptr(idx)))
		fmt.Println("After", ret)
	}
	return ret
}

// get the ith col of mat, where ith is count in bits
func Transpose(mat [][]byte, idx int) []byte {
	nrow, ncol := len(mat), len(mat[0])
	if ncol*8 <= idx {
		panic(fmt.Sprintln("idx out of range, should within", ncol*8, "find", idx))
	}
	ret := make([]byte, (nrow+7)/8)
	for i := range mat {
		b := mat[i][idx/8]
		ii := idx % 8
		bb := (b & (0x1 << ii)) >> ii
		ret[i/8] |= bb << (i % 8)
	}
	return ret
}

// len is the length in bits
func Remapping(in []byte, length int) []byte {
	l := (length + 7) / 8
	if len(in) == l {
		return in
	} else if len(in) > length {
		return in[:l]
	} else {
		dst := make([]byte, l)
		copy(dst, in)
		return dst
	}
}

// n is the bits length
func RandN(r *rand.Rand, n int) []byte {
	b := make([]byte, (n+7)/8)
	_, err := r.Read(b)
	always.Nil(err)
	return b
}
