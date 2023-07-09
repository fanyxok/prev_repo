package main

// import "s3l/mpcfgo/pkg/program"

// func matchv3(db1, db2, s1, s2 program.Sec[int32]) program.Sec[int32] {
// 	dist1 := Sub(db1, s1)
// 	dist2 := Sub(db2, s2)
// 	return Add(Mul(dist1, dist1), Mul(dist2, dist2))

// }
// func biometricv4(db program.Sec[[][]int32], sample program.Sec[[]int32]) program.Sec[int32] {
// 	db0 := Read(db, 0)
// 	min := matchv3(Read(db0, 0), Read(db0, 1), Read(sample, 0), Read(sample, 1))
// 	i := Public[int32](0)
// 	Loop(128)(func() {
// 		db_sample := Read(db, i)
// 		db1 := Read(db_sample, 0)
// 		db2 := Read(db_sample, 1)
// 		s1 := Read(sample, 0)
// 		s2 := Read(sample, 1)
// 		dist := matchv3(db1, db2, s1, s2)
// 		Update(&min, Mux(Lt(dist, min), dist, min))
// 		Update(&i, Addc(i, Public[int32](1)))
// 	})
// 	return min
// }
