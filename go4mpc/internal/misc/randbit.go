package misc

import (
	"math/rand"
	"time"
)

type boolpool struct {
	src       rand.Source
	cache     int64
	remaining int
}

var bp = boolpool{src: rand.NewSource(time.Now().UnixNano())}

func Bool() bool {
	if bp.remaining == 0 {
		bp.cache, bp.remaining = bp.src.Int63(), 63
	}

	result := bp.cache&0x01 == 1
	bp.cache >>= 1
	bp.remaining--

	return result
}
