package lz

/*
*
	Lazy
*
*/

// Simplified Switch for init a value
type _case[T any] struct {
	_Cond  bool
	_Value T
}

func Case[T any](c bool, v T) _case[T] {
	return _case[T]{
		c,
		v,
	}
}
func CondInit[T any](defaul T, cases ..._case[T]) T {
	for _, case_ := range cases {
		if case_._Cond {
			return case_._Value
		}
	}
	return defaul
}
