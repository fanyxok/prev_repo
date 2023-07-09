package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/pkg/calib"
	"strings"
)

func main() {
	var N int
	var dirpath string
	var role int
	var addr string

	flag.IntVar(&N, "N", 0, "The number of Repeation")
	flag.StringVar(&dirpath, "O", "", "The file location to store the calibration result")
	flag.IntVar(&role, "Role", -1, "The Role of this party, 0 for server side, 1 for client side.")
	flag.StringVar(&addr, "Addr", "", "IP address of all party, formatted as 0.0.0.0:0000, splited by ';' ")
	flag.Parse()
	if N < 0 {
		panic(fmt.Sprintf("N should be positive, got %d\n", N))
	}
	if role != 0 && role != 1 {
		panic(fmt.Sprintf("Role should be 0 or 1, got %d\n", role))
	}
	if addr == "" {
		panic("IP address can't be ")
	}

	// in 2-party setting, client(party 1) connect to server(party 0), the party-0 address is the only need.
	// in multi-party setting, every party should connect to each other, so there will need other party's address
	addrs := strings.Split(addr, ";")
	for i := range addrs {
		fmt.Println("IP of ", i, addrs[i])
	}
	var net network.Network
	if role == 0 {
		net = network.NewServer(addrs[0])
	} else {
		net = network.NewClient(addrs[0])
	}
	net.Connect()
	cost := calib.Calibration(N, role, net)
	byt, err := json.Marshal(cost)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		if net.Server {
			os.WriteFile(filepath.Join(dirpath, "cost.json"), byt, 0666)
		} else {
			os.WriteFile(filepath.Join(dirpath, "cost1.json"), byt, 0666)
		}
	}

	opcost := calib.CalibrationOps(N, role, net)
	byt, err = json.Marshal(opcost)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		if net.Server {
			os.WriteFile(filepath.Join(dirpath, "opcost.json"), byt, 0666)
		} else {
			os.WriteFile(filepath.Join(dirpath, "opcost1.json"), byt, 0666)
		}
	}
}
