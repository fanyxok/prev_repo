package pvt

import (
	"log"
	"s3l/mpcfgo/config"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/pkg/primitive/triple"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/value"
)

func emptyAshare(b bool) *Ashare {
	ret := new(Ashare)
	if b {
		ret.Pos = 0
	} else {
		ret.Pos = 1
	}
	return ret
}

var AShare = new(Ashare)
var tripleFactory0 *triple.TriplesFactory
var tripleFactory1 *triple.TriplesFactory
var TripleFactory = func(isServer bool) *triple.TriplesFactory {
	if config.Debug {
		if isServer {
			return tripleFactory0
		} else {
			return tripleFactory1
		}
	} else {
		return tripleFactory0
	}
}

func init() {
	tripleFactory0 = triple.NewTripleFactory()
	tripleFactory1 = triple.NewTripleFactory()
}

type Ashare struct {
	Plaintext pub.PubNum
	shares    [2]pub.PubNum
	Pos       int
}

func (ct *Ashare) GetShare() pub.PubNum {
	return ct.shares[ct.Pos]
}

func (ct *Ashare) GetPlaintext() pub.PubNum {
	return ct.Plaintext
}

func (ct *Ashare) New(net network.Network, num pub.PubNum) PvtNum {
	s := new(Ashare)
	s.Plaintext = num
	if net.Server {
		s.Pos = 0
	} else {
		s.Pos = 1
	}
	r := num.Rand()
	if num.Length() == 1 {
		s.shares[s.Pos] = pub.Bool(bool(r.(pub.Bool)) != bool(num.(pub.Bool)))
	} else {
		s.shares[s.Pos] = num.Sub(r)
	}
	net.Send(network.NewMsg(00, network.Sharing, []byte{byte(num.Length())}))
	net.Send(network.NewMsg(00, network.Sharing, r.Bytes()))
	return s
}
func (ct *Ashare) NewN(net network.Network, num []pub.PubNum) []PvtNum {
	length := num[0].Length()
	singleBytes := (length + 7) / 8
	totalBytes := make([]byte, singleBytes*len(num))
	rand_.Read(totalBytes)
	s := make([]PvtNum, len(num))
	p := 1
	if net.Server {
		p = 0
	}
	for i, v := range num {
		sv := new(Ashare)
		sv.Plaintext = v
		sv.Pos = p
		if length == 1 {
			totalBytes[i] &= 1
			sv.shares[p] = pub.Bool(bool(totalBytes[i] != v.Bytes()[0]))
		} else {
			sv.shares[p] = v.Sub(v.Decode(totalBytes[singleBytes*i : singleBytes*(i+1)]))
		}
		s[i] = sv
	}
	net.Send(network.NewMsg(00, network.Sharing, []byte{byte(length)}))
	net.Send(network.NewMsg(00, network.Sharing, misc.EncodeInt32(len(num))))
	net.Send(network.NewMsg(00, network.Sharing, totalBytes))
	return s
}
func (ct *Ashare) NewFromN(net network.Network) []PvtNum {
	length := int(net.Recv().Data[0])
	n := misc.DecodeInt32(net.Recv().Data)
	s := make([]PvtNum, n)
	singleBytes := (length + 7) / 8
	totalBytes := net.Recv().Data
	p := 1
	if net.Server {
		p = 0
	}
	for i := 0; i < int(n); i++ {
		v := new(Ashare)
		v.Pos = p
		if length == 1 {
			v.shares[p] = pub.ZeroBool.Decode(totalBytes[i : i+1])
		} else {
			v.shares[p] = pub.DecodePubNum(totalBytes[singleBytes*i : singleBytes*(i+1)])
		}
		s[i] = v
	}
	return s
}
func (ct *Ashare) NewFrom(net network.Network) PvtNum {
	s := new(Ashare)
	if net.Server {
		s.Pos = 0
	} else {
		s.Pos = 1
	}
	length := int(net.Recv().Data[0])
	sh := net.Recv()
	if length == 1 {
		s.shares[s.Pos] = pub.ZeroBool.Decode(sh.Data)
	} else {
		s.shares[s.Pos] = pub.DecodePubNum(sh.Data)
	}
	return s
}

/*
################################################

	PvtNum interface impl

################################################
*/
func (ct *Ashare) Number()     {}
func (ct *Ashare) Length() int { return ct.GetShare().Length() }
func (ct *Ashare) Private()    {}
func (ct *Ashare) Add(net network.Network, x value.Value) PvtNum {
	ret := new(Ashare)
	if net.Server {
		ret.Pos = 0
	} else {
		ret.Pos = 1
	}
	switch v := x.(type) {
	case pub.PubNum:
		if ret.Pos == 0 {
			ret.SetShare(ct.GetShare().Add(v))
		} else {
			ret.SetShare(ct.GetShare())
		}
	case PvtNum:
		sv := v.(*Ashare)
		ret.SetShare(ct.GetShare().Add(sv.GetShare()))
	}
	return ret
}

