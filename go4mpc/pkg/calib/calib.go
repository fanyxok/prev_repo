package calib

import (
	"log"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/pkg/functional"
	"s3l/mpcfgo/pkg/primitive/triple"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/pvt"
	"time"
)

var types = []pub.PubNum{pub.ZeroBool, pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
var itypes = []pub.PubNum{pub.ZeroInt8, pub.ZeroInt16, pub.ZeroInt32, pub.ZeroInt64}
var proInst = []pvt.PvtNum{pvt.AShare, pvt.YShare}
var proName = []string{"A", "Y"}

func formatValueType(s bool, p pub.PubNum) string {
	if s {
		switch p.(type) {
		case pub.Bool:
			return "sb1"
		case pub.Int8:
			return "si8"
		case pub.Int16:
			return "si16"
		case pub.Int32:
			return "si32"
		case pub.Int64:
			return "si64"
		default:
			log.Panicf("")
		}
	} else {
		switch p.(type) {
		case pub.Bool:
			return "b1"
		case pub.Int8:
			return "i8"
		case pub.Int16:
			return "i16"
		case pub.Int32:
			return "i32"
		case pub.Int64:
			return "i64"
		default:
			log.Panicf("")
		}
	}
	return " "
}

var nonZero = func(pn pub.PubNum) bool {
	if pn.Eq(pn.From(0)).(pub.Bool) {
		return false
	} else {
		return true
	}
}
var noCond = func(pn pub.PubNum) bool {
	return true
}
var positive = func(pn pub.PubNum) bool {
	if pn.Lt(pn.From(0)).(pub.Bool) {
		return false
	} else {
		return true
	}
}

func initOpCostJson(name string) OpCostJson {
	r := OpCostJson{}
	r.OpName = name
	r.OpCost = make(map[string]Protocols)
	return r
}

func initProtocols() Protocols {
	r := Protocols{}
	r.Cost = make(map[string]float64)
	return r
}

func Twice(net network.Network, pctype pvt.PvtNum, pn pub.PubNum, cond func(pn pub.PubNum) bool, N int) ([]pvt.PvtNum, []pvt.PvtNum) {
	X := genRandomPvt(net, pctype, pn, cond, N)
	Y := genRandomPvt(net, pctype, pn, cond, N)
	return X, Y
}

var t time.Time

func Time(x ...int) float64 {
	if len(x) == 0 {
		t = time.Now()
		return 0
	}
	elapsed := time.Since(t) / time.Duration(x[0])
	return functional.Time2MicroSecondFloat64(elapsed)
}

// "Add",
// "Sub",
// "Div",
// "Mul",

// "Eq",
// "Gt",
// "Lt",

// "And",
// "Or",
// "Not",

// "Mux",

// "Shr",
// "Shl",

// "New",

func CalibrationCoreNew(net network.Network, N int) OpCostJson {
	log.Println("New...")
	cost := initOpCostJson("New")
	for i := range types {
		p := initProtocols()
		for j := range proInst {
			if net.Server {
				X := genRandomPub(types[i], func(pn pub.PubNum) bool { return true }, N)
				Time()
				for k := range X {
					proInst[j].New(net, X[k])
				}
				p.Cost[proName[j]] = Time(N)
			} else {
				Time()
				for k := 0; k < N; k++ {
					proInst[j].NewFrom(net)
				}
				p.Cost[proName[j]] = Time(N)
			}
		}
		b := formatValueType(true, types[i])
		cost.OpCost[b] = p
	}
	return cost
}

// A, Y
func CalibrationCoreAdd(net network.Network, N int) OpCostJson {
	log.Println("Add...")
	cost := initOpCostJson("Add")
	for i := range itypes {
		p := initProtocols()
		for j := range proInst {
			X, Y := Twice(net, proInst[j], itypes[i], noCond, N)
			Time()
			for i := range X {
				_ = X[i].Add(net, Y[i])
			}
			p.Cost[proName[j]] = Time(N)
		}
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_"+b] = p
	}
	return cost
}
func CalibrationCoreSub(net network.Network, N int) OpCostJson {
	log.Println("Sub...")
	cost := initOpCostJson("Sub")
	for i := range itypes {
		p := initProtocols()
		for j := range proInst {
			X, Y := Twice(net, proInst[j], itypes[i], noCond, N)
			Time()
			for i := range X {
				_ = X[i].Sub(net, Y[i])
			}
			p.Cost[proName[j]] = Time(N)
		}
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_"+b] = p
	}
	return cost
}
func CalibrationCoreMul(net network.Network, N int) OpCostJson {
	log.Println("Mul...")
	cost := initOpCostJson("Mul")
	for i := range itypes {
		p := initProtocols()
		for j := range proInst {
			X, Y := Twice(net, proInst[j], itypes[i], noCond, N)
			if proInst[j] == pvt.AShare {
				pvt.TripleFactory(net.Server).SetTriples(triple.NewTriples(net, itypes[i], N))
			}
			Time()
			for i := range X {
				_ = X[i].Mul(net, Y[i])
			}
			p.Cost[proName[j]] = Time(N)
		}
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_"+b] = p
	}
	return cost
}
func CalibrationCoreDiv(net network.Network, N int) OpCostJson {
	log.Println("Div...")
	cost := initOpCostJson("Div")
	for i := range itypes {
		p := initProtocols()
		X, Y := Twice(net, pvt.YShare, itypes[i], noCond, N)
		Time()
		for i := range X {
			_ = X[i].Div(net, Y[i])
		}
		p.Cost["Y"] = Time(N)
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_"+b] = p
	}
	return cost
}
func CalibrationCoreAnd(net network.Network, N int) OpCostJson {
	log.Println("And...")
	cost := initOpCostJson("And")
	p := initProtocols()
	for j := range proInst {
		X, Y := Twice(net, proInst[j], pub.ZeroBool, noCond, N)
		if proInst[j] == pvt.AShare {
			pvt.TripleFactory(net.Server).SetTriples(triple.NewTriples(net, pub.ZeroBool, N))
		}
		Time()
		for i := range X {
			_ = X[i].And(net, Y[i])
		}
		p.Cost[proName[j]] = Time(N)
	}
	cost.OpCost["sb1_sb1"] = p
	return cost
}
func CalibrationCoreOr(net network.Network, N int) OpCostJson {
	log.Println("Or...")
	cost := initOpCostJson("Or")
	p := initProtocols()
	for j := range proInst {
		X, Y := Twice(net, proInst[j], pub.ZeroBool, noCond, N)
		if proInst[j] == pvt.AShare {
			pvt.TripleFactory(net.Server).SetTriples(triple.NewTriples(net, pub.ZeroBool, N))
		}
		Time()
		for i := range X {
			_ = X[i].Or(net, Y[i])
		}
		p.Cost[proName[j]] = Time(N)
	}
	cost.OpCost["sb1_sb1"] = p
	return cost
}

func CalibrationCoreNot(net network.Network, N int) OpCostJson {
	log.Println("Not...")

	cost := initOpCostJson("Not")
	p := initProtocols()
	for j := range proInst {
		X, _ := Twice(net, proInst[j], pub.ZeroBool, noCond, N)
		Time()
		for i := range X {
			_ = X[i].Not(net)
		}
		p.Cost[proName[j]] = Time(N)
	}
	cost.OpCost["sb1"] = p
	return cost
}
func CalibrationCoreEq(net network.Network, N int) OpCostJson {
	log.Println("Eq...")
	cost := initOpCostJson("Eq")
	pvt.TripleFactory(net.Server).SetTriples(triple.NewTriples(net, pub.ZeroBool, 116*N))
	for i := range types {
		p := initProtocols()
		for j := range proInst {
			X, Y := Twice(net, proInst[j], types[i], noCond, N)
			Time()
			for i := range X {
				_ = X[i].Eq(net, Y[i])
			}
			p.Cost[proName[j]] = Time(N)
		}
		b := formatValueType(true, types[i])
		cost.OpCost[b+"_"+b] = p
	}
	return cost
}

func CalibrationCoreGt(net network.Network, N int) OpCostJson {
	log.Println("Gt...")

	cost := initOpCostJson("Gt")
	for i := range itypes {
		p := initProtocols()
		X, Y := Twice(net, pvt.YShare, itypes[i], noCond, N)
		Time()
		for i := range X {
			_ = X[i].Gt(net, Y[i])
		}
		p.Cost["Y"] = Time(N)
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_"+b] = p
	}
	return cost
}

func CalibrationCoreLt(net network.Network, N int) OpCostJson {
	log.Println("Lt...")

	cost := initOpCostJson("Lt")
	for i := range itypes {
		p := initProtocols()
		X, Y := Twice(net, pvt.YShare, itypes[i], noCond, N)
		Time()
		for i := range X {
			_ = X[i].Lt(net, Y[i])
		}
		p.Cost["Y"] = Time(N)
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_"+b] = p
	}
	return cost
}

func CalibrationCoreShr(net network.Network, N int) OpCostJson {
	log.Println("Shr...")

	cost := initOpCostJson("Shr")
	for i := range itypes {
		p := initProtocols()
		for j := range proInst {
			X, _ := Twice(net, proInst[j], itypes[i], noCond, N)
			Y := genRandomPub(pub.ZeroInt8, positive, N)
			Time()
			for i := range X {
				_ = X[i].Shr(net, Y[i])
			}
			p.Cost[proName[j]] = Time(N)
		}
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_i8"] = p
	}
	return cost
}

func CalibrationCoreShl(net network.Network, N int) OpCostJson {
	log.Println("Shl...")

	cost := initOpCostJson("Shl")
	for i := range itypes {
		p := initProtocols()
		for j := range proInst {
			X, _ := Twice(net, proInst[j], itypes[i], noCond, N)
			Y := genRandomPub(pub.ZeroInt8, positive, N)
			Time()
			for i := range X {
				_ = X[i].Shl(net, Y[i])
			}
			p.Cost[proName[j]] = Time(N)
		}
		b := formatValueType(true, itypes[i])
		cost.OpCost[b+"_i8"] = p
	}
	return cost
}

func CalibrationCoreMux(net network.Network, N int) OpCostJson {
	log.Println("Mux...")

	cost := initOpCostJson("Mux")
	for i := range types {
		p := initProtocols()
		for j := range proInst {
			X, Y := Twice(net, proInst[j], types[i], noCond, N)
			B := genRandomPvt(net, proInst[j], pub.ZeroBool, noCond, N)
			if proInst[j] == pvt.AShare {
				pvt.TripleFactory(net.Server).SetTriples(triple.NewTriples(net, types[i], 3*N))
			}
			Time()
			for i := range X {
				_ = B[i].Mux(net, X[i], Y[i])
			}
			p.Cost[proName[j]] = Time(N)
		}
		b := formatValueType(true, types[i])
		cost.OpCost["sb1_"+b+"_"+b] = p
	}
	return cost
}
func CalibrationOps(N int, role int, net network.Network) []OpCostJson {
	log.SetPrefix("[CalibrationOps]")
	opcost := []OpCostJson{}
	opcost = append(opcost, CalibrationCoreNew(net, N))
	opcost = append(opcost, CalibrationCoreAdd(net, N))
	opcost = append(opcost, CalibrationCoreSub(net, N))
	opcost = append(opcost, CalibrationCoreMul(net, N))
	opcost = append(opcost, CalibrationCoreDiv(net, N))
	opcost = append(opcost, CalibrationCoreNot(net, N))
	opcost = append(opcost, CalibrationCoreAnd(net, N))
	opcost = append(opcost, CalibrationCoreOr(net, N))
	opcost = append(opcost, CalibrationCoreEq(net, N))
	opcost = append(opcost, CalibrationCoreGt(net, N))
	opcost = append(opcost, CalibrationCoreLt(net, N))
	opcost = append(opcost, CalibrationCoreShr(net, N))
	opcost = append(opcost, CalibrationCoreShl(net, N))
	opcost = append(opcost, CalibrationCoreMux(net, N))
	log.SetPrefix("")
	return opcost
}
func CalibrationCoreY2A(net network.Network, N int) CostJson {
	log.Println("Y2A...")
	cost := CostJson{}
	cost.Name = "Y2A"
	cost.Cost = map[string]float64{}
	for i := range types {
		X := genRandomPvt(net, pvt.YShare, types[i], noCond, N)
		Time()
		for i := range X {
			_ = pvt.Y2A(net, X[i])
		}
		b := formatValueType(true, types[i])
		cost.Cost[b] = Time(N)
	}
	return cost
}
func CalibrationCoreA2Y(net network.Network, N int) CostJson {
	log.Println("A2Y...")
	cost := CostJson{}
	cost.Name = "A2Y"
	cost.Cost = map[string]float64{}
	for i := range types {
		X := genRandomPvt(net, pvt.AShare, types[i], noCond, N)
		Time()
		for i := range X {
			_ = pvt.A2Y(net, X[i])
		}
		b := formatValueType(true, types[i])
		cost.Cost[b] = Time(N)
	}
	return cost
}
func Calibration(N int, role int, net network.Network) []CostJson {
	log.SetPrefix("[Calibration]")
	opcost := []CostJson{}
	opcost = append(opcost, CalibrationCoreA2Y(net, N))
	opcost = append(opcost, CalibrationCoreY2A(net, N))
	log.SetPrefix("")
	return opcost

}

func genRandomPvt(net network.Network, proto pvt.PvtNum, typ pub.PubNum, cond func(pub.PubNum) bool, N int) []pvt.PvtNum {
	nums := make([]pvt.PvtNum, N)
	if net.Server {
		for i := 0; i < N; i++ {
			v := typ.Rand()
			for !cond(v) {
				v = typ.Rand()
			}
			nums[i] = proto.New(net, v)
		}
	} else {
		for i := 0; i < N; i++ {
			nums[i] = proto.NewFrom(net)
		}
	}
	return nums
}

func genRandomPub(typ pub.PubNum, cond func(pub.PubNum) bool, N int) []pub.PubNum {
	nums := make([]pub.PubNum, N)
	for i := 0; i < N; i++ {
		v := typ.Rand()
		for !cond(v) {
			v = typ.Rand()
		}
		nums[i] = v
	}
	return nums
}
