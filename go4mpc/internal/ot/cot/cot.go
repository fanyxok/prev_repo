package cot

import (
	"math/rand"
	config "s3l/mpcfgo/config"
	"s3l/mpcfgo/internal/encrypt/prf"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	OT "s3l/mpcfgo/internal/ot"
	"s3l/mpcfgo/internal/ot/baseot"
	"s3l/mpcfgo/pkg/always"
	"time"
)

/*
m data, each has l bits.
*/
func SendN(net network.Network, m int, l int, f func(...interface{}) []byte) ([][]byte, [][]byte) {
	println("Sending", m, "messages. Each has", l, "bits.")
	net.Send(network.NewMsg(0, network.OTe, misc.EncodeInt32(m)))
	net.Send(network.NewMsg(0, network.OTe, misc.EncodeInt32(l)))
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
			gk := OT.RandN(rand_, m)
			u := net.Recv()
			colQ[i] = make([]byte, (m+7)/8)
			misc.BytesXorBytes(colQ[i], u.Data, gk)
		} else {
			net.Recv()
			colQ[i] = OT.RandN(rand_, m)
		}
	}

	S := misc.BoolsToBytes(s)
	x0, x1 := make([][]byte, m), make([][]byte, m)
	for i := 0; i < m; i++ {
		rowQ := OT.Transpose(colQ, i)
		x0[i] = OT.Remapping(prf.FixedKeyAES.Hash(rowQ), l)
		x1[i] = f(x0[i], i)
		y := make([]byte, (l+7)/8)
		QS := make([]byte, config.SymByte)
		misc.BytesXorBytes(QS, rowQ, S)
		h1 := OT.Remapping(prf.FixedKeyAES.Hash(QS), l)
		misc.BytesXorBytes(y, x1[i], h1)
		net.Send(network.NewMsg(11, network.OTe, y))
	}
	return x0, x1
}

func RecvN(net network.Network, r []bool) [][]byte {
	m := int(misc.DecodeInt32(net.Recv().Data))
	l := int(misc.DecodeInt32(net.Recv().Data))
	always.Eq(m, len(r))
	println("Receiving", m, "messages. Each has", l, "bits")
	// sample k random seeds, each length K
	seeds0, seeds1 := make([][]byte, config.SymK), make([][]byte, config.SymK)
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < config.SymK; i++ {
		seeds0[i], seeds1[i] = make([]byte, config.SymByte), make([]byte, config.SymByte)
		seed.Read(seeds0[i])
		seed.Read(seeds1[i])
	}

	// send seeds OT-ly
	baseot.SendN(net, seeds0, seeds1)

	//	T is m * K matrix, where colT is a col
	colT := make([][]byte, config.SymK)
	R := misc.BoolsToBytes(r)
	for i := 0; i < config.SymK; i++ {
		rand_0 := rand.New(rand.NewSource(int64(misc.DecodeInt64(seeds0[i][:8]))))
		colT[i] = OT.RandN(rand_0, m)
		rand_1 := rand.New(rand.NewSource(int64(misc.DecodeInt64(seeds1[i][:8]))))
		u := OT.RandN(rand_1, m)
		misc.BytesXorBytes(u, u, colT[i])
		misc.BytesXorBytes(u, u, R)
		net.Send(network.NewMsg(uint32(i), network.OTe, u))
	}

	// reveal the mask of X
	x := make([][]byte, m)
	for i := 0; i < m; i++ {
		rowT := OT.Transpose(colT, i)
		h := OT.Remapping(prf.FixedKeyAES.Hash(rowT), l)
		x[i] = make([]byte, (l+7)/8)
		y := net.Recv()
		if r[i] {
			misc.BytesXorBytes(x[i], y.Data, h)
		} else {
			x[i] = h
		}
	}
	return x
}