func (ct *Ashare) Sub(net network.Network, x value.Value) PvtNum {
	ret := new(Ashare)
	if net.Server {
		ret.Pos = 0
	} else {
		ret.Pos = 1
	}
	switch v := x.(type) {
	case pub.PubNum:
		if ret.Pos == 0 {
			ret.SetShare(ct.GetShare().Sub(v))
		} else {
			ret.SetShare(ct.GetShare())
		}
	case PvtNum:
		sv := v.(*Ashare)
		ret.SetShare(ct.GetShare().Sub(sv.GetShare()))
	}
	return ret
}

func (ct *Ashare) Mul(net network.Network, x value.Value) PvtNum {
	ret := new(Ashare)
	if net.Server {
		ret.Pos = 0
	} else {
		ret.Pos = 1
	}
	switch v := x.(type) {
	case pub.PubNum:
		ret.SetShare(ct.GetShare().Mul(v))
	case PvtNum:
		lv := v.(*Ashare)
		a, b, c := TripleFactory(net.Server).NextTriple(lv.Length())
		da := ct.GetShare().Sub(a)
		db := lv.GetShare().Sub(b)
		net.Send(network.NewMsg(0, network.Sharing, da.Bytes()))
		net.Send(network.NewMsg(0, network.Sharing, db.Bytes()))
		da = da.Add(da.Decode(net.Recv().Data))
		db = db.Add(db.Decode(net.Recv().Data))
		if net.Server {
			ret.SetShare(a.Mul(db).Add(b.Mul(da)).Add(c))
		} else {
			ret.SetShare(da.Mul(db).Add(a.Mul(db).Add(b.Mul(da)).Add(c)))
		}
	}
	return ret
}

func (ct *Ashare) Div(net network.Network, x value.Value) PvtNum {
	ret := new(Ashare)
	if net.Server {
		ret.Pos = 0
	} else {
		ret.Pos = 1
	}
	switch v := x.(type) {
	case pub.PubNum:
		if ret.Pos == 0 {
			ret.SetShare(ct.GetShare().Sub(v))
		} else {
			ret.SetShare(ct.GetShare())
		}
	case PvtNum:
		sv := v.(*Ashare)
		ret.SetShare(ct.GetShare().Sub(sv.GetShare()))
	}
	return ret
}

func (ct *Ashare) Not(net network.Network) PvtNum {
	if ct.GetShare().Length() == 1 {
		ret := new(Ashare)
		if net.Server {
			ret.Pos = 0
			ret.SetShare(ct.GetShare().Not())
		} else {
			ret.Pos = 1
			ret.SetShare(ct.GetShare())
		}
		return ret
	} else {
		log.Panicln("Not only support 1 bit Bool")
		return nil
	}
}

func (ct *Ashare) And(net network.Network, x value.Value) PvtNum {
	ret := emptyAshare(net.Server)
	if v, ok := x.(*Ashare); ok {
		a, b, c := TripleFactory(net.Server).NextTriple(v.Length())
		da := ct.GetShare().(pub.Bool).Xor(a)
		db := v.GetShare().(pub.Bool).Xor(b)
		net.Send(network.NewMsg(0, network.Sharing, da.Bytes()))
		net.Send(network.NewMsg(0, network.Sharing, db.Bytes()))
		da = da.Xor(da.Decode(net.Recv().Data))
		db = db.Xor(db.Decode(net.Recv().Data))
		if net.Server {
			ret.SetShare(a.And(db).(pub.Bool).Xor(b.And(da)).Xor(c))
		} else {
			ret.SetShare(da.And(db).(pub.Bool).Xor(a.And(db).(pub.Bool).Xor(b.And(da)).Xor(c)))
		}
	} else if v, ok := x.(pub.Bool); ok {
		ret.SetShare(ct.GetShare().And(v))
	} else {
		log.Panicln("")
	}
	return ret
}
func (ct *Ashare) Or(net network.Network, x value.Value) PvtNum {
	var ret PvtNum
	if v, ok := x.(*Ashare); ok {
		notct := ct.Not(net)
		notx := v.Not(net)
		notret := notct.And(net, notx)
		ret = notret.Not(net)
	} else if v, ok := x.(pub.Bool); ok {
		ret = ct.Not(net).And(net, v.Not()).Not(net)
	} else {
		log.Panicln("")
	}
	return ret
}
func (ct *Ashare) Eq(net network.Network, x value.Value) PvtNum {
	var ret PvtNum
	if x.Length() == 1 {
		if v, ok := x.(PvtNum); ok {
			notct := ct.Not(net)
			r := emptyAshare(net.Server)
			r.SetShare(notct.GetShare().(pub.Bool).Xor(v.GetShare()))
			ret = r
		} else if v, ok := x.(pub.PubNum); ok {
			ret = ct.Not(net).Or(net, v).And(net, ct.Or(net, v.Not()))
		} else {
			log.Panicln("")
		}
	} else {
		if v, ok := x.(PvtNum); ok {
			var diff pub.PubNum
			if net.Server {
				diff = ct.GetShare().Sub(v.GetShare())
			} else {
				diff = v.GetShare().Sub(ct.GetShare())
			}
			diffb := misc.BytesToBools(diff.Bytes())
			queue := make(chan pub.Bool, ct.Length())
			for i := 0; i < ct.Length(); i++ {
				queue <- pub.Bool(diffb[i])
			}
			rl := emptyAshare(net.Server)
			rr := emptyAshare(net.Server)
			for len(queue) > 1 {
				rl.SetShare(<-queue)
				rr.SetShare(<-queue)
				queue <- rl.Or(net, rr).GetShare().(pub.Bool)
			}
			rl.SetShare(<-queue)
			ret = rl.Not(net)

		} else if vx, ok := x.(pub.PubNum); ok {
			var diff pub.PubNum
			if net.Server {
				diff = ct.GetShare().Sub(vx)
			} else {
				diff = ct.GetShare()
			}
			diffb := misc.BytesToBools(diff.Bytes())
			queue := make(chan pub.Bool, ct.Length())
			for i := 0; i < ct.Length(); i++ {
				queue <- pub.Bool(diffb[i])
			}
			rl := emptyAshare(net.Server)
			rr := emptyAshare(net.Server)
			for len(queue) > 1 {
				rl.SetShare(<-queue)
				rr.SetShare(<-queue)
				queue <- rl.Or(net, rr).GetShare().(pub.Bool)
			}
			rl.SetShare(<-queue)
			ret = rl.Not(net)
		} else {
			log.Panicln("")

		}
	}
	return ret
}
func (ct *Ashare) Gt(net network.Network, x value.Value) PvtNum {
	return nil
}
func (ct *Ashare) Lt(net network.Network, x value.Value) PvtNum {
	return nil
}

