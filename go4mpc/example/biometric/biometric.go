package main

const N = 128
const D = 2

// func match(db1, db2, s1, s2 program.Pvt[int32]) program.Pvt[int32] {
// 	dist1 := Sub(db1, s1)
// 	dist2 := Sub(db2, s2)
// 	return Add(Mul(dist1, dist1), Mul(dist2, dist2))

// }
// func biometric(db [N][D]program.Pvt[int32], sample [D]program.Pvt[int32]) program.Pvt[int32] {
// 	min := match(db[0][0], db[0][1], sample[0], sample[1])
// 	for i := 1; i < N; i++ {
// 		dist := match(db[i][0], db[i][1], sample[0], sample[1])
// 		Update(&min, Mux(Lt(dist, min), dist, min))
// 	}
// 	return min
// }

// func match_c(db1, db2, s1, s2 int32) int32 {
// 	dist1 := db1 - s1
// 	dist2 := db2 - s2
// 	return dist1*dist1 + dist2*dist2
// }
// func biometric_c(db [N][D]int32, sample [D]int32) int32 {
// 	min := match_c(db[0][0], db[0][1], sample[0], sample[1])
// 	for i := 1; i < N; i++ {
// 		dist := match_c(db[i][0], db[i][1], sample[0], sample[1])
// 		if dist < min {
// 			min = dist
// 		}
// 	}
// 	return min
// }
