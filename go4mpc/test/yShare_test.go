package test

import (
	"log"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/pvt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYShare_New_NewFrom_Declassify(t *testing.T) {
	serNet, cliNet := Nets()
	N := 100
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
				x := pvt.YShare.NewFrom(serNet)
				x_prime := pvt.YShare.New(serNet, X[i])
				x.Declassify(serNet)
				x_prime.Declassify(serNet)
				assert.EqualValues(t, X[i], x.GetPlaintext())
				assert.EqualValues(t, X[i], x_prime.GetPlaintext())
				if i%Sample == 0 {
					log.Println(x.GetPlaintext(), "/", x_prime.GetPlaintext(), "=", X[i])
				}
			}
		},
		func() {
			for i := 0; i < 5*N; i++ {
				x := pvt.YShare.New(cliNet, X[i])
				x_prime := pvt.YShare.NewFrom(cliNet)
				x.Declassify(cliNet)
				x_prime.Declassify(cliNet)
				assert.EqualValues(t, X[i], x.GetPlaintext())
				assert.EqualValues(t, X[i], x_prime.GetPlaintext())
			}
		})
}

func TestYSharePub_Shr_Shl(t *testing.T) {
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

func TestYSharePub_Add_Sub(t *testing.T) {
	serNet, cliNet := Nets()
	N := 1000
	X := make([]pub.PubNum, 4*N)
	Y := make([]pub.PubNum, 4*N)

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
				x := pvt.YShare.New(serNet, X[i])
				add_z := x.Add(serNet, Y[i])
				add_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Add(Y[i]), add_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(add_z.GetPlaintext(), "/", X[i].Add(Y[i]), "=", X[i], " +", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				add_z := x.Add(cliNet, Y[i])
				add_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Add(Y[i]), add_z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.New(serNet, X[i])
				sub_z := x.Sub(serNet, Y[i])
				sub_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Sub(Y[i]), sub_z.GetPlaintext())
				if i%20 == 0 {
					log.Println(sub_z.GetPlaintext(), "/", X[i].Sub(Y[i]), "=", X[i], " -", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				sub_z := x.Sub(cliNet, Y[i])
				sub_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Sub(Y[i]), sub_z.GetPlaintext())
			}
		})

}
func TestYSharePub_Mul(t *testing.T) {
	serNet, cliNet := Nets()
	N := 200
	X := make([]pub.PubNum, 4*N)
	Y := make([]pub.PubNum, 4*N)

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
				x := pvt.YShare.New(serNet, X[i])
				mul_z := x.Mul(serNet, Y[i])
				mul_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Mul(Y[i]), mul_z.GetPlaintext())
				if i%20 == 0 {
					log.Println(mul_z.GetPlaintext(), "/", X[i].Mul(Y[i]), "=", X[i], " *", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				mul_z := x.Mul(cliNet, Y[i])
				mul_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Mul(Y[i]), mul_z.GetPlaintext())
			}
		})

}
func TestYSharePub_Div(t *testing.T) {
	serNet, cliNet := Nets()
	N := 200
	X := make([]pub.PubNum, 4*N)
	Y := make([]pub.PubNum, 4*N)

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
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.New(serNet, X[i])
				div_z := x.Div(serNet, Y[i])
				div_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Div(Y[i]), div_z.GetPlaintext())
				if i%20 == 0 {
					log.Println(div_z.GetPlaintext(), "/", X[i].Div(Y[i]), "=", X[i], " /", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				div_z := x.Div(cliNet, Y[i])
				div_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Div(Y[i]), div_z.GetPlaintext())
			}
		})
}

func TestYSharePub_And_Or(t *testing.T) {
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
			for i := 0; i < N; i++ {
				x := pvt.YShare.New(serNet, X[i])
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
				x := pvt.YShare.NewFrom(cliNet)
				and_z := x.And(cliNet, Y[i])
				and_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].And(Y[i]), and_z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < N; i++ {
				x := pvt.YShare.New(serNet, X[i])
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
				x := pvt.YShare.NewFrom(cliNet)
				or_z := x.Or(cliNet, Y[i])
				or_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Or(Y[i]), or_z.GetPlaintext())
			}
		})
}

