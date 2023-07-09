package main

// import (
// 	"fmt"
// 	"log"
// 	"s3l/mpcfgo/internal/network"
// 	"s3l/mpcfgo/pkg/type/pub"
// 	"testing"
// 	"time"
// )

// func Nets() (network.Network, network.Network) {
// 	ch := make(chan bool)
// 	serNet := network.NewServer(":22334")
// 	go func() {
// 		serNet.Connect()
// 		ch <- true
// 	}()

// 	cliNet := network.NewClient(":22334")
// 	go func() {
// 		time.Sleep(time.Millisecond * 100)
// 		cliNet.Connect()
// 		ch <- true
// 	}()
// 	<-ch
// 	<-ch
// 	log.Println("Server Net:", serNet, "\nClient Net:", cliNet)
// 	return serNet, cliNet
// }

// func Parallel(f0 func(), f1 func()) {
// 	ch := make(chan bool)
// 	go func() {
// 		f0()
// 		ch <- true
// 	}()
// 	go func() {
// 		f1()
// 		ch <- true
// 	}()
// 	<-ch
// 	<-ch
// }

// func TestPd(t *testing.T) {
// 	Pd()
// }
// func TestIR(t *testing.T) {
// 	Pd()
// 	IR()
// }
// func TestOptimize(t *testing.T) {
// 	Pd()
// 	IR()
// 	Optimize()
// }
// func TestMain(t *testing.T) {
// 	main()
// }

// func TestRun0(t *testing.T) {
// 	t.Parallel()
// 	db := make([]int32, N*D)
// 	for i := range db {
// 		db[i] = int32(i)
// 	}
// 	fmt.Print(db)
// 	Run(0, db)
// }
// func TestRun1(t *testing.T) {
// 	t.Parallel()
// 	db := []int32{300, 300}
// 	fmt.Print(db)
// 	Run(1, db)
// }

// func TestNative(t *testing.T) {
// 	ser, cli := Nets()
// 	N := 128
// 	DB := make([][]pub.Int32, N)
// 	for i := 0; i < N; i++ {
// 		DB[i] = []pub.Int32{pub.ZeroInt32.Rand().(pub.Int32), pub.ZeroInt32.Rand().(pub.Int32)}
// 	}
// 	Sample := []pub.Int32{pub.ZeroInt32.Rand().(pub.Int32), pub.ZeroInt32.Rand().(pub.Int32)}

// 	Parallel(func() {
// 		native_biometric(ser, DB, Sample)
// 	}, func() {
// 		native_biometric(cli, DB, Sample)

// 	})
// }
