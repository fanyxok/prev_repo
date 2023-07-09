package main

// import "s3l/mpcfgo/pkg/program"

// func rem(x program.Sec[int32], y program.Sec[int32]) program.Sec[int32] {
// 	rem := Secret[int32](0)
// 	Loop(32)(func() {
// 		Update(&rem, Shl(rem, Public(int8(1))))
// 		//shred := Shr(rem, Public(int8(k.Val)))
// 		rem2 := Sub(rem, y)
// 		b1 := Eq(rem, y)
// 		b2 := Gt(rem, y)
// 		cond := Or(b1, b2)
// 		If(cond)(func() {
// 			Update(&rem, rem2)
// 		})
// 	})
// 	return rem
// }
// func GCD(a program.Sec[int32], b program.Sec[int32]) program.Sec[int32] {
// 	gcd := Secret[int32](0)
// 	Loop(32)(func() {
// 		Update(&gcd, rem(a, b))
// 		temp := Clone(b)
// 		eq := Eq(b, Public[int32](0))
// 		cond := Not(eq)
// 		If(cond)(func() {
// 			Update(&b, gcd)
// 			Update(&a, temp)
// 		})
// 	})
// 	return gcd
// }
