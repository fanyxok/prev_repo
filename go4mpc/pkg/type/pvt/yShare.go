package pvt

import (
	"log"
	"math/rand"
	"reflect"
	"s3l/mpcfgo/config"
	"s3l/mpcfgo/internal/encrypt/prf"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/internal/ot/baseot"
	"s3l/mpcfgo/pkg/bristol"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/value"
	"time"
)

var rand_ = rand.New(rand.NewSource(time.Now().UnixNano()))
var DELTA []byte
var YShare = new(Yshare)

func init() {
	DELTA = make([]byte, config.SymByte+1)
	rand_.Read(DELTA)
	DELTA[config.SymByte] = 1
}

// Yshare The label and pointer are stored in a []byte, [0:len-1]byte, is the label, and [len-1]byte is the pointer
type Yshare struct {
	// Yshare is a share of l bit, Plaintext is a l-bit number
	Plaintext pub.PubNum
	wValue    [][]byte    // wValue[i] is the i-th bit's label and pointer, activated
	wTable    [][2][]byte // wValue[i][0] is the i-th bit's label and pointer of value 0, wValue[i][1] of value 1
}

func (ct *Yshare) GetShare() pub.PubNum {
	return pub.Zero(ct.Length())
}

func (ct *Yshare) GetPlaintext() pub.PubNum {
	return ct.Plaintext
}

func newYshare(length int) *Yshare {
	nShare := new(Yshare)
	nShare.wTable = make([][2][]byte, length)
	for i := 0; i < length; i++ {
		nShare.wTable[i][0] = make([]byte, config.SymByte+1)
		nShare.wTable[i][1] = make([]byte, config.SymByte+1)
		rand_.Read(nShare.wTable[i][0])
		misc.BytesXorBytes(nShare.wTable[i][1], nShare.wTable[i][0], DELTA)
		if randBool() {
			nShare.wTable[i][0][config.SymByte] = 1
			nShare.wTable[i][1][config.SymByte] = 0
		} else {
			nShare.wTable[i][0][config.SymByte] = 0
			nShare.wTable[i][1][config.SymByte] = 1
		}
	}
	return nShare
}

func newYShareOfValue(length int, x pub.PubNum) *Yshare {
	nShare := new(Yshare)
	nShare.wTable = make([][2][]byte, length)
	nShare.Plaintext = x
	bools := misc.BytesToBools(x.Bytes())
	bools = bools[:length]
	for i := 0; i < length; i++ {
		nShare.wTable[i][0] = make([]byte, config.SymByte+1)
		nShare.wTable[i][1] = make([]byte, config.SymByte+1)
		if bools[i] {
			copy(nShare.wTable[i][0], DELTA)
		} else {
			copy(nShare.wTable[i][1], DELTA)
		}
	}
	return nShare
}
func emptyYShare(length int) *Yshare {
	nShare := new(Yshare)
	nShare.wValue = make([][]byte, length)
	nShare.Plaintext = pub.Zero(length)
	for i := 0; i < length; i++ {
		nShare.wValue[i] = make([]byte, config.SymByte+1)
	}
	return nShare
}
func (ct *Yshare) New(net network.Network, num pub.PubNum) PvtNum {
	var nShare *Yshare
	if net.Server {
		// garbler provide the number
		nShare = newYshare(num.Length())
		nShare.Plaintext = num
		net.Send(network.NewMsg(00, network.Yshare, []byte{byte(nShare.Length())}))
		bits := misc.BytesToBools(nShare.Plaintext.Bytes())
		if nShare.Length() != 1 && len(bits) != nShare.Length() {
			log.Panicln("")
		}
		bits = bits[:nShare.Length()]

		for i, v := range bits {
			if v {
				net.Send(network.NewMsg(0, network.Yshare, nShare.wTable[i][1]))
			} else {
				net.Send(network.NewMsg(0, network.Yshare, nShare.wTable[i][0]))
			}
		}
	} else {
		nShare = new(Yshare)
		nShare.Plaintext = num
		net.Send(network.NewMsg(99, network.Yshare, []byte{byte(nShare.Length())}))
		bits := misc.BytesToBools(nShare.Plaintext.Bytes())
		if num.Length() == 1 {
			bits = bits[:1]
		}

		nShare.wValue = baseot.RecvN(net, bits)
	}
	return nShare
}
func (ct *Yshare) NewN(net network.Network, num []pub.PubNum) []PvtNum {
	return nil
}
func (ct *Yshare) NewFromN(net network.Network) []PvtNum {
	return nil
}
func (ct *Yshare) NewFrom(net network.Network) PvtNum {
	var nShare *Yshare
	if net.Server {
		zero := pub.Zero(int(net.Recv().Data[0]))
		nShare = newYshare(zero.Length())
		nShare.Plaintext = zero
		wtable0 := make([][]byte, nShare.Length())
		wtable1 := make([][]byte, nShare.Length())
		for i := 0; i < nShare.Length(); i++ {
			wtable0[i] = nShare.wTable[i][0]
			wtable1[i] = nShare.wTable[i][1]
		}
		baseot.SendN(net, wtable0, wtable1)
	} else {
		nShare = new(Yshare)
		nShare.Plaintext = pub.Zero(int(net.Recv().Data[0]))
		nShare.wValue = make([][]byte, nShare.Length())
		for i := 0; i < nShare.Length(); i++ {
			nShare.wValue[i] = net.Recv().Data
		}
	}
	return nShare
}

