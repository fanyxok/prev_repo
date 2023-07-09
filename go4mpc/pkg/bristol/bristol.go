package bristol

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"s3l/mpcfgo/config"
	"strconv"
	"strings"
)

type GateType byte

const (
	AND GateType = iota
	XOR
	NOT
	EQ
	EQW
)

func (ct GateType) String() (str string) {
	switch ct {
	case AND:
		str = "AND"
	case XOR:
		str = "XOR"
	case NOT:
		str = "INV"
	case EQ:
		str = "EQ"
	case EQW:
		str = "EQW"
	}
	return
}

type BristolGate struct {
	Type    GateType
	InWire  []int
	OutWire int
}

type BristolCircuit struct {
	Gate      []BristolGate
	Wire      int
	InLength  []int
	OutLength int

	ANDs int
	XORs int
	NOTs int
	EQs  int
	EQWs int
}

var (
	ADDi64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Add-64-(64,64).txt"))
	ADDi32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Add-32-(32,32).txt"))
	ADDi16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Add-16-(16,16).txt"))
	ADDi8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Add-8-(8,8).txt"))

	SUBi64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Sub-64-(64,64).txt"))
	SUBi32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Sub-32-(32,32).txt"))
	SUBi16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Sub-16-(16,16).txt"))
	SUBi8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Sub-8-(8,8).txt"))

	MULi64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mul-64-(64,64).txt"))
	MULi32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mul-32-(32,32).txt"))
	MULi16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mul-16-(16,16).txt"))
	MULi8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mul-8-(8,8).txt"))

	// msb doesn't
	DIVi64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Div-64-(64,64).txt"))
	DIVi32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Div-32-(32,32).txt"))
	DIVi16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Div-16-(16,16).txt"))
	DIVi8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Div-8-(8,8).txt"))

	NOTb1 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Not-1-(1).txt"))
	ANDb1 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-And-1-(1,1).txt"))
	ORb1  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Or-1-(1,1).txt"))
	// // msb is as value
	// UDIVi64 = FromFile(filepath.Join(config.Root, "pkg", "bristol", "UDIV64x64-64.txt"))
	// UDIVi32 BristolCircuit
	// UDIVi16 BristolCircuit
	// UDIVi8  BristolCircuit
	EQb1  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Eq-1-(1,1).txt"))
	EQi8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Eq-1-(8,8).txt"))
	EQi16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Eq-1-(16,16).txt"))
	EQi32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Eq-1-(32,32).txt"))
	EQi64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Eq-1-(64,64).txt"))

	Gti8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Gt-1-(8,8).txt"))
	Gti16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Gt-1-(16,16).txt"))
	Gti32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Gt-1-(32,32).txt"))
	Gti64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Gt-1-(64,64).txt"))

	Lti8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Lt-1-(8,8).txt"))
	Lti16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Lt-1-(16,16).txt"))
	Lti32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Lt-1-(32,32).txt"))
	Lti64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Lt-1-(64,64).txt"))

	Muxb1  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mux-1-(1,1,1).txt"))
	Muxi8  = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mux-8-(1,8,8).txt"))
	Muxi16 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mux-16-(1,16,16).txt"))
	Muxi32 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mux-32-(1,32,32).txt"))
	Muxi64 = FromFile(filepath.Join(config.Root, "pkg", "cbmc-gc", "GC-Mux-64-(1,64,64).txt"))

	NEGi64 = FromFile(filepath.Join(config.Root, "pkg", "bristol", "NEG64-64.txt"))
	NEGi32 = FromFile(filepath.Join(config.Root, "pkg", "bristol", "NEG64-64.txt"))
	NEGi16 = FromFile(filepath.Join(config.Root, "pkg", "bristol", "NEG64-64.txt"))
	NEGi8  = FromFile(filepath.Join(config.Root, "pkg", "bristol", "NEG64-64.txt"))

	AESi128 = FromFile(filepath.Join(config.Root, "pkg", "bristol", "AES128x128-128.txt"))
	ZEROi64 = FromFile(filepath.Join(config.Root, "pkg", "bristol", "ZERO64-1.txt"))
)

