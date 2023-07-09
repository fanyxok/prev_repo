package pvt

import (
	"s3l/mpcfgo/config"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/internal/ot/baseot"
	"s3l/mpcfgo/pkg/type/pub"
)

func Declassify(net network.Network, p PvtNum) pub.PubNum {
	p.Declassify(net)
	return p.GetPlaintext()
}

func DeclassifyN(net network.Network, p []PvtNum) []pub.PubNum {
	ret := make([]pub.PubNum, len(p))
	for i, v := range p {
		v.Declassify(net)
		ret[i] = v.GetPlaintext()
	}
	return ret
}

func A2YN(net network.Network, x []PvtNum) []PvtNum {
	r := make([]PvtNum, len(x))
	for i, v := range x {
		r[i] = Y2A(net, v)
	}
	return r
}

func Y2AN(net network.Network, x []PvtNum) []PvtNum {
	r := make([]PvtNum, len(x))
	for i, v := range x {
		r[i] = A2Y(net, v)
	}
	return r
}
func A2Y(net network.Network, x PvtNum) PvtNum {
	var sh0, sh1 PvtNum
	if net.Server {
		sh0 = YShare.New(net, x.GetShare())
		sh1 = YShare.NewFrom(net)
	} else {
		sh0 = YShare.NewFrom(net)
		sh1 = YShare.New(net, x.GetShare())
	}
	if x.Length() == 1 {
		return sh0.Eq(net, sh1).Not(net)
	} else {
		return sh0.Add(net, sh1)
	}
}

func Y2A(net network.Network, x_ PvtNum) PvtNum {
	x := x_.(*Yshare)
	if net.Server {
		if x.Length() == 1 {
			r := pub.ZeroBool.Rand().(pub.Bool)
			if x.wTable[0][0][config.SymByte] == 0 {
				baseot.Send(net, r.Xor(pub.Bool(true)).Bytes(), r.Xor(pub.ZeroBool).Bytes())
			} else {
				baseot.Send(net, r.Xor(pub.ZeroBool).Bytes(), r.Xor(pub.Bool(true)).Bytes())
			}
			ash := new(Ashare)
			ash.Pos = 0
			ash.SetShare(r)
			return ash
		} else {
			s0 := make([][]byte, x.Length())
			s1 := make([][]byte, x.Length())
			t := pub.Zero(x.Length())
			for i := 0; i < x.Length(); i++ {
				r := x.GetShare().Rand()
				t = t.Add(r)
				powi := r.From(1 << i)
				if x.wTable[i][0][config.SymByte] == 1 {
					s0[i] = powi.Sub(r).Bytes()
					s1[i] = r.From(0).Sub(r).Bytes()
				} else {
					s0[i] = r.From(0).Sub(r).Bytes()
					s1[i] = powi.Sub(r).Bytes()
				}
			}
			baseot.SendN(net, s0, s1)
			ash := new(Ashare)
			ash.Pos = 0
			ash.SetShare(t)
			return ash
		}
	} else {
		if x.Length() == 1 {
			var t pub.PubNum
			if x.wValue[0][config.SymByte] == 0 {
				t = pub.ZeroBool.Decode(baseot.Recv(net, true))
			} else {
				t = pub.ZeroBool.Decode(baseot.Recv(net, false))
			}
			ash := new(Ashare)
			ash.Pos = 1
			ash.SetShare(t)
			return ash
		} else {
			bitwise := make([]bool, x.Length())
			for i := 0; i < x.Length(); i++ {
				if x.wValue[i][config.SymByte] == 1 {
					bitwise[i] = true
				}
			}
			si := baseot.RecvN(net, bitwise)
			t := pub.Zero(x.Length())
			for i := 0; i < x.Length(); i++ {
				t = t.Add(t.Decode(si[i]))
			}
			ash := new(Ashare)
			ash.Pos = 1
			ash.SetShare(t)
			return ash
		}
	}
}
