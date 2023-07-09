package triple

import (
	"log"
	"s3l/mpcfgo/pkg/type/pub"
)

type TriplesFactory struct {
	Bool  *Triples
	Int8  *Triples
	Int16 *Triples
	Int32 *Triples
	Int64 *Triples
}

func NewTripleFactory() *TriplesFactory {
	return new(TriplesFactory)
}

func (ct *TriplesFactory) SetTriples(t *Triples) {
	switch t.Zero.Length() {
	case 1:
		ct.Bool = t
	case 8:
		ct.Int8 = t
	case 16:
		ct.Int16 = t
	case 32:
		ct.Int32 = t
	case 64:
		ct.Int64 = t
	default:
		log.Panicf("The Length() of Triples are not supported, with value %d", t.Zero.Length())
	}
}
func (ct *TriplesFactory) NextTriple(length int) (pub.PubNum, pub.PubNum, pub.PubNum) {
	var a, b, c pub.PubNum
	switch length {
	case 1:
		a, b, c = ct.Bool.Next()
	case 8:
		a, b, c = ct.Int8.Next()
	case 16:
		a, b, c = ct.Int16.Next()
	case 32:
		a, b, c = ct.Int32.Next()
	case 64:
		a, b, c = ct.Int64.Next()
	default:
		log.Panicf("The Length() of Triples are not supported, with value %d", length)
	}
	return a, b, c
}