func (ct *Yshare) decode0(net network.Network) {
	dTable := make([]byte, ct.Plaintext.Length()*2)
	for i, v := range ct.wTable {
		h0 := prf.FixedKeyAES.Hash(v[0])[0] & 1
		h1 := prf.FixedKeyAES.Hash(v[1])[0] & 1
		if v[0][config.SymByte] == 0 {
			dTable[i] = h0
			dTable[i+ct.Plaintext.Length()] = 1 - h1
		} else {
			dTable[i] = 1 - h1
			dTable[i+ct.Plaintext.Length()] = h0
		}
	}
	net.Send(network.NewMsg(00, network.Yshare, dTable))
	ct.Plaintext = ct.Plaintext.Decode(net.Recv().Data)
}

//go:noinline
func (ct *Yshare) decode1(net network.Network) {
	dTable := net.Recv().Data
	pv := make([]bool, ct.Length())
	for i := 0; i < ct.Length(); i++ {
		lsb := prf.FixedKeyAES.Hash(ct.wValue[i])[0] & 1
		if ct.wValue[i][config.SymByte] == 0 {
			pv[i] = dTable[i]^lsb == 1
		} else {
			pv[i] = dTable[i+ct.Length()]^lsb == 1
		}
	}
	bits := misc.BoolsToBytes(pv)
	ct.Plaintext = ct.Plaintext.Decode(bits)
	net.Send(network.NewMsg(00, network.Yshare, bits))
}

func (ct *Yshare) Declassify(net network.Network) {
	if net.Server {
		ct.decode0(net)
	} else {
		ct.decode1(net)
	}
}

func (ct *Yshare) GetShares() [][2][]byte {
	return ct.wTable
}

func (ct *Yshare) Label() [][]byte {
	return nil
}

func (ct *Yshare) Pointer() [][]byte {
	return nil
}

/*
################################################

	PvtNum interface impl

################################################
*/
func (ct *Yshare) Number()     {}
func (ct *Yshare) Length() int { return ct.Plaintext.Length() }
func (ct *Yshare) Private()    {}
func (ct *Yshare) Add(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	c := bristol.AddCirc[x.Length()]
	z := Eval(net, c, ct, x.(*Yshare))
	return z
}

func (ct *Yshare) Sub(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	c := bristol.SubCirc[ct.Length()]
	z := Eval(net, c, ct, x.(*Yshare))
	return z
}

func (ct *Yshare) Mul(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	c := bristol.MulCirc[ct.Length()]
	z := Eval(net, c, ct, x.(*Yshare))
	return z
}

func (ct *Yshare) Div(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	c := bristol.DivCirc[ct.Length()]
	z := Eval(net, c, ct, x.(*Yshare))
	return z
}

func (ct *Yshare) Not(net network.Network) PvtNum {
	z := Eval(net, bristol.NOTb1, ct)
	return z
}

