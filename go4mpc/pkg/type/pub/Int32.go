package pub

import (
	"fmt"
	"log"
	"math"
	"reflect"
)

type Int32 int32

func (ct Int32) Rand() PubNum {
	rng := math.MaxUint32 + 1
	return Int32(rand_.Intn(rng) + math.MinInt32)
}
func (ct Int32) From(r int) PubNum {
	return Int32(r)
}

const ZeroInt32 = Int32(0)

func (ct Int32) Public() {}
func (ct Int32) Number() {}

func (ct Int32) Length() int {
	return 32
}

func (ct Int32) BinaryString() string {
	return fmt.Sprintf("%08b", byte(ct))
}

func (ct Int32) Bytes() []byte {
	return append([]byte{}, byte(ct), byte(ct>>8), byte(ct>>16), byte(ct>>24))
}
func (ct Int32) Decode(b []byte) PubNum {
	_ = b[3]
	return Int32(int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24)
}

func (ct Int32) Sub(b PubNum) PubNum {
	return ct - b.(Int32)
}
func (ct Int32) Add(b PubNum) PubNum {
	return ct + b.(Int32)
}
func (ct Int32) Mul(b PubNum) PubNum {
	return ct * b.(Int32)
}
func (ct Int32) Div(b PubNum) PubNum {
	return ct / b.(Int32)
}

func (ct Int32) Eq(b PubNum) PubNum {
	return Bool(ct == b.(Int32))
}
func (ct Int32) Gt(b PubNum) PubNum {
	return Bool(ct > b.(Int32))
}
func (ct Int32) Lt(b PubNum) PubNum {
	return Bool(ct < b.(Int32))
}

func (ct Int32) Shr(b PubNum) PubNum {
	return ct >> b.(Int8)
}
func (ct Int32) Shl(b PubNum) PubNum {
	return ct << b.(Int8)
}

func (ct Int32) Not() PubNum {
	log.Panicln("Invalid Operator Not() of Int32")
	return nil
}
func (ct Int32) And(b PubNum) PubNum {
	log.Panicf("Invalid Operator And() of (Int32, %s) ", reflect.TypeOf(b))
	return nil
}
func (ct Int32) Or(b PubNum) PubNum {
	log.Panicln("Invalid Operator Not() of Int32")
	return nil
}
func (Ct Int32) Mux(x PubNum, y PubNum) PubNum {
	log.Panicln("Invalid Operator Mux() of Int32")
	return nil
}
