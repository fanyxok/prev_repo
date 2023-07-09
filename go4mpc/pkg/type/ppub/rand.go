package ppub

import "log"

func RandIntFrom(a Pub) Pub {
	switch a.Length() {
	case 8:
		return Int8(rand_.Intn(128))
	case 16:
		return Int16(rand_.Intn(16384))
	case 32:
		return Int32(rand_.Uint32())
	case 64:
		return Int64(rand_.Uint64())
	default:
		log.Panicln("Rand a length")
		return Bool(false)
	}
}

func RandOf[T Integer | Float | Boolean](a T) T {
	v, _ := Pub(a).Rand().(T)
	return v
}
func RandBool[T Boolean]() T {
	return rand_.Intn(2) != 0
}