func (ct *Yshare) And(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	z := Eval(net, bristol.ANDb1, ct, x.(*Yshare))
	return z
}
func (ct *Yshare) Or(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	z := Eval(net, bristol.ORb1, ct, x.(*Yshare))
	return z
}
func (ct *Yshare) Eq(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	c := bristol.EqCirc[ct.Length()]
	z := Eval(net, c, ct, x.(*Yshare))
	return z
}
func (ct *Yshare) Gt(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	c := bristol.GtCirc[ct.Length()]
	z := Eval(net, c, ct, x.(*Yshare))
	return z
}
func (ct *Yshare) Lt(net network.Network, x value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	c := bristol.LtCirc[ct.Length()]
	z := Eval(net, c, ct, x.(*Yshare))
	return z
}
func (ct *Yshare) Shr(net network.Network, x value.Value) (z PvtNum) {
	switch v := x.(type) {
	case pub.Int8:
		nShare := new(Yshare)
		nShare.Plaintext = ct.Plaintext
		if net.Server {
			nShare.wTable = make([][2][]byte, ct.Length())
			if step := int(v); step < 0 {
				log.Panicln("YShare Shr(): Do Shr with neg value")
			} else if step < ct.Length()-1 {
				for i := 0; i < ct.Length()-step-1; i++ {
					nShare.wTable[i][0] = make([]byte, config.SymByte+1)
					nShare.wTable[i][1] = make([]byte, config.SymByte+1)
					copy(nShare.wTable[i][0], ct.wTable[i+step][0])
					copy(nShare.wTable[i][1], ct.wTable[i+step][1])
				}
				for i := ct.Length() - step - 1; i < ct.Length()-1; i++ {
					nShare.wTable[i][0] = make([]byte, config.SymByte+1)
					nShare.wTable[i][1] = make([]byte, config.SymByte+1)
					copy(nShare.wTable[i][0], ct.wTable[ct.Length()-1][0])
					copy(nShare.wTable[i][1], ct.wTable[ct.Length()-1][1])
				}
			} else {
				for i := 0; i < ct.Length()-1; i++ {
					nShare.wTable[i][0] = make([]byte, config.SymByte+1)
					nShare.wTable[i][1] = make([]byte, config.SymByte+1)
					copy(nShare.wTable[i][0], ct.wTable[ct.Length()-1][0])
					copy(nShare.wTable[i][1], ct.wTable[ct.Length()-1][1])
				}
			}
			nShare.wTable[ct.Length()-1][0] = make([]byte, config.SymByte+1)
			nShare.wTable[ct.Length()-1][1] = make([]byte, config.SymByte+1)
			copy(nShare.wTable[ct.Length()-1][0], ct.wTable[ct.Length()-1][0])
			copy(nShare.wTable[ct.Length()-1][1], ct.wTable[ct.Length()-1][1])
		} else {
			nShare.wValue = make([][]byte, ct.Length())
			if step := int(v); step < 0 {
				log.Panicln("YShare Shr(): Do Shr with neg value")
			} else if step < ct.Length() {
				for i := 0; i < ct.Length()-step-1; i++ {
					nShare.wValue[i] = make([]byte, config.SymByte+1)
					copy(nShare.wValue[i], ct.wValue[i+step])
				}
				for i := ct.Length() - step - 1; i < ct.Length()-1; i++ {
					nShare.wValue[i] = make([]byte, config.SymByte+1)
					copy(nShare.wValue[i], ct.wValue[ct.Length()-1])
				}
			} else {
				for i := 0; i < ct.Length()-1; i++ {
					nShare.wValue[i] = make([]byte, config.SymByte+1)
					copy(nShare.wValue[i], ct.wValue[ct.Length()-1])
				}
			}
			nShare.wValue[ct.Length()-1] = make([]byte, config.SymByte+1)
			copy(nShare.wValue[ct.Length()-1], ct.wValue[ct.Length()-1])
		}
		z = nShare
	case PvtNum:
		log.Panicln("YShare Shr(): *YShare can only do Shr by PubNum(Int8)")
	default:
		log.Panicln("YShare Shr(): value x has unexpected interface type:", reflect.TypeOf(x))
	}
	return z
}
func (ct *Yshare) Shl(net network.Network, x value.Value) (z PvtNum) {
	switch v := x.(type) {
	case pub.Int8:
		nShare := new(Yshare)
		nShare.Plaintext = ct.Plaintext
		if net.Server {
			nShare.wTable = make([][2][]byte, ct.Length())
			if step := int(v); step < 0 {
				log.Panicln("YShare Shr(): Do Shr with neg value")
			} else if step < ct.Length() {
				for i := step; i < ct.Length(); i++ {
					nShare.wTable[i][0] = make([]byte, config.SymByte+1)
					nShare.wTable[i][1] = make([]byte, config.SymByte+1)
					copy(nShare.wTable[i][0], ct.wTable[i-step][0])
					copy(nShare.wTable[i][1], ct.wTable[i-step][1])
				}
				for i := 0; i < step; i++ {
					nShare.wTable[i][0] = make([]byte, config.SymByte+1)
					nShare.wTable[i][1] = make([]byte, config.SymByte+1)
					copy(nShare.wTable[i][1], DELTA)
				}
			} else {
				// all zero
				for i := 0; i < ct.Length(); i++ {
					nShare.wTable[i][0] = make([]byte, config.SymByte+1)
					nShare.wTable[i][1] = make([]byte, config.SymByte+1)
					copy(nShare.wTable[i][1], DELTA)
				}
			}
		} else {
			nShare.wValue = make([][]byte, ct.Length())
			if step := int(v); step < 0 {
				log.Panicln("YShare Shr(): Do Shr with neg value")
			} else if step < ct.Length() {
				for i := step; i < ct.Length(); i++ {
					nShare.wValue[i] = make([]byte, config.SymByte+1)
					copy(nShare.wValue[i], ct.wValue[i-step])
				}
				for i := 0; i < step; i++ {
					nShare.wValue[i] = make([]byte, config.SymByte+1)
				}
			} else {
				for i := 0; i < ct.Length(); i++ {
					nShare.wValue[i] = make([]byte, config.SymByte+1)
				}
			}
		}
		z = nShare
	case PvtNum:
		log.Panicln("YShare Shr(): *YShare can only do Shr by PubNum(Int8)")
	default:
		log.Panicln("YShare Shr(): value x has unexpected interface type:", reflect.TypeOf(x))
	}
	return z
}

