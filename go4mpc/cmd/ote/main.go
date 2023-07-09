package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/internal/ot/ote"
	"time"
)

func main() {
	var r int
	flag.IntVar(&r, "i", 0, "role")
	flag.Parse()
	bytes := 1

	if r == 0 {
		x0 := make([]byte, 1000000*bytes)
		x1 := make([]byte, 1000000*bytes)
		serNet := network.NewServer(":22334")
		serNet.Connect()

		rand.Read(x0)
		rand.Read(x1)
		//baseot.Send128(ser, x0, x1)
		//ote.SendN(ser, x0, x1)
		t := time.Now()
		ote.InitOtSender(serNet)
		ote.OTE.SendN(serNet, x0, x1, 1000000, bytes)
		fmt.Printf("Time: %v\n", time.Since(t))
	} else if r == 1 {
		c := make([]bool, 1000000)
		for i := 0; i < 1000000; i++ {
			c[i] = misc.Bool()
		}
		dst := make([]byte, 1000000)
		serNet := network.NewClient("127.0.0.1:22334")
		serNet.Connect()
		t := time.Now()
		ote.InitOtReceiver(serNet)
		ote.OTE.RecvN(serNet, dst, c, 1000000, bytes)
		fmt.Printf("Time: %v\n", time.Since(t))
	}

}
