package triple

import (
	"log"
	"math/rand"
	"s3l/mpcfgo/config"
	"s3l/mpcfgo/internal/encrypt/prf"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	OT "s3l/mpcfgo/internal/ot"
	"s3l/mpcfgo/internal/ot/baseot"
	"s3l/mpcfgo/pkg/type/pub"
	"sync/atomic"
	"time"
)

// Two party Triple
// C = A * B
type Triples struct {
	A    []pub.PubNum
	B    []pub.PubNum
	C    []pub.PubNum
	Zero pub.PubNum
	idx  atomic.Int32
}

// width : the bits of number
// cap: the capacity of this factory
func NewTriples(net network.Network, val pub.PubNum, cap int) *Triples {
	if cap == 0 {
		return nil
	}
	if val.Length() == 1 {
		tf := &Triples{
			A:    make([]pub.PubNum, cap),
			B:    make([]pub.PubNum, cap),
			C:    make([]pub.PubNum, cap),
			Zero: pub.ZeroBool,
			idx:  atomic.Int32{},
		}
		tf.BooleanTripleN(net, cap)
		return tf
	}
	tf := &Triples{Zero: pub.Zero(val.Length())}
	tf.PreprocessTripleN(net, tf.Zero, cap)
	return tf
}

// PreprocessTripleN 's second parameter could be Zero
func (ct *Triples) PreprocessTripleN(net network.Network, p pub.PubNum, n int) {
	ct.A = make([]pub.PubNum, n)
	ct.B = make([]pub.PubNum, n)
	if net.Server {
		for i := 0; i < n; i++ {
			p = p.Rand()
			r := p.Rand()
			ct.A[i] = p.Sub(r)
			net.Send(network.NewMsg(00, network.Sharing, r.Bytes()))
			ct.B[i] = pub.DecodePubNum(net.Recv().Data)
		}
	} else {
		for i := 0; i < n; i++ {
			ct.A[i] = pub.DecodePubNum(net.Recv().Data)
			p = p.Rand()
			r := p.Rand()
			ct.B[i] = p.Sub(r)
			net.Send(network.NewMsg(00, network.Sharing, r.Bytes()))
		}
	}
	ct.C = TripleN(net, ct.A, ct.B)
}

func TripleN(net network.Network, a, b []pub.PubNum) []pub.PubNum {
	c := make([]pub.PubNum, len(a))

	if net.Server {
		tripleN0(net, a, b, c)
	} else {
		tripleN1(net, a, b, c)
	}
	return c
}

func tripleN0(net network.Network, a, b, c []pub.PubNum) {
	num := len(a)
	fbuilder := func(T pub.PubNum) func([]byte, int) []byte {
		f := func(b []byte, j int) []byte {
			s := T
			x := s.Decode(b)
			j2 := s.From(1 << j)
			y := x.Add(T.Mul(j2))
			return y.Bytes()
		}
		return f
	}
	m := make([]int, num)
	fa0b1 := make([]func([]byte, int) []byte, num)
	fa1b0 := make([]func([]byte, int) []byte, num)
	a0b1 := make([]pub.PubNum, num)
	a1b0 := make([]pub.PubNum, num)

	for i, v := range a {
		m[i] = v.Length()
		fa0b1[i] = fbuilder(v)
		fa1b0[i] = fbuilder(b[i])
		a0b1[i] = v.From(0)
		a1b0[i] = v.From(0)
	}

	xq, _ := COTSendN(net, m, fa0b1)
	{
		// calculate (a0b1)0
		i := 0
		for j, v := range m {
			for k := 0; k < v; k++ {
				r := a0b1[j].Decode(xq[i])
				a0b1[j] = a0b1[j].Sub(r)
				i++
			}
		}
	}

	xw, _ := COTSendN(net, m, fa1b0)
	{
		// calculate (a1b0)0
		i := 0
		for j, v := range m {
			for k := 0; k < v; k++ {
				r := a1b0[j].Decode(xw[i])
				a1b0[j] = a1b0[j].Sub(r)
				i++
			}
		}
	}
	// t = a0b0 + (a0b1)0 + (a1b0)0
	for i := 0; i < len(c); i++ {
		c[i] = a[i].Mul(b[i]).Add(a0b1[i]).Add(a1b0[i])
	}
}

