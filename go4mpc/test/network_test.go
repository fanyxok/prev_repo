package test

import (
	"math/rand"
	"s3l/mpcfgo/internal/network"
	"testing"
	"time"
)

func TestNetServer(t *testing.T) {
	net := network.NewServer(":23334")
	net.Connect()
	bytes := make([]byte, 20)

	for i := 0; i < 1000000; i++ {
		rand.Read(bytes)
		msg := network.NewMsg(uint32(i), network.SimpleOT, bytes)
		net.Send(msg)
	}
	net.Recv()
}

func TestNetClient(t *testing.T) {
	net := network.NewClient("127.0.0.1:23334")
	net.Connect()
	for i := 0; i < 1000000; i++ {
		net.Recv()
	}
	net.Send(network.NewMsg(uint32(0), network.SimpleOT, nil))

}

func BenchmarkNet(b *testing.B) {
	ch := make(chan bool)
	b.ResetTimer()
	go func() {
		net := network.NewServer(":23334")
		net.Connect()
		bytes := make([]byte, 20)
		for i := 0; i < b.N; i++ {
			rand.Read(bytes)
			msg := network.NewMsg(uint32(i), network.SimpleOT, bytes)
			net.Send(msg)
		}
		ch <- true
	}()

	go func() {
		time.Sleep(time.Millisecond * 10)
		net := network.NewClient("127.0.0.1:23334")
		net.Connect()
		for i := 0; i < b.N; i++ {
			_ = net.Recv()
		}
		ch <- true
	}()

	<-ch
	<-ch
}
