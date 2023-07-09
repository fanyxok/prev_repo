package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"s3l/mpcfgo/pkg/bristol"
	"strconv"
)

func loadOldBristolFromFile(f *os.File) bristol.BristolCircuit {
	bc := bristol.BristolCircuit{}
	scanner := bufio.NewScanner(f)
	// 1 line, numberOfGate numberOfWire
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	numOfG, _ := strconv.Atoi(scanner.Text())
	bc.Gate = make([]bristol.BristolGate, numOfG)
	scanner.Scan()
	bc.Wire, _ = strconv.Atoi(scanner.Text())
	// 1 line, numberOfWireOfInput0 numberOfWireOfInput0 numberOfWireOfOutput
	scanner.Scan()
	numOfWireOfInput0, _ := strconv.Atoi(scanner.Text())
	bc.InLength = append(bc.InLength, numOfWireOfInput0)
	scanner.Scan()
	numOfWireOfInput1, _ := strconv.Atoi(scanner.Text())
	if numOfWireOfInput1 != 0 {
		bc.InLength = append(bc.InLength, numOfWireOfInput1)
	}
	scanner.Scan()
	bc.OutLength, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < len(bc.Gate); i++ {
		g := bristol.BristolGate{}
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
			g.Type = bristol.AND
		case "XOR":
			g.Type = bristol.XOR
		case "INV":
			g.Type = bristol.NOT
		case "EQ":
			g.Type = bristol.EQ
		case "EQW":
			g.Type = bristol.EQW
		case "MAND":
			log.Panicln("Can't accept MAND gate")
		}
		bc.Gate[i] = g
	}
	return bc
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		panic("args more than 1")
	}
	filePath := args[0]

	{
		f, err := os.Open(filePath)
		if err != nil {
			panic(err.Error())
		}
		loadOldBristolFromFile(f).DumpToFile(filePath + ".new")
		defer f.Close()
	}
}