func tripleN1(net network.Network, a, b, c []pub.PubNum) {
	num := len(a)
	m := make([]int, num)
	rb := make([][]bool, num)
	ra := make([][]bool, num)
	a0b1 := make([]pub.PubNum, num)
	a1b0 := make([]pub.PubNum, num)
	for i, v := range a {
		m[i] = v.Length()
		a0b1[i] = v.From(0)
		a1b0[i] = v.From(0)
		ra[i] = misc.BytesToBools(v.Bytes())
		rb[i] = misc.BytesToBools(b[i].Bytes())
	}

	xq := COTRecvN(net, rb)
	{
		// calculate (a0b1)0
		i := 0
		for j, v := range m {
			for k := 0; k < v; k++ {
				r := a0b1[j].Decode(xq[i])
				a0b1[j] = a0b1[j].Add(r)
				i++
			}
		}
	}

	xw := COTRecvN(net, ra)
	{
		// calculate (a1b0)0
		i := 0
		for j, v := range m {
			for k := 0; k < v; k++ {
				r := a1b0[j].Decode(xw[i])
				a1b0[j] = a1b0[j].Add(r)
				i++
			}
		}
	}

	for i := 0; i < num; i++ {
		c[i] = a[i].Mul(b[i]).Add(a0b1[i]).Add(a1b0[i])
	}
}
func (ct *Triples) Next() (pub.PubNum, pub.PubNum, pub.PubNum) {
	l := ct.idx.Load()
	ct.idx.Add(1)
	return ct.A[l], ct.B[l], ct.C[l]
}

/*
len(m) batchs of data, each batch has m[i] data. each data in i-th batch has m[i] bits
each batch has a correlation function f.
*/
func COTSendN(net network.Network, m []int, f []func([]byte, int) []byte) ([][]byte, [][]byte) {
	log.Println("SendCOT for Triple, with batches", len(m), ". First batch size ", m[0])
	totalM := 0
	for _, v := range m {
		totalM += v
	}

	// sample K randmon bits.
	s := make([]bool, config.SymK)
	for i := 0; i < config.SymK; i++ {
		s[i] = misc.Bool()
	}
	// recv seeds. [K x K] bits
	seeds := baseot.RecvN(net, s)

	colQ := make([][]byte, config.SymK)
	for i := 0; i < config.SymK; i++ {
		rand_ := rand.New(rand.NewSource(int64(misc.DecodeInt64(seeds[i][:8]))))
		if s[i] {
			gk := OT.RandN(rand_, totalM)
			u := net.Recv()
			colQ[i] = make([]byte, (totalM+7)/8)
			misc.BytesXorBytes(colQ[i], u.Data, gk)
		} else {
			net.Recv()
			colQ[i] = OT.RandN(rand_, totalM)
		}
	}

	S := misc.BoolsToBytes(s)
	x0, x1 := make([][]byte, totalM), make([][]byte, totalM)
	i := 0
	for j, v := range m {
		for k := 0; k < v; k++ {
			rowQ := OT.Transpose(colQ, i)
			x0[i] = OT.Remapping(prf.FixedKeyAES.Hash(rowQ), v)
			x1[i] = f[j](x0[i], k)
			y := make([]byte, (v+7)/8)
			QS := make([]byte, config.SymByte)
			misc.BytesXorBytes(QS, rowQ, S)
			h1 := OT.Remapping(prf.FixedKeyAES.Hash(QS), v)
			misc.BytesXorBytes(y, x1[i], h1)
			net.Send(network.NewMsg(11, network.OTe, y))
			i++
		}
	}
	return x0, x1
}

func COTRecvN(net network.Network, r [][]bool) [][]byte {
	num := len(r)
	m := make([]int, num)
	totalM := 0
	for j, v := range r {
		m[j] = len(v)
		totalM += m[j]
	}
	log.Println("RecvCOT for Triple generation, with batches", len(m), ". Total bits", totalM)
	// sample k random seeds, each length K
	seeds0, seeds1 := make([][]byte, config.SymK), make([][]byte, config.SymK)
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < config.SymK; i++ {
		seeds0[i], seeds1[i] = make([]byte, config.SymByte), make([]byte, config.SymByte)
		seed.Read(seeds0[i])
		seed.Read(seeds1[i])
	}
	// send seeds by OT
	baseot.SendN(net, seeds0, seeds1)

	//	T is m * K matrix, where colT is a col
	colT := make([][]byte, config.SymK)
	R := make([]byte, totalM/8)
	ll := 0
	for _, v := range r {
		bytv := misc.BoolsToBytes(v)
		copy(R[ll:], bytv)
		ll += len(bytv)
	}
	for i := 0; i < config.SymK; i++ {
		rand_0 := rand.New(rand.NewSource(int64(misc.DecodeInt64(seeds0[i][:8]))))
		rand_1 := rand.New(rand.NewSource(int64(misc.DecodeInt64(seeds1[i][:8]))))
		colT[i] = OT.RandN(rand_0, totalM)
		u := OT.RandN(rand_1, totalM)
		misc.BytesXorBytes(u, u, colT[i])
		misc.BytesXorBytes(u, u, R)
		net.Send(network.NewMsg(uint32(i), network.OTe, u))
	}

	// reveal the mask of X
	x := make([][]byte, totalM)
	i := 0
	for j, v := range m {
		for k := 0; k < v; k++ {
			rowT := OT.Transpose(colT, i)
			h := OT.Remapping(prf.FixedKeyAES.Hash(rowT), v)
			x[i] = make([]byte, (v+7)/8)
			y := net.Recv()
			if r[j][k] {
				misc.BytesXorBytes(x[i], y.Data, h)
			} else {
				x[i] = h
			}
			i++
		}
	}
	return x
}
