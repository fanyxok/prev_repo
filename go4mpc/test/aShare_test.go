package test

import (
	"log"
	"s3l/mpcfgo/pkg/primitive/triple"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/pvt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAShare_New_NewFrom_Declassify(t *testing.T) {
	serNet, cliNet := Nets()
	N := 256
	X := make([]pub.PubNum, 5*N)
	for i := 0; i < N; i++ {
		X[i] = pub.ZeroBool.Rand()
		X[i+N] = pub.ZeroInt8.Rand()
		X[i+2*N] = pub.ZeroInt16.Rand()
		X[i+3*N] = pub.ZeroInt32.Rand()
		X[i+4*N] = pub.ZeroInt64.Rand()
	}
	Parallel(
		func() {

			for i := 0; i < 5*N; i++ {
				_ = pvt.AShare.NewFrom(serNet)
				// x_prime := pvt.AShare.New(serNet, X[i])
				// x.Declassify(serNet)
				// x_prime.Declassify(serNet)
				// assert.EqualValues(t, X[i], x.GetPlaintext())
				// assert.EqualValues(t, X[i], x_prime.GetPlaintext())
				// if i%Sample == 0 {
				// 	log.Println(x.GetPlaintext(), "/", x_prime.GetPlaintext(), "=", X[i])
				// }
			}

		},
		func() {

			for i := 0; i < 5*N; i++ {
				_ = pvt.AShare.New(cliNet, X[i])
				// x_prime := pvt.AShare.NewFrom(cliNet)
				// x.Declassify(cliNet)
				// x_prime.Declassify(cliNet)
				// assert.EqualValues(t, X[i], x.GetPlaintext())
				// assert.EqualValues(t, X[i], x_prime.GetPlaintext())
			}
		})
}
func TestAShare_NewN_NewFromN(t *testing.T) {
	serNet, cliNet := Nets()
	N := 256
	X := make([]pub.PubNum, 5*N)
	for i := 0; i < N; i++ {
		X[i] = pub.ZeroBool.Rand()
		X[i+N] = pub.ZeroInt8.Rand()
		X[i+2*N] = pub.ZeroInt16.Rand()
		X[i+3*N] = pub.ZeroInt32.Rand()
		X[i+4*N] = pub.ZeroInt64.Rand()
	}
	Parallel(
		func() {
			for i := 0; i < 5; i++ {
				pvt.AShare.NewFromN(serNet)
				// for j, v := range x {
				// 	v.Declassify(serNet)
				// 	assert.EqualValues(t, X[i*N+j], v.GetPlaintext())
				// }
			}
		},
		func() {
			for i := 0; i < 5; i++ {
				pvt.AShare.NewN(cliNet, X[i*N:(i*N+N)])
				// for _, v := range x {
				// 	v.Declassify(cliNet)
				// 	//assert.EqualValues(t, X[i*N+j], v.GetPlaintext())
				// }
			}
		})
}
func TestAShare_Add_Sub(t *testing.T) {
	serNet, cliNet := Nets()
	N := 10000
	X := make([]pub.PubNum, N*4)
	Y := make([]pub.PubNum, N*4)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		Y[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		Y[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
		Y[i+3*N] = pub.ZeroInt64.Rand()
	}

	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.NewFrom(serNet)
				z := x.Add(serNet, y)
				z.Declassify(serNet)
				assert.EqualValues(t, X[i].Add(Y[i]), z.GetPlaintext())
				if N%20 == 0 {
					log.Println(z.GetPlaintext(), X[i].Add(Y[i]), "=", X[i], "+", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.New(cliNet, Y[i])
				z := x.Add(cliNet, y)
				z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Add(Y[i]), z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.NewFrom(serNet)
				sub_z := x.Sub(serNet, y)
				sub_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Sub(Y[i]), sub_z.GetPlaintext())
				if N%20 == 0 {
					log.Println(sub_z.GetPlaintext(), X[i].Sub(Y[i]), "=", X[i], "-", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.New(cliNet, Y[i])
				sub_z := x.Sub(cliNet, y)
				sub_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Sub(Y[i]), sub_z.GetPlaintext())
			}
		})
}
func TestASharePub_Add_Sub(t *testing.T) {
	serNet, cliNet := Nets()
	N := 10000
	X := make([]pub.PubNum, N*4)
	Y := make([]pub.PubNum, N*4)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		Y[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		Y[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
		Y[i+3*N] = pub.ZeroInt64.Rand()
	}

	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				z := x.Add(serNet, Y[i])
				z.Declassify(serNet)
				assert.EqualValues(t, X[i].Add(Y[i]), z.GetPlaintext())
				if N%20 == 0 {
					log.Println(z.GetPlaintext(), X[i].Add(Y[i]), "=", X[i], "+", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				z := x.Add(cliNet, Y[i])
				z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Add(Y[i]), z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				sub_z := x.Sub(serNet, Y[i])
				sub_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Sub(Y[i]), sub_z.GetPlaintext())
				if N%20 == 0 {
					log.Println(sub_z.GetPlaintext(), X[i].Sub(Y[i]), "=", X[i], "-", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				sub_z := x.Sub(cliNet, Y[i])
				sub_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Sub(Y[i]), sub_z.GetPlaintext())
			}
		})
}
func TestAShare_Mul(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, N*4)
	Y := make([]pub.PubNum, N*4)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		Y[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		Y[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
		Y[i+3*N] = pub.ZeroInt64.Rand()
	}

	Parallel(
		func() {
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt8, N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt16, N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt32, N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt64, N))
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.NewFrom(serNet)
				mul_z := x.Mul(serNet, y)
				mul_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Mul(Y[i]), mul_z.GetPlaintext())
				if N%20 == 0 {
					log.Println(mul_z.GetPlaintext(), X[i].Mul(Y[i]), "=", X[i], "*", Y[i])
				}
			}
		},
		func() {
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt8, N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt16, N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt32, N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt64, N))
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.New(cliNet, Y[i])
				mul_z := x.Mul(cliNet, y)
				mul_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Mul(Y[i]), mul_z.GetPlaintext())
			}
		})
}
func TestASharePub_Mul(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, N*4)
	Y := make([]pub.PubNum, N*4)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		Y[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		Y[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
		Y[i+3*N] = pub.ZeroInt64.Rand()
	}

	Parallel(
		func() {

			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				mul_z := x.Mul(serNet, Y[i])
				mul_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Mul(Y[i]), mul_z.GetPlaintext())
				if N%20 == 0 {
					log.Println(mul_z.GetPlaintext(), X[i].Mul(Y[i]), "=", X[i], "*", Y[i])
				}
			}
		},
		func() {

			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				mul_z := x.Mul(cliNet, Y[i])
				mul_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Mul(Y[i]), mul_z.GetPlaintext())
			}
		})
}
func TestAShare_Div(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, N*4)
	Y := make([]pub.PubNum, N*4)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		Y[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		Y[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
		Y[i+3*N] = pub.ZeroInt64.Rand()
	}
	for i := range Y {
		for Y[i] == Y[i].From(0) {
			Y[i] = Y[i].Rand()
		}
	}

	Parallel(
		func() {
			// pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt8, N))
			// pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt16, N))
			// pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt32, N))
			// pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt64, N))
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.NewFrom(serNet)
				div_z := x.Div(serNet, y)
				div_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Div(Y[i]), div_z.GetPlaintext())
				if N%20 == 0 {
					log.Println(div_z.GetPlaintext(), X[i].Div(Y[i]), "=", X[i], "/", Y[i])
				}
			}
		},
		func() {
			// pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt8, N))
			// pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt16, N))
			// pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt32, N))
			// pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt64, N))
			for i := 0; i < 4*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.New(cliNet, Y[i])
				div_z := x.Div(cliNet, y)
				div_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Div(Y[i]), div_z.GetPlaintext())
			}
		})
}

