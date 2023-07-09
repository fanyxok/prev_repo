package test

import (
	"log"
	"s3l/mpcfgo/pkg/primitive/triple"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/pvt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandBooleanTo(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	Rd := 1
	X := make([]pub.PubNum, Rd*N)
	B := make([]pub.PubNum, Rd*N)
	Type := []pub.PubNum{pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
	for j := range make([]int, Rd) {
		for i := 0; i < N; i++ {
			X[i+j*N] = Type[j].Rand()
			B[i+j*N] = pub.ZeroBool.Rand()
		}
	}
	Parallel(
		func() {
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroBool, 1*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt8, 1*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt16, 1*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt32, 1*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt64, 1*N))

			for i := 0; i < Rd*N; i++ {
				b := pvt.AShare.New(serNet, B[i])
				expand_b := b.(*pvt.Ashare).ExpandBooleanTo(serNet, X[i])
				expand_b.Declassify(serNet)
				if B[i] == pub.ZeroBool {
					assert.EqualValues(t, X[i].From(0), expand_b.GetPlaintext())
				} else {
					assert.EqualValues(t, X[i].From(1), expand_b.GetPlaintext())
				}
				if i%Sample == 0 {
					log.Println(expand_b.GetPlaintext(), "/", B[i])
				}
			}
		},
		func() {
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroBool, 1*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt8, 1*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt16, 1*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt32, 1*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt64, 1*N))
			for i := 0; i < Rd*N; i++ {
				b := pvt.AShare.NewFrom(cliNet)
				expand_b := b.(*pvt.Ashare).ExpandBooleanTo(cliNet, X[i])
				expand_b.Declassify(cliNet)
			}
		})
}
