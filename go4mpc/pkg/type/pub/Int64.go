package pub

import (
	"fmt"
	"log"
	"reflect"
)

type Int64 int64

func (ct Int64) Rand() PubNum {
	return Int64(rand_.Uint64())
}
func (ct Int64) From(r int) PubNum {
	return Int64(r)
}

const ZeroInt64 = Int64(0)

func (ct Int64) Number() {}

func (ct Int64) Public() {}

func (ct Int64) Length() int {
	return 64
}

func (ct Int64) BinaryString() string {
	return fmt.Sprintf("%08b", ct.Bytes())
}

func (ct Int64) Bytes() []byte {
	return append([]byte{}, byte(ct), byte(ct>>8), byte(ct>>16), byte(ct>>24), byte(ct>>32), byte(ct>>40), byte(ct>>48), byte(ct>>56))
}
func (ct Int64) Decode(b []byte) PubNum {
	_ = b[7]
	return Int64(int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56)
}

func (ct Int64) Add(b PubNum) PubNum {
	return ct + b.(Int64)
}
func (ct Int64) Sub(b PubNum) PubNum {
	return ct - b.(Int64)
}
func (ct Int64) Mul(b PubNum) PubNum {
	return ct * b.(Int64)
}
func (ct Int64) Div(b PubNum) PubNum {
	return ct / b.(Int64)
}
func (ct Int64) Eq(b PubNum) PubNum {
	return Bool(ct == b.(Int64))
}
func (ct Int64) Gt(b PubNum) PubNum {
	return Bool(ct > b.(Int64))
}
func (ct Int64) Lt(b PubNum) PubNum {
	return Bool(ct < b.(Int64))
}
func (ct Int64) Shr(b PubNum) PubNum {
	return ct >> b.(Int8)
}
func (ct Int64) Shl(b PubNum) PubNum {
	return ct << b.(Int8)
}

func (ct Int64) Not() PubNum {
	log.Panicln("Invalid Operator Not() of Int64")
	return nil
}
func (ct Int64) And(b PubNum) PubNum {
	log.Panicf("Invalid Operator And() of (Int64, %s) ", reflect.TypeOf(b))
	return nil
}
func (ct Int64) Or(b PubNum) PubNum {
	log.Panicln("Invalid Operator Not() of Int64")
	return nil
}
func (Ct Int64) Mux(x PubNum, y PubNum) PubNum {
	log.Panicln("Invalid Operator Mux() of Int64")
	return nil
}
