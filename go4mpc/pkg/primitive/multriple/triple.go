package multriple

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/internal/ot/ote"
	"s3l/mpcfgo/pkg/fast"
	"time"
	"unsafe"
)

var (
	rTypeInt8  = reflect.TypeOf(int8(0))
	rTypeInt16 = reflect.TypeOf(int16(0))
	rTypeInt32 = reflect.TypeOf(int32(0))
	rTypeInt64 = reflect.TypeOf(int64(0))
)

// int8 | int16 | int32 | int64
type MtVector[T ~int8 | ~int16 | ~int32 | ~int64] struct {
	// A || B || C
	Data []T
	Cap  int
	Next int
}

func NewMtVector[T ~int8 | ~int16 | ~int32 | ~int64](n int) *MtVector[T] {
	if n <= 0 {
		return nil
	}
	mtv := &MtVector[T]{
		Data: make([]T, n*3),
		Cap:  n,
		Next: 0,
	}
	return mtv
}

func (ct *MtVector[T]) ABC() (a, b, c []T) {
	return ct.Data[:ct.Cap], ct.Data[ct.Cap : ct.Cap*2], ct.Data[ct.Cap*2:]
}

func (ct *MtVector[T]) GetNext() (a, b, c T) {
	if ct.Next >= ct.Cap {
		log.Panicf("MtVector exceed Cap! (cap: %d, next: %d)\n", ct.Cap, ct.Next)
		return
	}
	a, b, c = ct.Data[ct.Next], ct.Data[ct.Cap+ct.Next], ct.Data[2*ct.Cap+ct.Next]
	ct.Next++
	return
}

func (mtv *MtVector[T]) Generate(net network.Network) {
	var t T
	switch reflect.TypeOf(t) {
	case rTypeInt8:
		genMtInt8(net, (*MtVector[int8])(unsafe.Pointer(mtv)))
	case rTypeInt16:
		genMtInt16(net, (*MtVector[int16])(unsafe.Pointer(mtv)))
	case rTypeInt32:
		//mtv.generateInt32(net)
	case rTypeInt64:
		//mtv.generateInt64(net)
	}
}