func (ct *Yshare) Mux(net network.Network, x value.Value, y value.Value) PvtNum {
	if v, ok := x.(pub.PubNum); ok {
		if net.Server {
			x = newYShareOfValue(x.Length(), v)
		} else {
			x = emptyYShare(x.Length())
		}
	}
	if v, ok := y.(pub.PubNum); ok {
		if net.Server {
			y = newYShareOfValue(y.Length(), v)
		} else {
			y = emptyYShare(y.Length())
		}
	}
	c := bristol.MuxCirc[x.Length()]
	x_ := x.(*Yshare)
	y_ := y.(*Yshare)
	z := Eval(net, c, ct, x_, y_)
	return z
}

func Eval(net network.Network, bc bristol.BristolCircuit, ins ...*Yshare) *Yshare {
	// assert check
	if len(bc.InLength) != len(ins) {
		log.Panicln("YShare Eval Bristol Circuit Has", bc.InLength, "Inputs, While Given", len(ins), "Inputs.")
	}
	for i, v := range bc.InLength {
		if v != ins[i].Length() {
			log.Panicln("")
		}
	}
	if net.Server {
		// garbler side
		gWires := make([][2][]byte, bc.Wire)
		idx := 0
		for _, v := range ins {
			for _, z := range v.wTable {
				gWires[idx] = z
				idx++
			}
		}
		for _, v := range bc.Gate {
			switch v.Type {
			case bristol.AND:
				gTable := [2][]byte{}
				lhs := gWires[v.InWire[0]]
				rhs := gWires[v.InWire[1]]
				r := lhs[0][config.SymByte]
				var crKeyP []byte
				if rhs[0][config.SymByte] == 0 {
					crKeyP = prf.FixedKeyAES.Hash(rhs[0])
					e := prf.FixedKeyAES.Hash(rhs[1])
					misc.BytesXorBytes(e, e, crKeyP)
					if r == 1 {
						misc.BytesXorBytes(e, e, DELTA)
					}
					gTable[0] = e
				} else {
					crKeyP = prf.FixedKeyAES.Hash(rhs[1])
					if r == 1 {
						misc.BytesXorBytes(crKeyP, crKeyP, DELTA)
					}
					e := prf.FixedKeyAES.Hash(rhs[0])
					misc.BytesXorBytes(e, e, crKeyP)
					gTable[0] = e
				}
				var clKeyP []byte
				if r == 0 {
					clKeyP = prf.FixedKeyAES.Hash(lhs[0])
					e := prf.FixedKeyAES.Hash(lhs[1])
					misc.BytesXorBytes(e, e, clKeyP)
					misc.BytesXorBytes(e, e, rhs[0])
					gTable[1] = e
				} else {
					clKeyP = prf.FixedKeyAES.Hash(lhs[1])
					e := prf.FixedKeyAES.Hash(lhs[0])
					misc.BytesXorBytes(e, e, clKeyP)
					misc.BytesXorBytes(e, e, rhs[0])
					gTable[1] = e
				}
				gWires[v.OutWire][0] = make([]byte, config.SymByte+1)
				misc.BytesXorBytes(gWires[v.OutWire][0], clKeyP, crKeyP)
				gWires[v.OutWire][1] = make([]byte, config.SymByte+1)
				misc.BytesXorBytes(gWires[v.OutWire][1], gWires[v.OutWire][0], DELTA)
				net.Send(network.NewMsg(00, network.Yshare, gTable[0]))
				net.Send(network.NewMsg(00, network.Yshare, gTable[1]))
			case bristol.EQ:
				gWires[v.OutWire][0] = make([]byte, config.SymByte+1)
				gWires[v.OutWire][1] = make([]byte, config.SymByte+1)
				rand_.Read(gWires[v.OutWire][0])
				misc.BytesXorBytes(gWires[v.OutWire][1], gWires[v.OutWire][0], DELTA)
				if v.InWire[0] == 0 {
					net.Send(network.NewMsg(00, network.Yshare, gWires[v.OutWire][0]))
				} else if v.InWire[0] == 1 {
					net.Send(network.NewMsg(00, network.Yshare, gWires[v.OutWire][1]))
				} else {
					log.Panicln("Eval() EQ: Can't assign non-0/1 value to a wire")
				}
			case bristol.EQW:
				gWires[v.OutWire] = gWires[v.InWire[0]]
			case bristol.NOT:
				gWires[v.OutWire][0] = make([]byte, config.SymByte+1)
				gWires[v.OutWire][1] = make([]byte, config.SymByte+1)
				copy(gWires[v.OutWire][0], gWires[v.InWire[0]][1])
				copy(gWires[v.OutWire][1], gWires[v.InWire[0]][0])
				gWires[v.OutWire][0][config.SymByte] = 1 - gWires[v.OutWire][0][config.SymByte]
				gWires[v.OutWire][1][config.SymByte] = 1 - gWires[v.OutWire][1][config.SymByte]
			case bristol.XOR:
				lhs := gWires[v.InWire[0]]
				rhs := gWires[v.InWire[1]]
				gWires[v.OutWire][0] = make([]byte, config.SymByte+1)
				misc.BytesXorBytes(gWires[v.OutWire][0], lhs[0], rhs[0])
				gWires[v.OutWire][1] = make([]byte, config.SymByte+1)
				misc.BytesXorBytes(gWires[v.OutWire][1], gWires[v.OutWire][0], DELTA)
			default:
				log.Panicln("")
			}
		}
		ret := new(Yshare)
		ret.wTable = gWires[bc.Wire-bc.OutLength:]
		ret.Plaintext = pub.Zero(len(ret.wTable))
		return ret
	} else {
		// evaluator side
		gWires := make([][]byte, bc.Wire)
		idx := 0
		for _, v := range ins {
			// for _, z := range v.wValue {
			// 	gWires[idx] = z
			// 	idx++
			// }
			for i := 0; i < v.Length(); i++ {
				gWires[idx] = v.wValue[i]
				idx++
			}
		}
		for _, v := range bc.Gate {
			switch v.Type {
			case bristol.AND:
				gTable := [2][]byte{}
				gTable[0] = net.Recv().Data
				gTable[1] = net.Recv().Data
				lhs := gWires[v.InWire[0]]
				rhs := gWires[v.InWire[1]]
				gHalf := prf.FixedKeyAES.Hash(rhs)
				if rhs[config.SymByte] == 1 {
					misc.BytesXorBytes(gHalf, gHalf, gTable[0])
				}
				eHalf := prf.FixedKeyAES.Hash(lhs)
				if lhs[config.SymByte] == 1 {
					misc.BytesXorBytes(eHalf, eHalf, gTable[1])
					misc.BytesXorBytes(eHalf, eHalf, rhs)
				}
				gWires[v.OutWire] = make([]byte, config.SymByte+1)
				misc.BytesXorBytes(gWires[v.OutWire], gHalf, eHalf)
			case bristol.EQ:
				gWires[v.OutWire] = net.Recv().Data
			case bristol.EQW:
				gWires[v.OutWire] = gWires[v.InWire[0]]
			case bristol.NOT:
				gWires[v.OutWire] = make([]byte, config.SymByte+1)
				copy(gWires[v.OutWire], gWires[v.InWire[0]])
				gWires[v.OutWire][config.SymByte] = 1 - gWires[v.OutWire][config.SymByte]
			case bristol.XOR:
				gWires[v.OutWire] = make([]byte, config.SymByte+1)
				misc.BytesXorBytes(gWires[v.OutWire], gWires[v.InWire[0]], gWires[v.InWire[1]])
			default:
				log.Panicln("")
			}
		}
		ret := new(Yshare)
		ret.wValue = gWires[bc.Wire-bc.OutLength:]
		ret.Plaintext = pub.Zero(len(ret.wValue))
		return ret
	}
}

// Useful function
type boolpool struct {
	src       rand.Source
	cache     int64
	remaining int
}

var bp = boolpool{src: rand.NewSource(time.Now().UnixNano())}

func randBool() bool {
	if bp.remaining == 0 {
		bp.cache, bp.remaining = bp.src.Int63(), 63
	}
	result := bp.cache&0x01 == 1
	bp.cache >>= 1
	bp.remaining--
	return result
}
