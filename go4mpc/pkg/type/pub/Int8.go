package pub

import (
	"fmt"
	"log"
	"reflect"
)

// signed, range from -128 to 127
type Int8 int8

func (ct Int8) Rand() PubNum {
	return Int8(rand_.Intn(256))
}

func (ct Int8) From(x int) PubNum {
	return Int8(x)
}

const ZeroInt8 = Int8(0)

func (ct Int8) Number() {}
func (ct Int8) Public() {}

func (ct Int8) Length() int { return 8 }
func (ct Int8) BinaryString() string {
	return fmt.Sprintf("%08b", byte(ct))
}
func (ct Int8) Bytes() []byte {
	return []byte{byte(ct)}
}
func (ct Int8) Decode(b []byte) PubNum {
	_ = b[0]
	return Int8(b[0])
}

// Operator
func (ct Int8) Sub(b PubNum) PubNum {
	return ct - b.(Int8)
}
func (ct Int8) Add(b PubNum) PubNum {
	return ct + b.(Int8)
}
func (ct Int8) Mul(b PubNum) PubNum {
	return ct * b.(Int8)
}
func (ct Int8) Div(b PubNum) PubNum {
	return ct / b.(Int8)
}

func (ct Int8) Eq(b PubNum) PubNum {
	return Bool(ct == b.(Int8))
}
func (ct Int8) Gt(b PubNum) PubNum {
	return Bool(ct > b.(Int8))
}
func (ct Int8) Lt(b PubNum) PubNum {
	return Bool(ct < b.(Int8))
}

func (ct Int8) Shr(b PubNum) PubNum {
	return ct >> b.(Int8)
}
func (ct Int8) Shl(b PubNum) PubNum {
	return ct << b.(Int8)
}

func (ct Int8) Not() PubNum {
	log.Panicln("Invalid Operator Not() of Int8")
	return nil
}
func (ct Int8) And(b PubNum) PubNum {
	log.Panicf("Invalid Operator And() of (Int8, %s) ", reflect.TypeOf(b))
	return nil
}
func (ct Int8) Or(b PubNum) PubNum {
	log.Panicln("Invalid Operator Not() of Int8")
	return nil
}
func (Ct Int8) Mux(x PubNum, y PubNum) PubNum {
	log.Panicln("Invalid Operator Mux() of Int8")
	return nil
}
