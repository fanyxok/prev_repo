package main

// import (
// 	"fmt"
// 	"path/filepath"
// 	"s3l/mpcfgo/config"
// 	"s3l/mpcfgo/pkg/program"
// 	p "s3l/mpcfgo/pkg/program"
// 	"s3l/mpcfgo/pkg/static"
// 	"s3l/mpcfgo/pkg/static/optimizer"
// )

// func Pd() {
// 	Set("biomatch", 2)
// 	in0 := SecretNOf[int32](0, 256)
// 	db := Secret(make([][]int32, 128))
// 	i := Public[int32](0)
// 	Loop(128)(func() {
// 		sample := Secret(make([]int32, 2))
// 		Write(sample, 0, Read(in0, Mulc(i, Public[int32](2))))
// 		Write(sample, 0, Read(in0, Addc(Mulc(i, Public[int32](2)), Public[int32](1))))
// 		Write(db, i, sample)
// 		Update(&i, Addc(i, Public[int32](1)))
// 	})
// 	sample := Secret(make([]int32, 2))
// 	Write(sample, 0, SecretOf[int32](1))
// 	Write(sample, 1, SecretOf[int32](1))
// 	min := biometricv4(db, sample)
// 	Output(min)
// 	fmt.Printf("p.Singleton(): %v\n", p.Singleton())
// }

// func IR() {
// 	prog := p.Singleton()
// 	prog.DumpToFile(filepath.Join(config.Root, "example", "biometric", "biometric_raw.pd"))
// 	//prog.EliminateOblivIf()
// 	//fmt.Printf("prog: %v\n", prog)
// 	prog.ToSSA()
// 	fmt.Printf("ssa prog: %v\n", prog)
// 	prog.DumpToFile(filepath.Join(config.Root, "example", "biometric", "biometric_ssa.pd"))
// 	pd := program.LoadProgram(filepath.Join(config.Root, "example", "biometric", "biometric_ssa.pd"))
// 	//fmt.Printf("pd: %v\n", pd)
// 	pd.DumpToFile(filepath.Join(config.Root, "example", "biometric", "biometric_load.pd"))

// }
// func Optimize() {
// 	json := static.LoadJson(filepath.Join(config.Root, "cmd", "calibration", "cost0.json"))
// 	pd := program.LoadProgram(filepath.Join(config.Root, "example", "biometric", "biometric_ssa.pd"))
// 	optimizer.Solve(pd, json, optimizer.TotalRunTime)
// }
// func Run(PID int, args_ []int32) {
// 	//pd := program.LoadProgram(filepath.Join(config.Root, "example", "biometric", "biometric.pd"))
// 	//runtype := static.LoadRunType(filepath.Join(config.Root, "example", "biometric", "runtype"))
// 	// r := runner.NewRunner(&pd,
// 	// 	[]string{
// 	// 		"127.0.0.1:23344",
// 	// 		"127.0.0.1:25566",
// 	// 	},
// 	// 	PID,
// 	// 	[]pvt.PvtNum{
// 	// 		pvt.AShare,
// 	// 		pvt.YShare,
// 	// 	},
// 	// 	[]runner.Converter{pvt.A2Y, pvt.Y2A},
// 	// 	runtype,
// 	// )
// 	// if PID == 0 {
// 	// 	args := make([]any, N*D)
// 	// 	for i := range args {
// 	// 		args[i] = args_[i]
// 	// 	}
// 	// 	r.Run(args)
// 	// } else {
// 	// 	args := make([]any, D)
// 	// 	for i := range args {
// 	// 		args[i] = args_[i]
// 	// 	}
// 	// 	r.Run(args)
// 	// }
// 	// o := r.Outputs()
// 	// for i, v := range o {
// 	// 	fmt.Println(i, v)
// 	// }
// }

// func main() {
// 	Pd()
// 	IR()
// }

// // func main() {
// // 	var PID int
// // 	flag.IntVar(&PID, "PID", -1, "Party ID of current party")
// // 	var RUN bool
// // 	flag.BoolVar(&RUN, "Run", false, "If to running in this execution")
// // 	flag.Parse()
// // 	fmt.Printf("PID %d\n", PID)
// // 	fmt.Printf("Run %t\n", RUN)
// // 	Main(PID, RUN, 0)
// // }
