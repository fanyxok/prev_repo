package test

import (
	"log"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/pvt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA2Y(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	Rd := 5
	X := make([]pub.PubNum, Rd*N)
	Type := []pub.PubNum{pub.ZeroBool, pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
	for j := range make([]int, Rd) {
		for i := 0; i < N; i++ {
			X[i+j*N] = Type[j].Rand()
		}
	}
	Parallel(
		func() {

			for i := 0; i < Rd*N; i++ {
				x := pvt.AShare.NewFrom(serNet)
				x_prime := pvt.A2Y(serNet, x)
				x_prime.Declassify(serNet)
				assert.EqualValues(t, X[i], x_prime.GetPlaintext())
				if i%Sample == 0 {
					log.Println(x.GetPlaintext(), "/", x_prime.GetPlaintext(), "=", X[i])
				}
			}
		},
		func() {

			for i := 0; i < Rd*N; i++ {
				x := pvt.AShare.New(cliNet, X[i])
				x_prime := pvt.A2Y(cliNet, x)
				x_prime.Declassify(cliNet)
				assert.EqualValues(t, X[i], x_prime.GetPlaintext())
			}
		})
}
func TestY2A(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	Rd := 5
	X := make([]pub.PubNum, Rd*N)
	Type := []pub.PubNum{pub.ZeroBool, pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
	for j := range make([]int, Rd) {
		for i := 0; i < N; i++ {
			X[i+j*N] = Type[j].Rand()
		}
	}
	Parallel(
		func() {

			for i := 0; i < Rd*N; i++ {
				x := pvt.YShare.NewFrom(serNet)
				x_prime := pvt.Y2A(serNet, x)
				x_prime.Declassify(serNet)
				assert.EqualValues(t, X[i], x_prime.GetPlaintext())
				if i%Sample == 0 {
					log.Println(x.GetPlaintext(), "/", x_prime.GetPlaintext(), "=", X[i])
				}
			}
		},
		func() {

			for i := 0; i < Rd*N; i++ {
				x := pvt.YShare.New(cliNet, X[i])
				x_prime := pvt.Y2A(cliNet, x)
				x_prime.Declassify(cliNet)
				assert.EqualValues(t, X[i], x_prime.GetPlaintext())
			}
		})
}
