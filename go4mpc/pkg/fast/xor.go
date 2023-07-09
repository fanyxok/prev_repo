package fast

import "unsafe"

func XorBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if n == 0 {
		return 0
	}
	_ = dst[n-1]
	xorBytesSSE2(&dst[0], &a[0], &b[0], n) // amd64 must have SSE2
	return n
}

//go:noescape
func xorBytesSSE2(dst, a, b *byte, n int)

// a,b,dst all have length 16
func Xor16(dst, a, b []byte) {
	dw := (*[2]uintptr)(unsafe.Pointer(&dst[0]))
	aw := (*[2]uintptr)(unsafe.Pointer(&a[0]))
	bw := (*[2]uintptr)(unsafe.Pointer(&b[0]))
	dw[0] = aw[0] ^ bw[0]
	dw[1] = aw[1] ^ bw[1]
}

func Xor8(dst, a, b []byte) {
	dw := (*[1]uintptr)(unsafe.Pointer(&dst[0]))
	aw := (*[1]uintptr)(unsafe.Pointer(&a[0]))
	bw := (*[1]uintptr)(unsafe.Pointer(&b[0]))
	dw[0] = aw[0] ^ bw[0]
}

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func Xor(dst, a, b []byte, n int) {
	// Assert dst has enough space
	_ = dst[n-1]

	w := n / wordSize
	if w > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		aw := *(*[]uintptr)(unsafe.Pointer(&a))
		bw := *(*[]uintptr)(unsafe.Pointer(&b))
		for i := 0; i < w; i++ {
			dw[i] = aw[i] ^ bw[i]
		}
	}

	for i := (n - n%wordSize); i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
}

func XorBatch128(dst, a, b []byte, repeat int) {
	// Assert dst has enough space
	_ = dst[16*repeat-1]
	dw := *(*[]uintptr)(unsafe.Pointer(&dst))
	aw := *(*[]uintptr)(unsafe.Pointer(&a))
	bw := *(*[]uintptr)(unsafe.Pointer(&b))
	for i := 0; i < repeat; i++ {
		dw[i*2] = aw[i*2] ^ bw[0]
		dw[i*2+1] = aw[i*2+1] ^ bw[1]
	}
}
