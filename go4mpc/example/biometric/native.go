package main

import (
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/pvt"
)

func native_match(net network.Network, db1 pvt.PvtNum, db2, s1, s2 pvt.PvtNum) pvt.PvtNum {
	dist1 := db1.Sub(net, s1)
	dist2 := db2.Sub(net, s2)
	return dist1.Mul(net, dist1).Add(net, dist2.Mul(net, dist2))
}
func native_biometric(net network.Network, db [][]pub.Int32, sample []pub.Int32) pvt.PvtNum {
	s_db := make([][]pvt.PvtNum, 128)
	for i, v := range db {
		s_db[i] = []pvt.PvtNum{}
		for _, vj := range v {
			if net.Server {
				s_db[i] = append(s_db[i], pvt.YShare.New(net, vj))
			} else {
				s_db[i] = append(s_db[i], pvt.YShare.NewFrom(net))
			}
		}
	}
	s_sample := []pvt.PvtNum{}
	for i := range sample {
		if net.Server {
			s_sample = append(s_sample, pvt.YShare.NewFrom(net))
		} else {
			s_sample = append(s_sample, pvt.YShare.New(net, sample[i]))
		}
	}
	min := native_match(net, s_db[0][0], s_db[0][1], s_sample[0], s_sample[1])
	for i := 1; i < N; i++ {
		dist := native_match(net, s_db[i][0], s_db[i][1], s_sample[0], s_sample[1])
		min = dist.Lt(net, min).Mux(net, dist, min)
	}
	return min
}
