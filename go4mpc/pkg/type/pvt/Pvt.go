package pvt

import (
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/pkg/type/pub"
	"s3l/mpcfgo/pkg/type/value"
)

type PvtNum interface {
	value.Value
	New(net network.Network, num pub.PubNum) PvtNum
	NewN(net network.Network, num []pub.PubNum) []PvtNum
	NewFrom(net network.Network) PvtNum
	NewFromN(net network.Network) []PvtNum
	Declassify(network network.Network)
	GetPlaintext() pub.PubNum
	GetShare() pub.PubNum
	// Operator
	Add(network.Network, value.Value) PvtNum
	Sub(network.Network, value.Value) PvtNum
	Mul(network.Network, value.Value) PvtNum
	Div(network.Network, value.Value) PvtNum

	Not(network.Network) PvtNum
	And(network.Network, value.Value) PvtNum
	Or(network.Network, value.Value) PvtNum

	Eq(network.Network, value.Value) PvtNum
	Gt(network.Network, value.Value) PvtNum
	Lt(network.Network, value.Value) PvtNum

	Shr(network.Network, value.Value) PvtNum
	Shl(network.Network, value.Value) PvtNum

	Mux(network.Network, value.Value, value.Value) PvtNum
}