var (
	AddCirc = map[int]BristolCircuit{8: ADDi8, 16: ADDi16, 32: ADDi32, 64: ADDi64}
	SubCirc = map[int]BristolCircuit{8: SUBi8, 16: SUBi16, 32: SUBi32, 64: SUBi64}
	MulCirc = map[int]BristolCircuit{8: MULi8, 16: MULi16, 32: MULi32, 64: MULi64}
	DivCirc = map[int]BristolCircuit{8: DIVi8, 16: DIVi16, 32: DIVi32, 64: DIVi64}
	EqCirc  = map[int]BristolCircuit{1: EQb1, 8: EQi8, 16: EQi16, 32: EQi32, 64: EQi64}
	GtCirc  = map[int]BristolCircuit{8: Gti8, 16: Gti16, 32: Gti32, 64: Gti64}
	LtCirc  = map[int]BristolCircuit{8: Lti8, 16: Lti16, 32: Lti32, 64: Lti64}
	MuxCirc = map[int]BristolCircuit{1: Muxb1, 8: Muxi8, 16: Muxi16, 32: Muxi32, 64: Muxi64}
)

func FromFile(file string) BristolCircuit {

	// log.Println(config.Root)

	f, err := os.Open(file)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer f.Close()
	bc := BristolCircuit{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	numOfG, _ := strconv.Atoi(scanner.Text())

	bc.Gate = make([]BristolGate, numOfG)
	scanner.Scan()
	bc.Wire, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	numOfIns, _ := strconv.Atoi(scanner.Text())
	bc.InLength = make([]int, numOfIns)
	for i := 0; i < numOfIns; i++ {
		scanner.Scan()
		bc.InLength[i], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	numOfOuts, _ := strconv.Atoi(scanner.Text())
	if numOfOuts != 1 {
		log.Panicln("")
	}
	for i := 0; i < numOfOuts; i++ {
		scanner.Scan()
		bc.OutLength, _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < len(bc.Gate); i++ {
		g := BristolGate{}
		scanner.Scan()
		inWnum, _ := strconv.Atoi(scanner.Text())
		g.InWire = make([]int, inWnum)
		scanner.Scan()
		strconv.Atoi(scanner.Text())

		for j := 0; j < len(g.InWire); j++ {
			scanner.Scan()
			g.InWire[j], _ = strconv.Atoi(scanner.Text())
		}
		scanner.Scan()
		g.OutWire, _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		gType := scanner.Text()
		switch gType {
		case "AND":
			g.Type = AND
		case "XOR":
			g.Type = XOR
		case "INV":
			g.Type = NOT
		case "EQ":
			g.Type = EQ
		case "EQW":
			g.Type = EQW
		case "MAND":
			log.Panicln("Can't accept MAND gate")
		}
		bc.Gate[i] = g
	}
	return bc
}

func (ct BristolCircuit) DumpToFile(filePath string) {
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			panic(err.Error())
		}
	} else if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%d %d\n", len(ct.Gate), ct.Wire))

	file.WriteString(strconv.Itoa(len(ct.InLength)) + " ")
	for _, v := range ct.InLength {
		file.WriteString(strconv.Itoa(v) + " ")
	}
	file.WriteString("\n")
	file.WriteString(fmt.Sprintf("1 %d", ct.OutLength))
	file.WriteString("\n")
	for _, v := range ct.Gate {
		var inwires []string
		for _, jv := range v.InWire {
			inwires = append(inwires, strconv.Itoa(jv))
		}
		inwire := strings.Join(inwires, " ")
		file.WriteString(fmt.Sprintf("%d 1 %s %d %s\n", len(v.InWire), inwire, v.OutWire, v.Type))
	}
}
