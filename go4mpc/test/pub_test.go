package test

import (
	"s3l/mpcfgo/pkg/type/ppub"
	"testing"
)

func TestPub(t *testing.T) {
	a := ppub.Int8(12)
	c := ppub.RandOf(a)
	// +-*/
	var d = ppub.Add(a, c)
	d = ppub.Sub(a, c)
	d = ppub.Div(a, c)
	d = ppub.Mul(a, c)
	// >>, <<
	d = ppub.Shl(a, 1)
	d = ppub.Shl(d, 2)
	// >=,<=, >,<
	b := ppub.Gt(a, c)
	b = ppub.Lt(a, c)
	// ==
	b = ppub.Eq(a, c)
	b = ppub.Eq(b, b)
	// Mux
	d = ppub.Mux(b, a, c)

}
