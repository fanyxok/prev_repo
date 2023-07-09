package main

// import (
// 	"flag"
// 	"fmt"
// 	"path/filepath"
// 	"s3l/mpcfgo/config"
// 	p "s3l/mpcfgo/pkg/program"
// 	"s3l/mpcfgo/pkg/static"
// )

// func Pd() {
// 	Set("GCD", 2)
// 	a := SecretOf[int32](0)
// 	b := SecretOf[int32](1)
// 	ret := GCD(a, b)
// 	p.Output(ret)
// 	{
// 		prog := p.Singleton()
// 		fmt.Printf("prog: %v\n", prog)
// 		//prog.EliminateOblivIf()
// 		//fmt.Printf("prog: %v\n", prog)

// 		//prog.ToSSA()
// 		//fmt.Printf("prog: %v\n", prog)

// 		prog.DumpToFile(filepath.Join(config.Root, "cmd", "gcd", "gcd.pd"))
// 		//pd := program.LoadProgram(filepath.Join(config.Root, "cmd", "gcd", "gcd.pd"))
// 		//json := static.LoadJson(filepath.Join(config.Root, "cmd", "calibration", "cost0.json"))
// 		//optimizer.Solve(pd, json, optimizer.TotalRunTime)
// 	}
// }
// func Main(PID int, RUN bool, arg int32) {
// 	if !RUN {
// 		if PID == 0 {
// 			Pd()
// 		}
// 	} else {
// 		//pd := program.LoadProgram(filepath.Join(config.Root, "cmd", "gcd", "gcd.pd"))
// 		runtype := static.LoadRunType(filepath.Join(config.Root, "cmd", "gcd", "runtype"))
// 		fmt.Println(runtype)
// 		// r := runner.NewRunner(&pd,
// 		// 	[]string{
// 		// 		"127.0.0.1:23344",
// 		// 		"127.0.0.1:25566",
// 		// 	},
// 		// 	PID,
// 		// 	[]pvt.PvtNum{
// 		// 		pvt.AShare,
// 		// 		pvt.YShare,
// 		// 	},
// 		// 	[]runner.Converter{pvt.A2Y, pvt.Y2A},
// 		// 	runtype,
// 		// )
// 		// r.Run([]any{arg})
// 		// o := r.Outputs()
// 		// for i, v := range o {
// 		// 	fmt.Println(i, v)
// 		// }
// 	}
// }
// func main() {
// 	var PID int
// 	flag.IntVar(&PID, "PID", -1, "Party ID of current party")
// 	var RUN bool
// 	flag.BoolVar(&RUN, "Run", false, "If to running in this execution")
// 	flag.Parse()
// 	fmt.Printf("PID %d\n", PID)
// 	fmt.Printf("Run %t\n", RUN)
// 	Main(PID, RUN, 0)
// }