func (ct *Ashare) Shr(net network.Network, x value.Value) PvtNum {
	ret := new(Ashare)
	if net.Server {
		ret.Pos = 0
	} else {
		ret.Pos = 1
	}
	if v, ok := x.(pub.PubNum); ok {
		if v.Lt(v.From(0)).(pub.Bool) {
			log.Panicln("Shr support positive only")
		} else {
			ret.SetShare(ct.GetShare().Shr(x.(pub.Int8)))
		}
	} else {
		log.Panicln("Shl support PubNum only")
	}
	return ret
}

func (ct *Ashare) Shl(net network.Network, x value.Value) PvtNum {
	ret := new(Ashare)
	if net.Server {
		ret.Pos = 0
	} else {
		ret.Pos = 1
	}
	if v, ok := x.(pub.PubNum); ok {
		if v.Lt(v.From(0)).(pub.Bool) {
			log.Panicln("Shr support positive only")
		} else {
			ret.SetShare(ct.GetShare().Shl(x.(pub.Int8)))
		}
	} else {
		log.Panicln("Shl support PubNum only")
	}
	return ret
}
func (ct *Ashare) Mux(net network.Network, x value.Value, y value.Value) PvtNum {
	var ret PvtNum
	var b PvtNum
	if x.Length() == 1 {
		lop := ct.And(net, x)
		rop := ct.Not(net).And(net, y)
		tmp := emptyAshare(net.Server)
		tmp.SetShare(pub.Bool(lop.GetShare().(pub.Bool) != rop.GetShare().(pub.Bool)))
		ret = tmp
		return ret
	} else {
		if vx, ok := x.(PvtNum); ok {
			b = ct.ExpandBooleanTo(net, vx.GetShare())
		} else if vx, ok := x.(pub.PubNum); ok {
			b = ct.ExpandBooleanTo(net, vx)
		}
		ret = b.Mul(net, x).Add(net, y).Sub(net, b.Mul(net, y))
		return ret
	}
}

/*
################################################

	Self-own methods

################################################
*/
// all party get the unshare
func (ct *Ashare) Declassify(net network.Network) {
	net.Send(network.NewMsg(00, network.Sharing, ct.GetShare().Bytes()))
	y := net.Recv()
	ct.shares[1-ct.Pos] = ct.GetShare().Decode(y.Data)
	if ct.GetShare().Length() == 1 {
		// XOR for boolean
		ct.Plaintext = pub.Bool(ct.shares[0].(pub.Bool) != ct.shares[1].(pub.Bool))
	} else {
		// ADD for integer
		ct.Plaintext = ct.shares[0].Add(ct.shares[1])
	}
}

func (ct *Ashare) GetShares() [2]pub.PubNum {
	return ct.shares
}
func (ct *Ashare) SetShare(x pub.PubNum) {
	ct.shares[ct.Pos] = x
}

func (ct *Ashare) ExpandBooleanTo(net network.Network, target pub.PubNum) PvtNum {
	x := emptyAshare(net.Server)
	y := emptyAshare(net.Server)
	if net.Server {
		if ct.GetShare().(pub.Bool) == pub.ZeroBool {
			x.SetShare(target.From(0))
		} else {
			x.SetShare(target.From(1))
		}
		y.SetShare(target.From(0))
	} else {
		if ct.GetShare().(pub.Bool) == pub.ZeroBool {
			y.SetShare(target.From(0))
		} else {
			y.SetShare(target.From(1))
		}
		x.SetShare(target.From(0))
	}
	return x.Add(net, y).Sub(net, x.Mul(net, y).Mul(net, target.From(2)))
}
