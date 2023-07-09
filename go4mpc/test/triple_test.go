package test

import (
	"fmt"
	"s3l/mpcfgo/pkg/primitive/triple"
	"s3l/mpcfgo/pkg/type/pub"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkTriple(b *testing.B) {
	// Sender input && Routine
	ser, cli := Nets()
	Parallel(
		func() {
			_ = triple.NewTriples(ser, pub.ZeroInt8, b.N)
		},
		func() {
			triple.NewTriples(cli, pub.ZeroInt8, b.N)
		})
}
func TestTriple(t *testing.T) {

	serNet, cliNet := Nets()
	N := 1000
	Parallel(
		func() {
			_ = triple.NewTriples(serNet, pub.ZeroBool, N)
		},
		func() {
			triple.NewTriples(cliNet, pub.ZeroBool, N)
		})

}
func TestTripleInt(t *testing.T) {

	serNet, cliNet := Nets()
	N := 1000
	tp0 := make(chan *triple.Triples, 10)
	tp1 := make(chan *triple.Triples, 10)
	Parallel(
		func() {
			tp0 <- triple.NewTriples(serNet, pub.ZeroInt8, N)
			tp0 <- triple.NewTriples(serNet, pub.ZeroInt16, N)
			tp0 <- triple.NewTriples(serNet, pub.ZeroInt32, N)
			tp0 <- triple.NewTriples(serNet, pub.ZeroInt64, N)

		},
		func() {
			tp1 <- triple.NewTriples(cliNet, pub.ZeroInt8, N)
			tp1 <- triple.NewTriples(cliNet, pub.ZeroInt16, N)
			tp1 <- triple.NewTriples(cliNet, pub.ZeroInt32, N)
			tp1 <- triple.NewTriples(cliNet, pub.ZeroInt64, N)

		})
	for range []int{0, 1, 2, 3} {
		TP0 := <-tp0
		TP1 := <-tp1
		for j := 0; j < N; j++ {
			assert.EqualValues(t, TP0.C[j].Add(TP1.C[j]), TP0.A[j].Add(TP1.A[j]).Mul(TP0.B[j].Add(TP1.B[j])))
			if j%Sample == 0 {
				fmt.Println(TP0.C[j].Add(TP1.C[j]), TP0.A[j].Add(TP1.A[j]), TP0.B[j].Add(TP1.B[j]))
			}
		}
	}
}
func XOR(x, y pub.PubNum) pub.Bool {
	return x.(pub.Bool) != y.(pub.Bool)
}
func TestTripleBool(t *testing.T) {

	serNet, cliNet := Nets()
	N := 10000
	tp0 := make(chan *triple.Triples, 10)
	tp1 := make(chan *triple.Triples, 10)
	Parallel(
		func() {
			tp0 <- triple.NewTriples(serNet, pub.ZeroBool, N)

		},
		func() {
			tp1 <- triple.NewTriples(cliNet, pub.ZeroBool, N)

		})
	for range []int{0} {
		TP0 := <-tp0
		TP1 := <-tp1
		for j := 0; j < N; j++ {
			//fmt.Println(j, "C", TP0.C[j], TP1.C[j], XOR(TP0.C[j], TP1.C[j]), "A", TP0.A[j], TP1.A[j], XOR(TP0.A[j], TP1.A[j]), "B", TP0.B[j], TP1.B[j], XOR(TP0.B[j], TP1.B[j]))
			assert.EqualValues(t, XOR(TP0.C[j], TP1.C[j]), XOR(TP0.A[j], (TP1.A[j])).And(XOR(TP0.B[j], TP1.B[j])), j)
			if j%Sample == 0 {
				fmt.Println(XOR(TP0.C[j], TP1.C[j]), XOR(TP0.A[j], (TP1.A[j])), XOR(TP0.B[j], TP1.B[j]))
			}
		}
	}
}
