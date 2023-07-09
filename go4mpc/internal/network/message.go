package network

import (
	"encoding/binary"
)

type MsgType byte

type Msg struct {
	SN   uint32  // 4 bytes
	Type MsgType // 1 byte
	Data []byte  // many bytes
}

const (
	SimpleOT MsgType = iota
	ROT              // random OT
	OTe              // OT extension
	COT              // correlated OT
	Sharing          // Sharing , UnSharing
	Triple           // Triple generation
	Yshare           // Yao's Sharing
)

func NewMsg(sn uint32, mtype MsgType, data []byte) Msg {
	return Msg{sn, mtype, data}
}

func DecodeMsg(encode []byte) Msg {
	return Msg{
		binary.LittleEndian.Uint32(encode[0:4]),
		MsgType(encode[4]),
		encode[5:]}
}

/*
*
F + uint32 + (uint32 + byte + many byte)
*/
func (m Msg) Encode() []byte {
	length := 4 + 1 + len(m.Data)
	out := make([]byte, 10)
	out[0] = 'F'
	binary.LittleEndian.PutUint32(out[1:5], uint32(length))
	binary.LittleEndian.PutUint32(out[5:9], m.SN)
	out[9] = byte(m.Type)
	out = append(out, m.Data...)
	return out
}