func genMtInt8(net network.Network, mtv *MtVector[int8]) {
	log.Printf("Generate Triple [int8] %v\n", mtv.Cap)
	ote.COT.EleSize = 1
	if net.Server {
		// a0, b0
		nOTs := mtv.Cap * 8
		A, B, C := mtv.ABC()
		A0 := unsafe.Slice((*byte)(unsafe.Pointer(&A[0])), mtv.Cap)
		B0 := unsafe.Slice((*byte)(unsafe.Pointer(&B[0])), mtv.Cap)
		//A0 := (*[]byte)(unsafe.Pointer(&A))
		//B0 := (*[]byte)(unsafe.Pointer(&B))
		rand.Read(A0)
		rand.Read(B0)
		X0 := make([]byte, nOTs)
		X1 := make([]byte, nOTs)
		ote.COT.F = func(args ...interface{}) {
			n := args[3].(int)
			phead := *(*[]int8)(fast.PointerOf(args[0]))
			dst := unsafe.Slice((*int8)(unsafe.Pointer(&phead[0])), n)
			phead = *(*[]int8)(fast.PointerOf(args[1]))
			src := unsafe.Slice((*int8)(unsafe.Pointer(&phead[0])), n)
			//dst := *(*[]int8)(fast.PointerOf(args[0]))
			//src := *(*[]int8)(fast.PointerOf(args[1]))
			ld := args[2].(int)
			for i := 0; i < n; i++ {
				(dst)[i] = A[ld/8]<<(7-ld%8) + (src)[i]
				ld++
			}
		}
		t := time.Now()
		ote.COT.SendN(net, X0, X1, nOTs, 1)
		fmt.Printf("Mtv-COT-1: %v\n", time.Since(t))

		for i := 0; i < mtv.Cap; i++ {
			c := byte(0)
			for j := 0; j < 8; j++ {
				c -= X0[i*8+j]
			}
			C[i] = A[i]*B[i] + int8(c)
		}
		ote.COT.F = func(args ...interface{}) {
			dst := *(*[]int8)(fast.PointerOf(args[0]))
			src := *(*[]int8)(fast.PointerOf(args[1]))
			ld := args[2].(int)
			n := args[3].(int)
			for i := 0; i < n; i++ {
				(dst)[i] = (B)[ld/8]<<(7-ld%8) + (src)[i]
				ld++
			}
		}
		t = time.Now()
		ote.COT.SendN(net, X0, X1, nOTs, 1)
		fmt.Printf("Mtv-COT-0: %v\n", time.Since(t))
		// // c0 :=  a0b0 + (a0b1)0 + (a1b0)0
		for i := 0; i < mtv.Cap; i++ {
			c := byte(0)
			for j := 0; j < 8; j++ {
				c -= X0[i*8+j]
			}
			C[i] += int8(c)
		}
	} else {
		// a1, b1
		nOTs := mtv.Cap * 8
		A, B, C := mtv.ABC()
		A1 := unsafe.Slice((*byte)(unsafe.Pointer(&A[0])), mtv.Cap)
		B1 := unsafe.Slice((*byte)(unsafe.Pointer(&B[0])), mtv.Cap)
		rand.Read(A1)
		rand.Read(B1)
		Xi := make([]byte, nOTs)
		R := make([]bool, nOTs)
		fast.Bytes2Bools(R, B1)
		ote.COT.RecvN(net, Xi, R, nOTs, 1)
		for i := 0; i < mtv.Cap; i++ {
			c := byte(0)
			for j := 0; j < 8; j++ {
				c += Xi[i*8+j]
			}
			C[i] = A[i]*B[i] + int8(c)
		}
		fast.Bytes2Bools(R, A1)
		ote.COT.RecvN(net, Xi, R, nOTs, 1)
		for i := 0; i < mtv.Cap; i++ {
			c := byte(0)
			for j := 0; j < 8; j++ {
				c += Xi[i*8+j]
			}
			C[i] += int8(c)
		}
	}
}

func Int16sToBools(dst []bool, src []int16, nint int) {
	for i := 0; i < nint; i++ {
		s := src[i]
		for j := 15; j >= 0; j-- {
			dst[16*i+j] = s&0x1 == 1
			s >>= 1
		}
	}
}