func TestAShare_Not(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, N)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroBool.Rand()
	}

	Parallel(
		func() {
			for i := 0; i < N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				not_z := x.Not(serNet)
				not_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Not(), not_z.GetPlaintext())
				if N%20 == 0 {
					log.Println(not_z.GetPlaintext(), X[i].Not(), "= Not", X[i])
				}
			}
		},
		func() {
			for i := 0; i < N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				not_z := x.Not(cliNet)
				not_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Not(), not_z.GetPlaintext())
			}
		})
}
func TestASharePub_Shl(t *testing.T) {
	serNet, cliNet := Nets()
	N := 10000
	X := make([]pub.PubNum, 4*N)
	Y := make([]pub.PubNum, 4*N)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
	}
	for i := 0; i < 4*N; i++ {
		Y[i] = pub.ZeroInt8.Rand()
		for Y[i].(pub.Int8) < 0 {
			Y[i] = pub.ZeroInt8.Rand()
		}
	}
	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.New(serNet, X[i])
				shl_z := x.Shl(serNet, Y[i])
				shl_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Shl(Y[i]), shl_z.GetPlaintext())
				if i%20 == 0 {
					log.Println(shl_z.GetPlaintext(), "/", X[i].Shl(Y[i]), "=", X[i], " <<", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				shl_z := x.Shl(cliNet, Y[i])
				shl_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Shl(Y[i]), shl_z.GetPlaintext())
			}
		})
}
func TestYSharePub_Shr(t *testing.T) {
	serNet, cliNet := Nets()
	N := 10000
	X := make([]pub.PubNum, 4*N)
	Y := make([]pub.PubNum, 4*N)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
	}
	for i := 0; i < 4*N; i++ {
		Y[i] = pub.ZeroInt8.Rand()
		for Y[i].(pub.Int8) < 0 {
			Y[i] = pub.ZeroInt8.Rand()
		}
	}
	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.New(serNet, X[i])
				shr_z := x.Shr(serNet, Y[i])
				shr_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Shr(Y[i]), shr_z.GetPlaintext())
				if i%20 == 0 {
					log.Println(shr_z.GetPlaintext(), "/", X[i].Shr(Y[i]), "=", X[i], " >>", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				shr_z := x.Shr(cliNet, Y[i])
				shr_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Shr(Y[i]), shr_z.GetPlaintext())
			}
		})
}
func TestAShareAnd(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, N)
	Y := make([]pub.PubNum, N)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroBool.Rand()
		Y[i] = pub.ZeroBool.Rand()
	}
	Parallel(
		func() {
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroBool, N))
			for i := 0; i < N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.New(serNet, Y[i])
				and_z := x.And(serNet, y)
				and_z.Declassify(serNet)
				assert.EqualValues(t, X[i].And(Y[i]), and_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(and_z.GetPlaintext(), "/", X[i].And(Y[i]), "=", X[i], "&&", Y[i])
				}
			}
		},
		func() {
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroBool, N))
			for i := 0; i < N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.NewFrom(cliNet)
				and_z := x.And(cliNet, y)
				and_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].And(Y[i]), and_z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				and_z := x.And(serNet, Y[i])
				and_z.Declassify(serNet)
				assert.EqualValues(t, X[i].And(Y[i]), and_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(and_z.GetPlaintext(), "/", X[i].And(Y[i]), "=", X[i], "&&", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				and_z := x.And(cliNet, Y[i])
				and_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].And(Y[i]), and_z.GetPlaintext())
			}
		})
}

