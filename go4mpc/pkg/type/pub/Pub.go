package pub

import (
	"log"
	"math/rand"
	number "s3l/mpcfgo/pkg/type/value"
	"time"
)

var rand_ = rand.New(rand.NewSource(time.Now().UnixNano()))

type PubNum interface {
	number.Value
	Public() // dummy method for PubNum distinguish

	BinaryString() string // right most is the least significant bit, // debug use
	Bytes() []byte        // binary representation, debug use
	Rand() PubNum
	From(int) PubNum
	Decode([]byte) PubNum
	// Operator
	Add(PubNum) PubNum
	Sub(PubNum) PubNum
	Mul(PubNum) PubNum
	Div(PubNum) PubNum

	Not() PubNum
	And(PubNum) PubNum
	Or(PubNum) PubNum
	Gt(PubNum) PubNum
	Lt(PubNum) PubNum
	Eq(PubNum) PubNum
	Shl(PubNum) PubNum
	Shr(PubNum) PubNum

	Mux(PubNum, PubNum) PubNum
}

func Zero(length int) PubNum {
	switch length {
	case 1:
		return ZeroBool
	case 8:
		return ZeroInt8
	case 16:
		return ZeroInt16
	case 32:
		return ZeroInt32
	case 64:
		return ZeroInt64
	default:
		log.Panicf("Query Zero with unsupported length, with value %d", length)
		return ZeroInt8
	}
}

// decode a little endian byte sequence of PubNum to PubNum
func DecodePubNum(b []byte) PubNum {
	l := len(b)
	if l == 1 {
		return Int8(b[0])
	} else if l == 2 {
		return Int16(int16(b[0]) | int16(b[1])<<8)
	} else if l == 4 {
		return Int32(int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24)
	} else if l == 8 {
		return Int64(int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56)
	} else {
		log.Panicln("The byte slice has a not expected length", l)
		return nil
	}
}
