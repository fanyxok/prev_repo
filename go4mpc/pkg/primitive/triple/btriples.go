package triple

import (
	"crypto/rand"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	"s3l/mpcfgo/internal/ot/rot"
	"s3l/mpcfgo/pkg/type/pub"
)

var xor = func(x pub.PubNum, y pub.PubNum) pub.Bool {
	return x.(pub.Bool) != y.(pub.Bool)
}

func (ct *Triples) BooleanTripleN(net network.Network, n int) {
	if net.Server {
		as := make([]byte, (n+7)/8)
		rand.Read(as)
		a := misc.BytesToBools(as)
		xa_ := rot.RecvN(net, a[:n])
		for i := 0; i < n; i++ {
			r := pub.ZeroBool.Rand()
			net.Send(network.NewMsg(00, network.Sharing, r.Bytes()))
			ct.A[i] = xor(r, pub.Bool(a[i]))
			ct.B[i] = pub.ZeroBool.Decode(net.Recv().Data)
			ct.C[i] = pub.ZeroBool.Decode(xa_[i])
		}
	} else {
		x0, x1 := rot.SendN(net, n, 1)
		for i := 0; i < n; i++ {
			ct.C[i] = pub.ZeroBool.Decode(x0[i])
			ct.A[i] = pub.ZeroBool.Decode(net.Recv().Data)
			r := pub.ZeroBool.Rand()
			net.Send(network.NewMsg(00, network.Sharing, r.Bytes()))
			ct.B[i] = xor(pub.ZeroBool.Decode(x1[i]), pub.ZeroBool.Decode(x0[i]))
			ct.B[i] = xor(ct.B[i], r)
		}
	}
}
