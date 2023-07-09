package pub

import (
	"log"
)

// signed, range from -128 to 127
type Bool bool

func (ct Bool) Rand() PubNum {
	if rand_.Intn(2) == 0 {
		return Bool(false)
	} else {
		return Bool(true)
	}
}

func (ct Bool) From(x int) PubNum {
	if x == 0 {
		return Bool(false)
	} else if x == 1 {
		return Bool(true)
	}
	log.Panicf("Bool From() Invalid %d", x)
	return Bool(true)
}

const ZeroBool = Bool(false)

func (ct Bool) Number() {}
func (ct Bool) Public() {}

func (ct Bool) Length() int { return 1 }
func (ct Bool) BinaryString() string {
	if ct {
		return "0"
	} else {
		return "1"
	}
}
func (ct Bool) Bytes() []byte {
	if ct {
		return []byte{1}
	} else {
		return []byte{0}
	}
}
func (ct Bool) Decode(b []byte) PubNum {
	_ = b[0]
	if b[0] == 0 {
		return Bool(false)
	} else {
		return Bool(true)
	}
}

// Operator
func (ct Bool) Sub(b PubNum) PubNum {
	log.Panicf("Invalid Operation Sub() of Bool: %v", b)
	return nil
}
func (ct Bool) Add(b PubNum) PubNum {
	log.Panicf("Invalid Operation Add() of Bool: %v", b)
	return nil
}
func (ct Bool) Mul(b PubNum) PubNum {
	log.Panicf("Invalid Operation Mul() of Bool: %v", b)
	return nil
}
func (ct Bool) Div(b PubNum) PubNum {
	log.Panicf("Invalid Operation Div() of Bool: %v", b)
	return nil
}
func (ct Bool) Not() PubNum {
	return Bool(!ct)
}
func (ct Bool) And(b PubNum) PubNum {
	return Bool(ct && b.(Bool))
}
func (ct Bool) Or(b PubNum) PubNum {
	return Bool(ct || b.(Bool))
}
func (ct Bool) Eq(b PubNum) PubNum {
	return Bool(ct == b.(Bool))
}
func (ct Bool) Gt(b PubNum) PubNum {
	log.Panicf("Invalid Operation Gt() of Bool: %v", b)
	return nil
}
func (ct Bool) Lt(b PubNum) PubNum {
	log.Panicf("Invalid Operation Lt() of Bool: %v", b)
	return nil
}
func (ct Bool) Shr(b PubNum) PubNum {
	log.Panicf("Invalid Operation Shr() of Bool: %v", b)
	return nil
}
func (ct Bool) Shl(b PubNum) PubNum {
	log.Panicf("Invalid Operation Shl() of Bool: %v", b)
	return nil
}
func (ct Bool) Mux(x PubNum, y PubNum) PubNum {
	if ct {
		return x
	} else {
		return y
	}
}
func (ct Bool) Xor(x PubNum) Bool {
	return ct != x.(Bool)
}