func TestAShareOr(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, N)
	Y := make([]pub.PubNum, N)

	for i := 0; i < N; i++ {
		X[i] = pub.ZeroBool.Rand()
		Y[i] = pub.ZeroBool.Rand()
	}
	Parallel(
		func() {
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroBool, N))
			for i := 0; i < N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.New(serNet, Y[i])
				or_z := x.Or(serNet, y)
				or_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Or(Y[i]), or_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(or_z.GetPlaintext(), "/", X[i].Or(Y[i]), "=", X[i], "||", Y[i])
				}
			}
		},
		func() {
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroBool, N))

			for i := 0; i < N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.NewFrom(cliNet)
				or_z := x.Or(cliNet, y)
				or_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Or(Y[i]), or_z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				or_z := x.Or(serNet, Y[i])
				or_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Or(Y[i]), or_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(or_z.GetPlaintext(), "/", X[i].Or(Y[i]), "=", X[i], "||", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				or_z := x.Or(cliNet, Y[i])
				or_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Or(Y[i]), or_z.GetPlaintext())
			}
		})
}
func TestAShare_Eq(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, 5*N)
	Y := make([]pub.PubNum, 5*N)

	for i := 0; i < N/2; i++ {
		X[i] = pub.ZeroBool.Rand()
		X[i+N] = pub.ZeroInt8.Rand()
		X[i+2*N] = pub.ZeroInt16.Rand()
		X[i+3*N] = pub.ZeroInt32.Rand()
		X[i+4*N] = pub.ZeroInt64.Rand()
		Y[i] = pub.ZeroBool.Rand()
		Y[i+N] = pub.ZeroInt8.Rand()
		Y[i+2*N] = pub.ZeroInt16.Rand()
		Y[i+3*N] = pub.ZeroInt32.Rand()
		Y[i+4*N] = pub.ZeroInt64.Rand()
	}
	for i := N / 2; i < N; i++ {
		X[i] = pub.ZeroBool.Rand()
		X[i+N] = pub.ZeroInt8.Rand()
		X[i+2*N] = pub.ZeroInt16.Rand()
		X[i+3*N] = pub.ZeroInt32.Rand()
		X[i+4*N] = pub.ZeroInt64.Rand()
		Y[i] = X[i]
		Y[i+N] = X[i+N]
		Y[i+2*N] = X[i+2*N]
		Y[i+3*N] = X[i+3*N]
		Y[i+4*N] = X[i+4*N]
	}
	Parallel(
		func() {
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroBool, 116*N))

			for i := 0; i < 5*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.New(serNet, Y[i])
				eq_z := x.Eq(serNet, y)
				eq_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Eq(Y[i]), eq_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(eq_z.GetPlaintext(), "/", X[i].Eq(Y[i]), "=", X[i], "==", Y[i])
				}
			}
		},
		func() {
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroBool, 116*N))

			for i := 0; i < 5*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.NewFrom(cliNet)
				eq_z := x.Eq(cliNet, y)
				eq_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Eq(Y[i]), eq_z.GetPlaintext())
			}
		})
}
func TestASharePub_Eq(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	Rd := 5
	X := make([]pub.PubNum, 5*N)
	Y := make([]pub.PubNum, 5*N)
	Type := []pub.PubNum{pub.ZeroBool, pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
	for j := range make([]int, Rd) {
		for i := 0; i < N/2; i++ {
			X[i+j*N] = Type[j].Rand()
			Y[i+j*N] = Type[j].Rand()
		}
		for i := N / 2; i < N; i++ {
			X[i+j*N] = Type[j].Rand()
			Y[i+j*N] = X[i+j*N]
		}
	}
	Parallel(
		func() {
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroBool, 116*N))
			for i := 0; i < Rd*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.New(serNet, Y[i])
				eq_z := x.Eq(serNet, y)
				eq_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Eq(Y[i]), eq_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(eq_z.GetPlaintext(), "/", X[i].Eq(Y[i]), "=", X[i], "==", Y[i])
				}
			}
		},
		func() {
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroBool, 116*N))
			for i := 0; i < Rd*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.NewFrom(cliNet)
				eq_z := x.Eq(cliNet, y)
				eq_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Eq(Y[i]), eq_z.GetPlaintext())
			}
		})
}
func TestAShare_Mux(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	Rd := 5
	X := make([]pub.PubNum, Rd*N)
	Y := make([]pub.PubNum, Rd*N)
	B := make([]pub.PubNum, Rd*N)
	Type := []pub.PubNum{pub.ZeroBool, pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
	for j := range make([]int, Rd) {
		for i := 0; i < N; i++ {
			X[i+j*N] = Type[j].Rand()
			Y[i+j*N] = Type[j].Rand()
			B[i+j*N] = pub.ZeroBool.Rand()
		}
	}

	Parallel(
		func() {
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroBool, 2*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt8, 3*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt16, 3*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt32, 3*N))
			pvt.TripleFactory(true).SetTriples(triple.NewTriples(serNet, pub.ZeroInt64, 3*N))

			for i := 0; i < Rd*N; i++ {
				x := pvt.AShare.New(serNet, X[i])
				y := pvt.AShare.New(serNet, Y[i])
				b := pvt.AShare.New(serNet, B[i])
				mux_z := b.Mux(serNet, x, y)
				mux_z.Declassify(serNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(mux_z.GetPlaintext(), "/", B[i].Mux(X[i], Y[i]), "= MUX", B[i], X[i], Y[i])
				}
			}
		},
		func() {
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroBool, 2*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt8, 3*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt16, 3*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt32, 3*N))
			pvt.TripleFactory(false).SetTriples(triple.NewTriples(cliNet, pub.ZeroInt64, 3*N))
			for i := 0; i < Rd*N; i++ {
				x := pvt.AShare.NewFrom(cliNet)
				y := pvt.AShare.NewFrom(cliNet)
				b := pvt.AShare.NewFrom(cliNet)
				mux_z := b.Mux(cliNet, x, y)
				mux_z.Declassify(cliNet)

			}
		})
}
func TestASharePub_Mux(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	Rd := 5
	X := make([]pub.PubNum, Rd*N)
	Y := make([]pub.PubNum, Rd*N)
	B := make([]pub.PubNum, Rd*N)
	Type := []pub.PubNum{pub.ZeroBool, pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
	for j := range make([]int, Rd) {
		for i := 0; i < N; i++ {
			X[i+j*N] = Type[j].Rand()
			Y[i+j*N] = Type[j].Rand()
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
				mux_z := b.Mux(serNet, X[i], Y[i])
				mux_z.Declassify(serNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(mux_z.GetPlaintext(), "/", B[i].Mux(X[i], Y[i]), "= MUX", B[i], X[i], Y[i])
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
				mux_z := b.Mux(cliNet, X[i], Y[i])
				mux_z.Declassify(cliNet)

			}
		})
}
