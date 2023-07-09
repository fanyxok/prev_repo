package test

import (
	"log"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/pvt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCircuit_Add(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 1000
	X := make([]pub.PubNum, n*4)
	Y := make([]pub.PubNum, n*4)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+n] = pub.ZeroInt16.Rand()
		Y[i+n] = pub.ZeroInt16.Rand()
		X[i+2*n] = pub.ZeroInt32.Rand()
		Y[i+2*n] = pub.ZeroInt32.Rand()
		X[i+3*n] = pub.ZeroInt64.Rand()
		Y[i+3*n] = pub.ZeroInt64.Rand()
	}

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Add(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Add(Y[i]), z.GetPlaintext())
			if n%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Add(Y[i]), "=", X[i], "+", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Add(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Add(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Sub(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 1000
	X := make([]pub.PubNum, n*4)
	Y := make([]pub.PubNum, n*4)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+n] = pub.ZeroInt16.Rand()
		Y[i+n] = pub.ZeroInt16.Rand()
		X[i+2*n] = pub.ZeroInt32.Rand()
		Y[i+2*n] = pub.ZeroInt32.Rand()
		X[i+3*n] = pub.ZeroInt64.Rand()
		Y[i+3*n] = pub.ZeroInt64.Rand()
	}

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Sub(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Sub(Y[i]), z.GetPlaintext())
			if n%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Sub(Y[i]), "=", X[i], "-", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Sub(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Sub(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Mul(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	N := 100
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

	go func() {
		for i := 0; i < 4*N; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Mul(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Mul(Y[i]), z.GetPlaintext(), i)
			if i%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Mul(Y[i]), "=", X[i], "*", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 4*N; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Mul(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Mul(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Div(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 100
	X := make([]pub.PubNum, n*4)
	Y := make([]pub.PubNum, n*4)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+n] = pub.ZeroInt16.Rand()
		Y[i+n] = pub.ZeroInt16.Rand()
		X[i+2*n] = pub.ZeroInt32.Rand()
		Y[i+2*n] = pub.ZeroInt32.Rand()
		X[i+3*n] = pub.ZeroInt64.Rand()
		Y[i+3*n] = pub.ZeroInt64.Rand()
	}
	for i := range Y {
		for Y[i] == Y[i].From(0) {
			Y[i] = Y[i].Rand()
		}
	}

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Div(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Div(Y[i]), z.GetPlaintext(), i)
			if n%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Div(Y[i]), "=", X[i], "/", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Div(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Div(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Not(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 100
	X := make([]pub.PubNum, n)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroBool.Rand()
	}

	go func() {
		for i := 0; i < n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			z := x.Not(serNet)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Not(), z.GetPlaintext(), i)
			if n%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Not(), "= !", X[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			z := x.Not(cliNet)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Not(), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_And(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 100
	X := make([]pub.PubNum, n)
	Y := make([]pub.PubNum, n)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroBool.Rand()
		Y[i] = pub.ZeroBool.Rand()
	}

	go func() {
		for i := 0; i < n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.And(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].And(Y[i]), z.GetPlaintext(), i)
			if n%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].And(Y[i]), "=", X[i], "&&", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.And(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].And(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Or(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 100
	X := make([]pub.PubNum, n)
	Y := make([]pub.PubNum, n)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroBool.Rand()
		Y[i] = pub.ZeroBool.Rand()
	}

	go func() {
		for i := 0; i < n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Or(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Or(Y[i]), z.GetPlaintext(), i)
			if n%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Or(Y[i]), "=", X[i], "||", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Or(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Or(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Eq(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 1000
	X := make([]pub.PubNum, n*5)
	Y := make([]pub.PubNum, n*5)

	for i := 0; i < n/2; i++ {
		X[i] = pub.ZeroBool.Rand()
		Y[i] = pub.ZeroBool.Rand()
		X[i+n] = pub.ZeroInt8.Rand()
		Y[i+n] = pub.ZeroInt8.Rand()
		X[i+2*n] = pub.ZeroInt16.Rand()
		Y[i+2*n] = pub.ZeroInt16.Rand()
		X[i+3*n] = pub.ZeroInt32.Rand()
		Y[i+3*n] = pub.ZeroInt32.Rand()
		X[i+4*n] = pub.ZeroInt64.Rand()
		Y[i+4*n] = pub.ZeroInt64.Rand()
	}
	for i := n / 2; i < n; i++ {
		X[i] = pub.ZeroBool.Rand()
		Y[i] = X[i]
		X[i+n] = pub.ZeroInt8.Rand()
		Y[i+n] = X[i+n]
		X[i+2*n] = pub.ZeroInt16.Rand()
		Y[i+2*n] = X[i+2*n]
		X[i+3*n] = pub.ZeroInt32.Rand()
		Y[i+3*n] = X[i+3*n]
		X[i+4*n] = pub.ZeroInt64.Rand()
		Y[i+4*n] = X[i+4*n]
	}

	go func() {
		for i := 0; i < 5*n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Eq(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Eq(Y[i]), z.GetPlaintext(), i)
			if i%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Eq(Y[i]), "=", X[i], "==", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 5*n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Eq(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Eq(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Gt(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 1000
	X := make([]pub.PubNum, n*4)
	Y := make([]pub.PubNum, n*4)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+n] = pub.ZeroInt16.Rand()
		Y[i+n] = pub.ZeroInt16.Rand()
		X[i+2*n] = pub.ZeroInt32.Rand()
		Y[i+2*n] = pub.ZeroInt32.Rand()
		X[i+3*n] = pub.ZeroInt64.Rand()
		Y[i+3*n] = pub.ZeroInt64.Rand()
	}

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Gt(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Gt(Y[i]), z.GetPlaintext(), i)
			if i%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Gt(Y[i]), "=", X[i], ">", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Gt(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Gt(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Lt(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	n := 1000
	X := make([]pub.PubNum, n*4)
	Y := make([]pub.PubNum, n*4)

	for i := 0; i < n; i++ {
		X[i] = pub.ZeroInt8.Rand()
		Y[i] = pub.ZeroInt8.Rand()
		X[i+n] = pub.ZeroInt16.Rand()
		Y[i+n] = pub.ZeroInt16.Rand()
		X[i+2*n] = pub.ZeroInt32.Rand()
		Y[i+2*n] = pub.ZeroInt32.Rand()
		X[i+3*n] = pub.ZeroInt64.Rand()
		Y[i+3*n] = pub.ZeroInt64.Rand()
	}

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := x.Lt(serNet, y)
			z.Declassify(serNet)
			assert.EqualValues(t, X[i].Lt(Y[i]), z.GetPlaintext(), i)
			if i%20 == 0 {
				log.Println(z.GetPlaintext(), X[i].Lt(Y[i]), "=", X[i], ">", Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 4*n; i++ {
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := x.Lt(cliNet, y)
			z.Declassify(cliNet)
			assert.EqualValues(t, X[i].Lt(Y[i]), z.GetPlaintext())
		}
		ch <- true
	}()
	<-ch
	<-ch
}

func TestCircuit_Mux(t *testing.T) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	cliNet := network.NewClient(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	N := 1000
	B := make([]pub.PubNum, 5*N)
	X := make([]pub.PubNum, N*5)
	Y := make([]pub.PubNum, N*5)
	for i := 0; i < 5*N; i++ {
		B[i] = pub.ZeroBool.Rand()
	}
	for i := 0; i < N; i++ {
		X[i] = pub.ZeroBool.Rand()
		Y[i] = pub.ZeroBool.Rand()
		X[i+N] = pub.ZeroInt8.Rand()
		Y[i+N] = pub.ZeroInt8.Rand()
		X[i+2*N] = pub.ZeroInt16.Rand()
		Y[i+2*N] = pub.ZeroInt16.Rand()
		X[i+3*N] = pub.ZeroInt32.Rand()
		Y[i+3*N] = pub.ZeroInt32.Rand()
		X[i+4*N] = pub.ZeroInt64.Rand()
		Y[i+4*N] = pub.ZeroInt64.Rand()
	}

	go func() {
		for i := 0; i < 5*N; i++ {
			b := pvt.YShare.New(serNet, B[i])
			x := pvt.YShare.New(serNet, X[i])
			y := pvt.YShare.NewFrom(serNet)
			z := b.Mux(serNet, x, y)
			expect := B[i].Mux(X[i], Y[i])
			z.Declassify(serNet)
			actual := z.GetPlaintext()
			assert.EqualValues(t, actual, expect, i)
			if i%20 == 0 {
				log.Println(actual, expect, "= Mux", B[i], X[i], Y[i])
			}
		}
		ch <- true
	}()

	go func() {
		for i := 0; i < 5*N; i++ {
			b := pvt.YShare.NewFrom(cliNet)
			x := pvt.YShare.NewFrom(cliNet)
			y := pvt.YShare.New(cliNet, Y[i])
			z := b.Mux(cliNet, x, y)
			expect := B[i].Mux(X[i], Y[i])
			z.Declassify(cliNet)
			actual := z.GetPlaintext()
			assert.EqualValues(t, actual, expect, i)
		}
		ch <- true
	}()
	<-ch
	<-ch
}
