package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type Network struct {
	Server  bool
	addr    string
	queue   chan []byte
	channel net.Conn
}

/**
Server impl
*/

func NewServer(addr string) Network {
	return Network{true, addr, make(chan []byte, 1024), nil}
}

func NewClient(addr string) Network {
	return Network{false, addr, make(chan []byte, 1024), nil}
}

func (ct *Network) Connect() {
	if ct.Server {
		ct.connect0()
	} else {
		ct.connect1()
	}
}
func (ct *Network) connect1() {
	// use ResolveTCPAddr to create address to connect to
	remoteIPv4, err := net.ResolveTCPAddr("tcp", ct.addr)
	letItCrash(err)
	conn, err := net.DialTCP("tcp", nil, remoteIPv4)
	letItCrash(err)
	ct.channel = conn
	go ct.run()
}

func (ct *Network) connect0() {
	// use ResolveTCPAddr to create address to connect to
	localIPv4, err := net.ResolveTCPAddr("tcp", ct.addr)
	letItCrash(err)
	l, err := net.ListenTCP("tcp", localIPv4)
	letItCrash(err)
	conn, err := l.Accept()
	letItCrash(err)
	fmt.Println("Incoming connection", conn.RemoteAddr())
	ct.channel = conn
	go ct.run()
}

func (ct *Network) Terminate() {
	ct.channel.Close()
}

func (ct *Network) run() {
	defer ct.channel.Close()
	for {
		header := make([]byte, 5)
		_, err := io.ReadFull(ct.channel, header)
		if err == io.EOF {
			continue
		}
		letItCrash(err)
		if header[0] != 'F' {
			panic("Network header is not 'F'" + string(header))
		}
		length := binary.LittleEndian.Uint32(header[1:5])
		msgBody := make([]byte, length)
		io.ReadFull(ct.channel, msgBody)
		ct.queue <- msgBody
	}
}

func (ct *Network) Send(data Msg) {
	_, err := ct.channel.Write(data.Encode())
	letItCrash(err)
}

func (ct *Network) Recv() Msg {
	encode := <-ct.queue
	msg := DecodeMsg(encode)
	return msg
}

func letItCrash(e error) {
	if e != nil {
		fmt.Println("let it crash:" + e.Error())
		panic(e)
	}
}
