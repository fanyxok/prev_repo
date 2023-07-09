package ppub

import (
	"math/rand"
	"time"
)

var rand_ = rand.New(rand.NewSource(time.Now().UnixNano()))

type (
	Integer interface {
		Int8 | Int16 | Int32 | Int64
	}
	Boolean interface {
		Bool
	}
	Float interface {
		Float64 | Float32
	}

	Pub interface {
		Bytes() []byte // binary representation, debug use
		Length() int
		From(int) Pub
		Rand() Pub
		Decode([]byte) Pub
	}
	// Int8
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Bool    bool
	Float32 float32
	Float64 float64
)

const (
	ZeroInt8  = Int8(0)
	ZeroInt16 = Int16(0)
	ZeroInt32 = Int32(0)
	ZeroInt64 = Int64(0)
	ZeroBool  = Bool(false)
)

func (ct Int8) Rand() Pub {
	return Int8(rand_.Intn(128))
}

func (ct Int8) Length() int {
	return 8
}
func (ct Int8) Bytes() []byte {
	return []byte{byte(ct)}
}
func (ct Int8) From(x int) Pub {
	return Int8(x)
}
func (ct Int8) Decode(bytes []byte) Pub {
	return Int8(bytes[0])
}

/*
Int16
*/

func (ct Int16) Rand() Pub {
	return Int16(rand_.Intn(16384))
}
func (ct Int16) Length() int {
	return 16
}
func (ct Int16) Bytes() []byte {
	return append([]byte{}, byte(ct), byte(ct>>8))
}
func (ct Int16) From(x int) Pub {
	return Int16(x)
}
func (ct Int16) Decode(b []byte) Pub {
	_ = b[1]
	return Int16(int16(b[0]) | int16(b[1])<<8)
}

// Int32

func (ct Int32) Rand() Pub {
	return Int32(rand_.Uint32())
}
func (ct Int32) Length() int {
	return 32
}
func (ct Int32) Bytes() []byte {
	return append([]byte{}, byte(ct), byte(ct>>8), byte(ct>>16), byte(ct>>24))
}
func (ct Int32) From(x int) Pub {
	return Int32(x)
}
func (ct Int32) Decode(b []byte) Pub {
	_ = b[3]
	return Int32(int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24)
}

// Int64

func (ct Int64) Rand() Pub {
	return Int64(rand_.Uint64())
}
func (ct Int64) Length() int {
	return 64
}
func (ct Int64) Bytes() []byte {
	return append([]byte{}, byte(ct), byte(ct>>8), byte(ct>>16), byte(ct>>24), byte(ct>>32), byte(ct>>40), byte(ct>>48), byte(ct>>56))
}
func (ct Int64) From(x int) Pub {
	return Int64(x)
}
func (ct Int64) Decode(b []byte) Pub {
	_ = b[7]
	return Int64(int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56)
}

// Bool

func (ct Bool) Rand() Pub {
	return Bool(rand_.Intn(2) != 0)
}
func (ct Bool) Length() int {
	return 1
}
func (ct Bool) Bytes() []byte {
	if ct {
		return []byte{1}
	} else {
		return []byte{0}
	}
}
func (ct Bool) From(x int) Pub {
	return Bool(x != 0)
}
func (ct Bool) Decode(bytes []byte) Pub {
	return Bool(bytes[0] != 0)
}

// float
func (ct Float32) Rand() Pub {
	return Bool(rand_.Intn(2) != 0)
}
func (ct Float32) Length() int {
	return 1
}
func (ct Float32) Bytes() []byte {
	return nil
}
func (ct Float32) From(x int) Pub {
	return Bool(x != 0)
}
func (ct Float32) Decode(bytes []byte) Pub {
	return Bool(bytes[0] != 0)
}

func (ct Float64) Rand() Pub {
	return Bool(rand_.Intn(2) != 0)
}
func (ct Float64) Length() int {
	return 1
}
func (ct Float64) Bytes() []byte {
	return nil
}
func (ct Float64) From(x int) Pub {
	return Bool(x != 0)
}
func (ct Float64) Decode(bytes []byte) Pub {
	return Bool(bytes[0] != 0)
}
