package misc

func GCD(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	if a%b == 0 {
		return b
	}
	return GCD(b, a%b)
}
