package main

const N = 128
const D = 2

/*
	syntax requirement

1. 	Can decl const in global, one const one line
2.	Can decl var by 'var' in non-global scope, one var one line
3. 	Can update var's value by '=', once update one line
4.  Can use three address expr in ='s right side
5. 	Can use if use atomic condition, not normal exprssion
6. 	Can use for with one init var, one post
*/
func match(db0, db1, s0, s1 sint32) sint32 {
	var l0 sint32 = db0 - s0
	var l1 sint32 = db1 - s1
	var l02 sint32 = l0 * l0
	var l12 sint32 = l1 * l1
	var dis sint32 = l02 + l12
	return dis
}

func biometric(db [][]sint32, sample []sint32) sint32 {
	var min sint32 = match(db[0][0], db[0][1], sample[0], sample[1])
	for i := 1; i < N; i = i + 1 {
		var dist sint32 = match(db[i][0], db[i][1], sample[0], sample[1])
		var cond bool = dist < min
		if cond {
			min = dist
		}
	}
	return min
}

func main() {
	var in0 []sint32 = i32n(0, N*D)
	var db [][]sint32 = make([][]sint32, N)
	for i := 0; i < N; i = i + 1 {
		db[i] = make([]sint32, D)
		db[i][0] = in0[i*2]
		db[i][1] = in0[i*2+1]
	}
	var sample []sint32 = make([]sint32, D)
	sample[0] = i32(1)
	sample[1] = i32(1)
	biometric(db, sample)
}