func TestYSharePub_Eq(t *testing.T) {
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
			for i := 0; i < 5*N; i++ {
				x := pvt.YShare.New(serNet, X[i])
				eq_z := x.Eq(serNet, Y[i])
				eq_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Eq(Y[i]), eq_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(eq_z.GetPlaintext(), "/", X[i].Eq(Y[i]), "=", X[i], "==", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 5*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				eq_z := x.Eq(cliNet, Y[i])
				eq_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Eq(Y[i]), eq_z.GetPlaintext())
			}
		})
}
func TestYSharePub_Gt_Lt(t *testing.T) {
	serNet, cliNet := Nets()
	N := 100
	X := make([]pub.PubNum, 4*N)
	Y := make([]pub.PubNum, 4*N)
	for i := 0; i < N; i++ {
		X[i] = pub.ZeroInt8.Rand()
		X[i+N] = pub.ZeroInt16.Rand()
		X[i+2*N] = pub.ZeroInt32.Rand()
		X[i+3*N] = pub.ZeroInt64.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		Y[i+N] = pub.ZeroInt16.Rand()
		Y[i+2*N] = pub.ZeroInt32.Rand()
		Y[i+3*N] = pub.ZeroInt64.Rand()
	}
	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.New(serNet, X[i])
				gt_z := x.Gt(serNet, Y[i])
				gt_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Gt(Y[i]), gt_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(gt_z.GetPlaintext(), "/", X[i].Gt(Y[i]), "=", X[i], ">", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				gt_z := x.Gt(cliNet, Y[i])
				gt_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Gt(Y[i]), gt_z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.New(serNet, X[i])
				lt_z := x.Lt(serNet, Y[i])
				lt_z.Declassify(serNet)
				assert.EqualValues(t, X[i].Lt(Y[i]), lt_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(lt_z.GetPlaintext(), "/", X[i].Lt(Y[i]), "=", X[i], "<", Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 4*N; i++ {
				x := pvt.YShare.NewFrom(cliNet)
				lt_z := x.Lt(cliNet, Y[i])
				lt_z.Declassify(cliNet)
				assert.EqualValues(t, X[i].Lt(Y[i]), lt_z.GetPlaintext())
			}
		})
}
func TestYSharePub_Mux(t *testing.T) {
	serNet, cliNet := Nets()
	N := 2000
	X := make([]pub.PubNum, 5*N)
	Y := make([]pub.PubNum, 5*N)
	B := make([]pub.PubNum, 5*N)
	for i := 0; i < N; i++ {
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
	for i := 0; i < 5*N; i++ {
		B[i] = pub.ZeroBool.Rand()
	}
	Parallel(
		func() {
			for i := 0; i < 5*N; i++ {
				b := pvt.YShare.NewFrom(serNet)
				mux_z := b.Mux(serNet, X[i], Y[i])
				mux_z.Declassify(serNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(mux_z.GetPlaintext(), "/", B[i].Mux(X[i], Y[i]), "= Mux", B[i], X[i], Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 5*N; i++ {
				b := pvt.YShare.New(cliNet, B[i])
				mux_z := b.Mux(cliNet, X[i], Y[i])
				mux_z.Declassify(cliNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < 5*N; i++ {
				b := pvt.YShare.NewFrom(serNet)
				x := pvt.YShare.New(serNet, X[i])
				mux_z := b.Mux(serNet, x, Y[i])
				mux_z.Declassify(serNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(mux_z.GetPlaintext(), "/", B[i].Mux(X[i], Y[i]), "= Mux", B[i], X[i], Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 5*N; i++ {
				b := pvt.YShare.New(cliNet, B[i])
				x := pvt.YShare.NewFrom(cliNet)
				mux_z := b.Mux(cliNet, x, Y[i])
				mux_z.Declassify(cliNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
			}
		})
	Parallel(
		func() {
			for i := 0; i < 5*N; i++ {
				b := pvt.YShare.NewFrom(serNet)
				y := pvt.YShare.New(serNet, Y[i])
				mux_z := b.Mux(serNet, X[i], y)
				mux_z.Declassify(serNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
				if i%Sample == 0 {
					log.Println(mux_z.GetPlaintext(), "/", B[i].Mux(X[i], Y[i]), "= Mux", B[i], X[i], Y[i])
				}
			}
		},
		func() {
			for i := 0; i < 5*N; i++ {
				b := pvt.YShare.New(cliNet, B[i])
				y := pvt.YShare.NewFrom(cliNet)
				mux_z := b.Mux(cliNet, X[i], y)
				mux_z.Declassify(cliNet)
				assert.EqualValues(t, B[i].Mux(X[i], Y[i]), mux_z.GetPlaintext())
			}
		})
}