func genMtInt16(net network.Network, mtv *MtVector[int16]) {
	log.Printf("Generate Triple [int16] %v\n", mtv.Cap)
	ote.COT.EleSize = 16
	runtime.GC()
	const debug = 250
	if net.Server {
		// a0, b0
		nOTs := mtv.Cap * 16
		A, B, C := mtv.ABC()
		A0 := unsafe.Slice((*byte)(unsafe.Pointer(&A[0])), mtv.Cap*2)
		B0 := unsafe.Slice((*byte)(unsafe.Pointer(&B[0])), mtv.Cap*2)
		rand.Read(A0)
		rand.Read(B0)
		X0 := make([]byte, 2*nOTs)
		X1 := make([]byte, 2*nOTs)
		ote.COT.F = func(args ...interface{}) {
			lI16 := args[2].(int) // in bytes
			nI16 := args[3].(int) // num of bytes
			dst := unsafe.Slice((*int16)(unsafe.Pointer(&(args[0].([]byte)[0]))), nI16)
			src := unsafe.Slice((*int16)(unsafe.Pointer(&(args[1].([]byte)[0]))), nI16)

			start := lI16 / 16
			end := start + nI16
			for j := start; j < end; j++ {
				ele := (A)[j]
				idx := (j - start) * 16
				for i := 0; i < 16; i++ {
					(dst)[idx+i] = ele<<(15-i) + (src)[idx+i]
				}
			}

			// if ld/16 == debug {
			// 	fmt.Println("S", A[ld/16])
			// }
			// for i := 0; i < n; i++ {
			// 	(dst)[i] = A[ld/16]<<(15-ld%16) + (src)[i]
			// 	if ld/16 == debug {
			// 		fmt.Println("S", dst[i], A[ld/16]<<(15-ld%16), src[i])
			// 	}
			// 	ld++
			// }
		}
		t := time.Now()
		ote.COT.SendN(net, X0, X1, nOTs, 2)
		fmt.Printf("Mtv-COT-1: %v\n", time.Since(t))

		X0P := unsafe.Slice((*int16)(unsafe.Pointer(&X0[0])), nOTs)
		X1P := unsafe.Slice((*int16)(unsafe.Pointer(&X1[0])), nOTs)
		runtime.GC()

		for i := 0; i < mtv.Cap; i++ {
			c := int16(0)
			for j := 0; j < 16; j++ {
				c -= X0P[i*16+j]
				if i == debug {
					fmt.Println("S", i, j, X0P[i*16+j], X1P[i*16+j])
				}
			}
			C[i] = A[i]*B[i] + int16(c)
		}
		ote.COT.F = func(args ...interface{}) {
			lI16 := args[2].(int) // in bytes
			nI16 := args[3].(int) // num of bytes
			dst := unsafe.Slice((*int16)(unsafe.Pointer(&(args[0].([]byte)[0]))), nI16)
			src := unsafe.Slice((*int16)(unsafe.Pointer(&(args[1].([]byte)[0]))), nI16)

			start := lI16 / 16
			end := start + nI16
			for j := start; j < end; j++ {
				ele := (B)[j]
				idx := (j - start) * 16
				for i := 0; i < 16; i++ {
					(dst)[idx+i] = ele<<(15-i) + (src)[idx+i]
				}
			}

			//for i := 0; i < n; i++ {
			//	(dst)[i] = (B)[ld/16]<<(15-ld%16) + (src)[i]
			//	ld++
			//}
		}
		t = time.Now()
		ote.COT.SendN(net, X0, X1, nOTs, 2)
		fmt.Printf("Mtv-COT-0: %v\n", time.Since(t))
		// // c0 :=  a0b0 + (a0b1)0 + (a1b0)0
		for i := 0; i < mtv.Cap; i++ {
			c := int16(0)
			for j := 0; j < 16; j++ {
				c -= X0P[i*16+j]
			}
			C[i] += int16(c)
		}
	} else {
		// a1, b1
		nOTs := mtv.Cap * 16
		A, B, C := mtv.ABC()

		A1 := unsafe.Slice((*byte)(unsafe.Pointer(&A[0])), mtv.Cap*2)
		B1 := unsafe.Slice((*byte)(unsafe.Pointer(&B[0])), mtv.Cap*2)
		rand.Read(A1)
		rand.Read(B1)
		Xi := make([]byte, 2*nOTs)
		R := make([]bool, nOTs)
		Int16sToBools(R, B, mtv.Cap)

		//		fast.Bytes2Bools(R, B1)
		ote.COT.RecvN(net, Xi, R, nOTs, 2)
		XiP := unsafe.Slice((*int16)(unsafe.Pointer(&Xi[0])), nOTs)
		for i := 0; i < mtv.Cap; i++ {
			c := int16(0)
			if i == debug {
				fmt.Printf("%v %#016b\n", B[debug], uint16(B[debug]))
				fmt.Println()
			}
			for j := 0; j < 16; j++ {
				c += XiP[i*16+j]
				if i == debug {
					if R[i*16+j] {
						fmt.Println("R", 1, XiP[i*16+j])
					} else {
						fmt.Println("R", 0, XiP[i*16+j])
					}
				}
			}
			C[i] = A[i]*B[i] + c
		}
		Int16sToBools(R, A, mtv.Cap)
		//fast.Bytes2Bools(R, A1)
		ote.COT.RecvN(net, Xi, R, nOTs, 2)
		XiP = unsafe.Slice((*int16)(unsafe.Pointer(&Xi[0])), nOTs)
		for i := 0; i < mtv.Cap; i++ {
			c := int16(0)
			for j := 0; j < 16; j++ {
				c += XiP[i*16+j]
			}
			C[i] += c
		}
	}
}

