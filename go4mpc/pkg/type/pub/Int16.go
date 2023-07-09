package pub

import (
	"fmt"
	"log"
	"reflect"
)

type Int16 int16

func (ct Int16) Rand() PubNum {
	return Int16(rand_.Intn(65536))
}

func (ct Int16) From(i int) PubNum {
	return Int16(i)
}

const ZeroInt16 = Int16(0)

func (ct Int16) Public() {}
func (ct Int16) Number() {}

func (ct Int16) Length() int {
	return 16
}
func (ct Int16) BinaryString() string {
	return fmt.Sprintf("%08b", byte(ct))
}
func (ct Int16) Bytes() []byte {
	return append([]byte{}, byte(ct), byte(ct>>8))
}
func (ct Int16) Decode(b []byte) PubNum {
	_ = b[1]
	return Int16(int16(b[0]) | int16(b[1])<<8)
}
func (ct Int16) Sub(b PubNum) PubNum {
	return ct - b.(Int16)
}
func (ct Int16) Add(b PubNum) PubNum {
	return ct + b.(Int16)
}
func (ct Int16) Mul(b PubNum) PubNum {
	return ct * b.(Int16)
}
func (ct Int16) Div(b PubNum) PubNum {
	return ct / b.(Int16)
}

func (ct Int16) Eq(b PubNum) PubNum {
	return Bool(ct == b.(Int16))
}
func (ct Int16) Gt(b PubNum) PubNum {
	return Bool(ct > b.(Int16))
}
func (ct Int16) Lt(b PubNum) PubNum {
	return Bool(ct < b.(Int16))
}

func (ct Int16) Shr(b PubNum) PubNum {
	return ct >> b.(Int8)
}
func (ct Int16) Shl(b PubNum) PubNum {
	return ct << b.(Int8)
}

func (ct Int16) Not() PubNum {
	log.Panicln("Invalid Operator Not() of Int16")
	return nil
}
func (ct Int16) And(b PubNum) PubNum {
	log.Panicf("Invalid Operator And() of (Int16, %s) ", reflect.TypeOf(b))
	return nil
}
func (ct Int16) Or(b PubNum) PubNum {
	log.Panicln("Invalid Operator Not() of Int16")
	return nil
}
func (Ct Int16) Mux(x PubNum, y PubNum) PubNum {
	log.Panicln("Invalid Operator Mux() of Int16")
	return nil
}
