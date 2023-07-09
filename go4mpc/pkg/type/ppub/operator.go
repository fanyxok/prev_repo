package ppub

func Add[T Integer | Float](a, b T) T {
	return a + b
}
func Sub[T Integer | Float](a, b T) T {
	return a - b
}
func Div[T Integer | Float](a, b T) T {
	return a / b
}
func Mul[T Integer | Float](a, b T) T {
	return a * b
}
func Eq[T Integer | Float | Boolean, K Boolean](a, b T) K {
	return a == b
}
func Lt[T Integer | Float, K Boolean](a, b T) K {
	return a < b
}
func Gt[T Integer | Float, K Boolean](a, b T) K {
	return a > b
}
func And[T Boolean](a, b T) T {
	return a && b
}
func Or[T Boolean](a, b T) T {
	return a || b
}
func Not[T Boolean](a T) T {
	return !a
}
func Mux[B Boolean, T Integer | Float | Boolean](c B, a, b T) T {
	if c {
		return a
	} else {
		return b
	}
}
func Shr[T Integer](a T, b int) T {
	return a >> b
}
func Shl[T Integer](a T, b int) T {
	return a << b
}
