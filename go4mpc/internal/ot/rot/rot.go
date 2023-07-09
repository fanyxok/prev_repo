package rot

import (
	"math/rand"
	"s3l/mpcfgo/config"
	"s3l/mpcfgo/internal/encrypt/prf"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	OT "s3l/mpcfgo/internal/ot"
	"s3l/mpcfgo/internal/ot/baseot"
	"time"
)

/*
x0 has m L bit string, occupy m * L/8 bytes
x1 has m L bit string, occupy m * L/8 bytes
*/
func SendN(net network.Network, m, l int) ([][]byte, [][]byte) {
	println("RandomOT: Sending", m, "messages. Each has", l, "bits.")
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
	x0 := make([][]byte, m)
	x1 := make([][]byte, m)
	lbyte := (l + 7) / 8
	var rand_ func([]byte)
	if l == 1 {
		rand_ = func(b []byte) {
			if misc.Bool() {
				b[0] = 1
			} else {
				b[0] = 0
			}
		}
	} else {
		rand_ = func(b []byte) {
			rand.Read(b)
		}
	}

	for j := 0; j < m; j++ {
		rowQ := OT.Transpose(colQ, j)
		h0 := OT.Remapping(prf.FixedKeyAES.Hash(rowQ), l)
		y0 := make([]byte, lbyte)
		x0[j] = make([]byte, lbyte)
		rand_(x0[j])
		misc.BytesXorBytes(y0, x0[j], h0)

		QS := make([]byte, config.SymByte)
		misc.BytesXorBytes(QS, rowQ, S)

		h1 := OT.Remapping(prf.FixedKeyAES.Hash(QS), l)
		y1 := make([]byte, lbyte)
		x1[j] = make([]byte, lbyte)
		rand_(x1[j])
		misc.BytesXorBytes(y1, x1[j], h1)
		net.Send(network.NewMsg(11, network.OTe, y0))
		net.Send(network.NewMsg(11, network.OTe, y1))
	}
	return x0, x1
}

func RecvN(net network.Network, r []bool) [][]byte {
	m := len(r)
	tmp := net.Recv()
	l := int(misc.DecodeInt32(tmp.Data))
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
		if r[i] {
			_ = net.Recv()
			t := net.Recv()
			misc.BytesXorBytes(x[i], t.Data, h)
		} else {
			t := net.Recv()
			_ = net.Recv()
			misc.BytesXorBytes(x[i], t.Data, h)
		}
	}
	return x
}
