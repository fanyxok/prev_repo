package test

import (
	"log"
	"s3l/mpcfgo/internal/network"
	"time"
)

const Sample = 20

func Nets() (network.Network, network.Network) {
	ch := make(chan bool)
	serNet := network.NewServer(":22334")
	go func() {
		serNet.Connect()
		ch <- true
	}()

	cliNet := network.NewClient(":22334")
	go func() {
		time.Sleep(time.Millisecond * 100)
		cliNet.Connect()
		ch <- true
	}()
	<-ch
	<-ch
	log.Println("Server Net:", serNet, "\nClient Net:", cliNet)
	return serNet, cliNet
}

func Parallel(f0 func(), f1 func()) {
	ch := make(chan bool)
	go func() {
		f0()
		ch <- true
	}()
	go func() {
		f1()
		ch <- true
	}()
	<-ch
	<-ch
}