// func genMtInt8(net network.Network, mtv *MtVector[int8]) {
// 	log.Printf("Generate Triple [int8] %v\n", mtv.Cap)
// 	// fi := func(base, bias int8, shift int) int8 {
// 	// 	//fmt.Println("base:", base, "bias:", bias, "shift:", shift)
// 	// 	return base<<shift - bias
// 	// }
// 	if net.Server {
// 		// a0, b0
// 		nOTs := mtv.Cap * 8
// 		A, B, C := mtv.ABC()
// 		A0 := (*[]byte)(unsafe.Pointer(&A))
// 		B0 := (*[]byte)(unsafe.Pointer(&B))
// 		rand.Read(*A0)
// 		rand.Read(*B0)
// 		X0 := make([]byte, nOTs)
// 		X1 := make([]byte, nOTs)
// 		ote.COT.F = func(args ...interface{}) {
// 			dst := args[0].([]byte)
// 			src := args[1].([]byte)
// 			ld := args[2].(int)
// 			n := args[3].(int)
// 			for i := 0; i < n; i++ {
// 				dst[i] = (*A0)[ld/8]<<(7-ld%8) + src[i]
// 				ld++
// 			}
// 		}
// 		t := time.Now()
// 		ote.COT.SendN(net, X0, X1, nOTs, 1)
// 		fmt.Printf("Mtv-COT-1: %v\n", time.Since(t))

// 		for i := 0; i < mtv.Cap; i++ {
// 			c := byte(0)
// 			for j := 0; j < 8; j++ {
// 				c -= X0[i*8+j]
// 			}
// 			C[i] = A[i]*B[i] + int8(c)
// 		}
// 		ote.COT.F = func(args ...interface{}) {
// 			dst := args[0].([]byte)
// 			src := args[1].([]byte)
// 			ld := args[2].(int)
// 			n := args[3].(int)
// 			for i := 0; i < n; i++ {
// 				dst[i] = (*B0)[ld/8]<<(7-ld%8) + src[i]
// 				ld++
// 			}
// 		}
// 		t = time.Now()
// 		ote.COT.SendN(net, X0, X1, nOTs, 1)
// 		fmt.Printf("Mtv-COT-0: %v\n", time.Since(t))
// 		// // c0 :=  a0b0 + (a0b1)0 + (a1b0)0
// 		for i := 0; i < mtv.Cap; i++ {
// 			c := byte(0)
// 			for j := 0; j < 8; j++ {
// 				c -= X0[i*8+j]
// 			}
// 			C[i] += int8(c)
// 		}
// 	} else {
// 		// a1, b1
// 		nOTs := mtv.Cap * 8
// 		A, B, C := mtv.ABC()
// 		A1 := (*[]byte)(unsafe.Pointer(&A))
// 		B1 := (*[]byte)(unsafe.Pointer(&B))
// 		rand.Read(*A1)
// 		rand.Read(*B1)
// 		Xi := make([]byte, nOTs)
// 		R := make([]bool, nOTs)
// 		fast.Bytes2Bools(R, *B1)
// 		ote.COT.RecvN(net, Xi, R, nOTs, 1)
// 		for i := 0; i < mtv.Cap; i++ {
// 			c := byte(0)
// 			for j := 0; j < 8; j++ {
// 				c += Xi[i*8+j]
// 			}
// 			C[i] = A[i]*B[i] + int8(c)
// 		}
// 		fast.Bytes2Bools(R, *A1)
// 		ote.COT.RecvN(net, Xi, R, nOTs, 1)
// 		for i := 0; i < mtv.Cap; i++ {
// 			c := byte(0)
// 			for j := 0; j < 8; j++ {
// 				c += Xi[i*8+j]
// 			}
// 			C[i] += int8(c)
// 		}
// 	}
// }
