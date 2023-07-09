package functional

import (
	"strconv"
	"time"
)

func Time2MicroSecondFloat64(t time.Duration) float64 {
	t = t / time.Microsecond
	return float64(t)
}
func Float642String(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}
