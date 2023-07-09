package test

import (
	"fmt"
	"math/rand"
	"s3l/mpcfgo/internal/ot/ote"
	"s3l/mpcfgo/pkg/primitive/multriple"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestNewTriple(t *testing.T) {
	ser, cli := Nets()
	type tType = int16
	var mtvS, mtvR *multriple.MtVector[tType]
	n := 256
	Parallel(
		func() {
			ote.InitOtSender(ser)
			mtvS = multriple.NewMtVector[tType](n)
			mtvS.Generate(ser)

		},
		func() {
			ote.InitOtReceiver(cli)

			mtvR = multriple.NewMtVector[tType](n)
			mtvR.Generate(cli)

		})
	// v := byte(255)
	// u := (*int8)(unsafe.Pointer(&v))
	// *u = -1
	// fmt.Printf("%#08b %#08b\n", v, *u)
	// j := 1
	// fmt.Printf("C:[%v,%v]=%v,", mtvS.C[j], mtvR.C[j], mtvS.C[j]+mtvR.C[j])
	// fmt.Printf("A:[%v,%v]=%v,", mtvS.A[j], mtvR.A[j], mtvS.A[j]+mtvR.A[j])
	// fmt.Printf("B:[%v,%v]=%v,", mtvS.B[j], mtvR.B[j], mtvS.B[j]+mtvR.B[j])
	// fmt.Printf("C=A*B=%v\n", (mtvS.A[j]+mtvR.A[j])*(mtvS.B[j]+mtvR.B[j]))
	// fmt.Printf("A0*B1=%v\n", mtvS.A[j]*mtvR.B[j])
	// fmt.Printf("A1*B0=%v\n", mtvS.B[j]*mtvR.A[j])

	cnt := 1
	for i := 0; i < n; i++ {
		sa, sb, sc := mtvS.GetNext()
		ra, rb, rc := mtvR.GetNext()
		// 	// for i := 0; i < 128; i++ {
		// 	fmt.Printf("C:[%v,%v]=%v,", mtvS.C[i], mtvR.C[i], mtvS.C[i]+mtvR.C[i])
		// 	fmt.Printf("A:[%v,%v]=%v,", mtvS.A[i], mtvR.A[i], mtvS.A[i]+mtvR.A[i])
		// 	fmt.Printf("B:[%v,%v]=%v,", mtvS.B[i], mtvR.B[i], mtvS.B[i]+mtvR.B[i])
		// 	fmt.Printf("C=A*B=%v\n", (mtvS.A[i]+mtvR.A[i])*(mtvS.B[i]+mtvR.B[i]))
		// 	fmt.Printf("A0*B1=%v\n", mtvS.A[i]*mtvR.B[i])
		// 	fmt.Printf("A1*B0=%v\n", mtvS.B[i]*mtvR.A[i])
		// 	// }
		if !assert.EqualValues(t, (sa+ra)*(sb+rb), sc+rc, fmt.Sprintf("Errors %d at %d", cnt, i)) {
			cnt++
		}
		const sample = 10
		if i%sample == 0 {
			fmt.Println("--- --- --- Sample", i/sample, "-", i, "--- --- ---")
			fmt.Printf("A(%v+%v)%v*B(%v+%v)%v=%v=C(%v+%v)%v\n", sa, ra, sa+ra, sb, rb, sb+rb, (sa+ra)*(sb+rb), sc, rc, sc+rc)
		}
	}

}

func BenchmarkNewTriple(b *testing.B) {
	ser, cli := Nets()
	var mtvS, mtvR *multriple.MtVector[int8]

	Parallel(
		func() {
			ote.InitOtSender(ser)
			mtvS = multriple.NewMtVector[int8](b.N)
			mtvS.Generate(ser)

		},
		func() {
			ote.InitOtReceiver(cli)

			mtvR = multriple.NewMtVector[int8](b.N)
			mtvR.Generate(cli)

		})
	// for i := 0; i < 128; i++ {
	// 	fmt.Printf("C:[%v,%v]=%v,", mtvS.C[i], mtvR.C[i], mtvS.C[i]+mtvR.C[i])
	// 	fmt.Printf("A:[%v,%v]=%v,", mtvS.A[i], mtvR.A[i], mtvS.A[i]+mtvR.A[i])
	// 	fmt.Printf("B:[%v,%v]=%v,", mtvS.B[i], mtvR.B[i], mtvS.B[i]+mtvR.B[i])
	// 	fmt.Printf("%v\n", (mtvS.A[i]+mtvR.A[i])*(mtvS.B[i]+mtvR.B[i]))
	// }
}
func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	m := make([]int8, 1000000)
	bm := (*[]byte)(unsafe.Pointer(&m))
	n, err := rand.Read(*bm)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("n: %v\n", n)
}
