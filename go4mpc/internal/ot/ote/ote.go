package ote

import (
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/rand"
	config "s3l/mpcfgo/config"
	"s3l/mpcfgo/internal/encrypt/prf"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/network"
	OT "s3l/mpcfgo/internal/ot"
	"s3l/mpcfgo/internal/ot/baseot"
	"sync"
	"unsafe"

	"s3l/mpcfgo/pkg/always"
	"s3l/mpcfgo/pkg/fast"
	"time"

	"github.com/forgoer/openssl"
)

/*
*
ALSZ OTE
*
*/

var OtInit_S []byte
var OtInit_Sbs []bool
var OtInit_K []*rand.Rand
var OtInit_K0 []*rand.Rand
var OtInit_K1 []*rand.Rand

func init() {
	OtInit_S = make([]byte, config.SymByte)
	OtInit_Sbs = make([]bool, config.SymK)
	OtInit_K = make([]*rand.Rand, config.SymK)
	OtInit_K0 = make([]*rand.Rand, config.SymK)
	OtInit_K1 = make([]*rand.Rand, config.SymK)
}

func RAND(b []byte) {
	_, err := rand.Read(b)
	if err != nil {
		log.Panicf("RAND: %v\n", err)
	}
}

/*
* InitOT initializes the OT extension
*
 */
func InitOtSender(net network.Network) {
	RAND(OtInit_S)
	fast.Bytes2Bools(OtInit_Sbs, OtInit_S)
	kb := baseot.Recv128(net, OtInit_Sbs)
	var wg sync.WaitGroup
	workload := config.SymK / 4
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * workload
			end := start + workload
			for i := start; i < end; i++ {
				a := binary.LittleEndian.Uint64(kb[i][0:8])
				b := binary.LittleEndian.Uint64(kb[i][8:16])
				OtInit_K[i] = rand.New(rand.NewSource(int64(a ^ b)))
			}
		}(i)
	}
	wg.Wait()
}

func InitOtReceiver(net network.Network) {
	kb0 := make([][]byte, config.SymK)
	kb1 := make([][]byte, config.SymK)
	var wg sync.WaitGroup
	workload := config.SymK / 4
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * workload
			end := start + workload
			for i := start; i < end; i++ {
				kb0[i] = make([]byte, config.SymByte)
				kb1[i] = make([]byte, config.SymByte)
				RAND(kb0[i])
				RAND(kb1[i])
			}
		}(i)
	}
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		baseot.Send128(net, kb0, kb1)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < config.SymK; i++ {
			a0 := binary.LittleEndian.Uint64(kb0[i][0:8])
			b0 := binary.LittleEndian.Uint64(kb0[i][8:16])
			OtInit_K0[i] = rand.New(rand.NewSource(int64(a0 ^ b0)))
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < config.SymK; i++ {
			a1 := binary.LittleEndian.Uint64(kb1[i][0:8])
			b1 := binary.LittleEndian.Uint64(kb1[i][8:16])
			OtInit_K1[i] = rand.New(rand.NewSource(int64(a1 ^ b1)))
		}
	}()
	wg.Wait()
}

/*
* ALSZ OTE
*
 */
type OTE_ struct{}
type COT_ struct {
	F       func(...interface{})
	EleSize int // in byte
}
type ROT_ struct{}

var (
	OTE = OTE_{}
	COT = COT_{}
	ROT = ROT_{}
)

func HashValues(dst []byte, src []byte, bytes int) {
	if bytes > len(src) {
		tmp := make([]byte, bytes)
		copy(tmp, src)
		cipher.NewCFBEncrypter(prf.AES_Key_Fixed, prf.IV).XORKeyStream(dst, tmp)
	} else {
		cipher.NewCFBEncrypter(prf.AES_Key_Fixed, prf.IV).XORKeyStream(dst, src[:bytes])
	}
}

var _GCM, _ = cipher.NewGCM(prf.AES_Key_Fixed)

func HashValues0(dst []byte, src []byte, bytes int) {
	if bytes > len(src) {
		tmp := make([]byte, bytes)
		copy(tmp, src)
		_GCM.Seal(dst, prf.IV[:12], tmp, nil)
	} else {
		_GCM.Seal(dst, prf.IV[:12], src[:bytes], nil)
	}
}

func HashValues1(dst, src []byte, bytes int) {
	paddingCount := prf.AesBlockSizeInBytes - len(src)%prf.AesBlockSizeInBytes
	if paddingCount == 0 {
		_ECB.CryptBlocks(dst, src)
	} else {
		src_ := fast.ZerosPadding(src, paddingCount)
		dst_ := make([]byte, len(src_))
		_ECB.CryptBlocks(dst_, src_)
		copy(dst, dst_[:bytes])
	}
}

var _ECB = openssl.NewECBEncrypter(prf.AES_Key_Fixed)

func HashValuesJ(dst []byte, src []byte, j int, bytes int) {
	paddingCount := prf.AesBlockSizeInBytes - len(src)%prf.AesBlockSizeInBytes
	if paddingCount == prf.AesBlockSizeInBytes {
		head := (*int)(unsafe.Pointer(&src[0]))
		*head ^= j
		//fmt.Println(len(dst), len(src))
		dst_ := make([]byte, len(src))
		_ECB.CryptBlocks(dst_, src)
		copy(dst, dst_[:bytes])
		*head ^= j
	} else {
		//panic("should not be ")
		src_ := fast.ZerosPadding(src, paddingCount)
		head := (*int)(unsafe.Pointer(&src_[0]))
		*head ^= j
		dst_ := make([]byte, len(src_))
		_ECB.CryptBlocks(dst_, src_)
		copy(dst, dst_[:bytes])
	}
}

func HashValuesJBatch128(dst []byte, src []byte, Jstart, Jend int, bytes int) {
	for i := Jstart; i < Jend; i++ {
		head := (*int)(unsafe.Pointer(&src[(i-Jstart)*16]))
		*head ^= i
	}
	dst_ := make([]byte, len(src))
	_ECB.CryptBlocks(dst_, src)
	for i := 0; i < Jend-Jstart; i++ {
		copy(dst[i*bytes:(i+1)*bytes], dst_[i*16:i*16+bytes])
	}
	for i := Jstart; i < Jend; i++ {
		head := (*int)(unsafe.Pointer(&src[(i-Jstart)*16]))
		*head ^= i
	}
}

const OTE_Workers = 1

// Sender own OtInit_K
func (ct OTE_) SendN(net network.Network, x0, x1 []byte, nOTs, bytes int) {
	nOTsByte := (nOTs + 7) / 8

	var wg sync.WaitGroup
	t := time.Now()
	var colQ []byte = make([]byte, config.SymK*nOTsByte)
	var Y []byte = make([]byte, 2*bytes*nOTs)
	// U has 16 * nOTs bytes
	var U []byte = net.Recv().Data
	fmt.Printf("Send-recvU: %v\n", time.Since(t))
	t = time.Now()
	// colQ divided by 8
	wg.Add(OTE_Workers)
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			for i := p; i < config.SymK; i += OTE_Workers {
				lbound := i * nOTsByte
				rbound := (i + 1) * nOTsByte
				OtInit_K[i].Read(colQ[lbound:rbound])
				if OtInit_Sbs[i] {
					fast.Xor(colQ[lbound:rbound], colQ[lbound:rbound], U[lbound:rbound], nOTsByte)
				}
			}

		}(p)
	}
	wg.Wait()
	fmt.Printf("Send-colQ: %v\n", time.Since(t))
	t = time.Now()
	rowQ := make([]byte, config.SymK*nOTsByte)
	fast.Transpose128n(rowQ, colQ, nOTsByte)
	fmt.Printf("Send-Transpose: %v\n", time.Since(t))
	t = time.Now()
	xBytes := bytes * nOTs
	HQ_QS := make([]byte, xBytes)
	//HQS := make([]byte, xBytes)
	QS := rowQ //make([]byte, config.SymByte*nOTs)
	wg.Add(OTE_Workers)
	pload := nOTs/OTE_Workers + 1
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			ld := p * pload
			rd := fast.MinInt(ld+pload, nOTs)
			HashValuesJBatch128(HQ_QS[ld*bytes:rd*bytes], rowQ[ld*config.SymByte:rd*config.SymByte], ld, rd, bytes)
			fast.Xor(Y[ld*bytes:rd*bytes], HQ_QS[ld*bytes:rd*bytes], x0[ld*bytes:rd*bytes], rd-ld)
			fast.XorBatch128(QS[ld*config.SymByte:rd*config.SymByte], rowQ[ld*config.SymByte:rd*config.SymByte], OtInit_S, rd-ld)
			HashValuesJBatch128(HQ_QS[ld*bytes:rd*bytes], QS[ld*config.SymByte:rd*config.SymByte], ld, rd, bytes)
			fast.Xor(Y[xBytes+ld*bytes:xBytes+rd*bytes], HQ_QS[ld*bytes:rd*bytes], x1[ld*bytes:rd*bytes], rd-ld)
		}(p)
	}
	wg.Wait()
	fmt.Printf("Send-HQ+Xor3+HQS: %v\n", time.Since(t))
	t = time.Now()
	net.Send(network.NewMsg(0, network.OTe, Y))
	fmt.Printf("Send-sendY: %v\n", time.Since(t))
}

func (ct OTE_) RecvN(net network.Network, X []byte, r []bool, nOTs, bytes int) {
	nOTsByte := (nOTs + 7) / 8
	var wg sync.WaitGroup
	var colT = make([]byte, config.SymK*nOTsByte)
	var U = make([]byte, config.SymK*nOTsByte)
	var R = make([]byte, nOTsByte)
	fast.Bools2Bytes(R, r)
	t := time.Now()
	wg.Add(OTE_Workers)
	pload := config.SymK/OTE_Workers + 1
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			for i := p; i < config.SymK; i += OTE_Workers {
				lbound := i * nOTsByte
				rbound := (i + 1) * nOTsByte
				// colT[i]
				OtInit_K0[i].Read(colT[lbound:rbound])
				// G(K1)
				OtInit_K1[i].Read(U[lbound:rbound])
				// G(K1) ^ r
				fast.Xor(U[lbound:rbound], U[lbound:rbound], R, nOTsByte)
				// G(K1) ^ r ^ colT[i]
				//fast.Xor(U[lbound:rbound], U[lbound:rbound], colT[lbound:rbound], nOTsByte)
			}
			//fast.Xor(U[ld*nOTsByte:rd*nOTsByte], U[ld*nOTsByte:rd*nOTsByte], colT[ld*nOTsByte:rd*nOTsByte], (rd-ld)*nOTsByte)
		}(p)
	}
	wg.Wait()
	fast.Xor(U, U, colT, config.SymK*nOTsByte)
	fmt.Printf("Recv-computeU: %v\n", time.Since(t))
	t = time.Now()
	net.Send(network.NewMsg(0, network.OTe, U))
	fmt.Printf("Recv-sendU: %v\n", time.Since(t))
	t = time.Now()
	Y := net.Recv().Data
	fmt.Printf("Recv-recvY: %v\n", time.Since(t))
	rowT := make([]byte, config.SymK*nOTsByte)
	t = time.Now()
	fast.Transpose128n(rowT, colT, nOTsByte)
	fmt.Printf("Transpose: %v\n", time.Since(t))

	t = time.Now()
	xBytes := bytes * nOTs
	HT := make([]byte, xBytes)
	wg.Add(OTE_Workers)
	pload = nOTs/OTE_Workers + 1
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			ld := p * pload
			rd := fast.MinInt(ld+pload, nOTs)
			HashValuesJBatch128(HT[ld*bytes:rd*bytes], rowT[ld*config.SymByte:rd*config.SymByte], ld, rd, bytes)
			for i := ld; i < rd; i++ {
				if r[i] {
					fast.Xor(X[i*bytes:(i+1)*bytes], Y[xBytes+i*bytes:xBytes+(i+1)*bytes], HT[i*bytes:(i+1)*bytes], bytes)
				} else {
					fast.Xor(X[i*bytes:(i+1)*bytes], Y[i*bytes:(i+1)*bytes], HT[i*bytes:(i+1)*bytes], bytes)
				}
			}
		}(p)
	}
	wg.Wait()
	fmt.Printf("Recv-Hash_Xor: %v\n", time.Since(t))
}

func (ct COT_) SendN(net network.Network, X0, X1 []byte, nOTs, bytes int) {
	nOTsByte := (nOTs + 7) / 8

	var wg sync.WaitGroup
	t := time.Now()
	var colQ []byte = make([]byte, config.SymK*nOTsByte)
	var Y []byte = make([]byte, 2*bytes*nOTs)
	// U has 16 * nOTs bytes
	var U []byte = net.Recv().Data
	fmt.Printf("Send-recvU: %v\n", time.Since(t))
	t = time.Now()
	// colQ divided by 8
	wg.Add(OTE_Workers)
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			for i := p; i < config.SymK; i += OTE_Workers {
				lbound := i * nOTsByte
				rbound := (i + 1) * nOTsByte
				OtInit_K[i].Read(colQ[lbound:rbound])
				if OtInit_Sbs[i] {
					fast.Xor(colQ[lbound:rbound], colQ[lbound:rbound], U[lbound:rbound], nOTsByte)
				}
			}

		}(p)
	}
	wg.Wait()
	fmt.Printf("Send-colQ: %v\n", time.Since(t))
	t = time.Now()
	rowQ := make([]byte, config.SymK*nOTsByte)
	fast.Transpose128n(rowQ, colQ, nOTsByte)
	fmt.Printf("Send-Transpose: %v\n", time.Since(t))
	t = time.Now()
	xBytes := bytes * nOTs
	HQ_QS := make([]byte, xBytes)
	//HQS := make([]byte, xBytes)
	QS := rowQ //make([]byte, config.SymByte*nOTs)
	wg.Add(OTE_Workers)
	nEle := nOTs / ct.EleSize
	eload := int(math.Ceil(float64(nEle) / OTE_Workers))
	pload := eload * ct.EleSize
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			ld := fast.MinInt(p*pload, nOTs)
			rd := fast.MinInt(ld+pload, nOTs)
			// X0 = H(j, Q[j])
			HashValuesJBatch128(X0[ld*bytes:rd*bytes], rowQ[ld*config.SymByte:rd*config.SymByte], ld, rd, bytes)
			// X1 = F(j,X0)
			ct.F(X1[ld*bytes:rd*bytes], X0[ld*bytes:rd*bytes], ld, rd-ld)
			//copy(X1, X0)
			// QS = Q[j]^S
			fast.XorBatch128(QS[ld*config.SymByte:rd*config.SymByte], rowQ[ld*config.SymByte:rd*config.SymByte], OtInit_S, rd-ld)
			// Y[j] = X1 ^ H(j, QS)
			HashValuesJBatch128(HQ_QS[ld*bytes:rd*bytes], QS[ld*config.SymByte:rd*config.SymByte], ld, rd, bytes)
			fast.Xor(Y[ld*bytes:rd*bytes], HQ_QS[ld*bytes:rd*bytes], X1[ld*bytes:rd*bytes], rd-ld)
			//fast.Xor(Y[xBytes+ld*bytes:xBytes+rd*bytes], HQ_QS[ld*bytes:rd*bytes], X1[ld*bytes:rd*bytes], rd-ld)
		}(p)
	}
	wg.Wait()
	fmt.Printf("Send-HQ+Xor3+HQS: %v\n", time.Since(t))
	t = time.Now()
	net.Send(network.NewMsg(0, network.OTe, Y))
	fmt.Printf("Send-sendY: %v\n", time.Since(t))
}

func (ct COT_) RecvN(net network.Network, X []byte, r []bool, nOTs, bytes int) {
	nOTsByte := (nOTs + 7) / 8
	var wg sync.WaitGroup
	var colT = make([]byte, config.SymK*nOTsByte)
	var U = make([]byte, config.SymK*nOTsByte)
	var R = make([]byte, nOTsByte)
	fast.Bools2Bytes(R, r)
	t := time.Now()
	wg.Add(OTE_Workers)
	pload := int(math.Ceil(float64(config.SymK) / OTE_Workers))
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			for i := p; i < config.SymK; i += OTE_Workers {
				lbound := i * nOTsByte
				rbound := (i + 1) * nOTsByte
				// colT[i]
				OtInit_K0[i].Read(colT[lbound:rbound])
				// G(K1)
				OtInit_K1[i].Read(U[lbound:rbound])
				// G(K1) ^ r
				fast.Xor(U[lbound:rbound], U[lbound:rbound], R, nOTsByte)
				// G(K1) ^ r ^ colT[i]
				//fast.Xor(U[lbound:rbound], U[lbound:rbound], colT[lbound:rbound], nOTsByte)
			}
			//fast.Xor(U[ld*nOTsByte:rd*nOTsByte], U[ld*nOTsByte:rd*nOTsByte], colT[ld*nOTsByte:rd*nOTsByte], (rd-ld)*nOTsByte)
		}(p)
	}
	wg.Wait()
	fast.Xor(U, U, colT, config.SymK*nOTsByte)
	fmt.Printf("Recv-computeU: %v\n", time.Since(t))
	t = time.Now()
	net.Send(network.NewMsg(0, network.OTe, U))
	fmt.Printf("Recv-sendU: %v\n", time.Since(t))
	t = time.Now()
	Y := net.Recv().Data
	fmt.Printf("Recv-recvY: %v\n", time.Since(t))
	rowT := make([]byte, config.SymK*nOTsByte)
	t = time.Now()
	fast.Transpose128n(rowT, colT, nOTsByte)
	fmt.Printf("Transpose: %v\n", time.Since(t))

	t = time.Now()
	//xBytes := bytes * nOTs
	//HT := make([]byte, xBytes)
	wg.Add(OTE_Workers)
	eN := nOTs / ct.EleSize
	eload := int(math.Ceil(float64(eN) / OTE_Workers))
	pload = eload * ct.EleSize
	for p := 0; p < OTE_Workers; p++ {
		go func(p int) {
			defer wg.Done()
			ld := fast.MinInt(p*pload, nOTs)
			rd := fast.MinInt(ld+pload, nOTs)
			HashValuesJBatch128(X[ld*bytes:rd*bytes], rowT[ld*config.SymByte:rd*config.SymByte], ld, rd, bytes)
			for i := ld; i < rd; i++ {
				if r[i] {
					fast.Xor(X[i*bytes:(i+1)*bytes], Y[i*bytes:(i+1)*bytes], X[i*bytes:(i+1)*bytes], bytes)
				}
			}
		}(p)
	}
	wg.Wait()
	fmt.Printf("Recv-Hash_Xor: %v\n", time.Since(t))
}

var oteSeed [2][]*rand.Rand
var oteKey []*rand.Rand
var oteChoice []bool
var initialized = false

func OteSeed(net network.Network) [2][]*rand.Rand {
	if !initialized {
		start := time.Now()
		oteSeed[0], oteSeed[1] = make([]*rand.Rand, config.SymK), make([]*rand.Rand, config.SymK)
		// sample k random seeds, each length K
		seed0, seed1 := make([][]byte, config.SymK), make([][]byte, config.SymK)
		seed := rand.New(rand.NewSource(time.Now().UnixNano()))
		var wg sync.WaitGroup
		for i := 0; i < config.SymK; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done() // 函数结束时，通知此 wait 任务已经完成
				seed0[i], seed1[i] = make([]byte, config.SymByte), make([]byte, config.SymByte)
				seed.Read(seed0[i])
				seed.Read(seed1[i])
			}(i)
		}
		wg.Wait()
		for i := 0; i < config.SymK; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done() // 函数结束时，通知此 wait 任务已经完成
				oteSeed[0][i] = rand.New(rand.NewSource(int64(misc.DecodeInt64(seed0[i][0:8]) ^ misc.DecodeInt64(seed0[i][8:16]))))
				oteSeed[1][i] = rand.New(rand.NewSource(int64(misc.DecodeInt64(seed1[i][0:8]) ^ misc.DecodeInt64(seed1[i][8:16]))))
			}(i)
		}
		{
			wg.Add(1)
			go func() {
				defer wg.Done() // 函数结束时，通知此 wait 任务已经完成
				baseot.SendN(net, seed0, seed1)
			}()
		}

		wg.Wait()
		initialized = true
		log.Printf("BaseOT Time: %v\n", time.Since(start))
	}
	return oteSeed
}

func OteKey(net network.Network) ([]bool, []*rand.Rand) {
	if !initialized {
		// sample K randmon bits.
		oteChoice = make([]bool, config.SymK)
		oteKey = make([]*rand.Rand, config.SymK)
		for i := 0; i < config.SymK; i++ {
			oteChoice[i] = misc.Bool()
		}
		// recv seeds. [K x K] bits
		seed := baseot.RecvN(net, oteChoice)
		var wg sync.WaitGroup
		for i := 0; i < config.SymK; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done() // 函数结束时，通知此 wait 任务已经完成
				oteKey[i] = rand.New(rand.NewSource(int64(misc.DecodeInt64(seed[i][0:8]) ^ misc.DecodeInt64(seed[i][8:16]))))
			}(i)
		}
		wg.Wait()
		initialized = true
	}
	return oteChoice, oteKey
}

/*
x0 has m L bit string, occupy m * L/8 bytes
x1 has m L bit string, occupy m * L/8 bytes
*/
func SendN(net network.Network, x0 [][]byte, x1 [][]byte) {
	always.Eq(len(x0), len(x1))
	m := len(x0)
	l := len(x0[0]) * 8
	println("Sending", m, "messages. Each has", l, "bits.")

	net.Send(network.NewMsg(0, network.OTe, misc.EncodeInt32(l)))
	s, seeds := OteKey(net)

	colQ := make([][]byte, config.SymK)
	for i := 0; i < config.SymK; i++ {
		if s[i] {
			gk := OT.RandN(seeds[i], m)
			u := net.Recv()
			colQ[i] = make([]byte, (m+7)/8)
			misc.BytesXorBytes(colQ[i], u.Data, gk)
		} else {
			net.Recv()
			colQ[i] = OT.RandN(seeds[i], m)
		}
	}
	S := misc.BoolsToBytes(s)
	yByteLen := (l + 7) / 8

	batch := 8
	for k := 0; k < m/batch; k++ {
		y := make([]byte, (yByteLen+yByteLen)*batch)
		for i := 0; i < batch; i++ {
			j := k*batch + i
			rowQ := OT.Transpose(colQ, j)
			h0 := OT.Remapping(prf.FixedKeyAES.Hash(rowQ), l)
			QS := make([]byte, config.SymByte)
			misc.BytesXorBytes(QS, rowQ, S)
			h1 := OT.Remapping(prf.FixedKeyAES.Hash(QS), l)
			misc.BytesXorBytes(y[2*i*yByteLen:(2*i+1)*yByteLen], x0[j], h0)
			misc.BytesXorBytes(y[(2*i+1)*yByteLen+yByteLen:(2*i+2)*yByteLen], x1[j], h1)
		}
		net.Send(network.NewMsg(11, network.OTe, y))
	}
	for j := m / batch * batch; j < m; j++ {
		rowQ := OT.Transpose(colQ, j)
		h0 := OT.Remapping(prf.FixedKeyAES.Hash(rowQ), l)
		QS := make([]byte, config.SymByte)
		misc.BytesXorBytes(QS, rowQ, S)
		h1 := OT.Remapping(prf.FixedKeyAES.Hash(QS), l)
		y := make([]byte, yByteLen+yByteLen)
		misc.BytesXorBytes(y[:yByteLen], x0[j], h0)
		misc.BytesXorBytes(y[yByteLen:], x1[j], h1)
		net.Send(network.NewMsg(11, network.OTe, y))
	}
}

func RecvN(net network.Network, r []bool) [][]byte {
	m := len(r)
	tmp := net.Recv()
	l := int(misc.DecodeInt32(tmp.Data))
	println("Receiving", m, "messages. Each has", l, "bits")
	seeds := OteSeed(net)
	seeds0, seeds1 := seeds[0], seeds[1]

	//	T is m * K matrix, where colT is a col
	colT := make([][]byte, config.SymK)
	R := misc.BoolsToBytes(r)
	for i := 0; i < config.SymK; i++ {
		colT[i] = OT.RandN(seeds0[i], m)
		u := OT.RandN(seeds1[i], m)
		misc.BytesXorBytes(u, u, colT[i])
		misc.BytesXorBytes(u, u, R)
		net.Send(network.NewMsg(uint32(i), network.OTe, u))
	}

	// reveal the mask of X
	x := make([][]byte, m)
	xByteLen := (l + 7) / 8
	batch := 8
	for k := 0; k < m/batch; k++ {
		y := net.Recv().Data
		for i := 0; i < batch; i++ {
			j := k*batch + i
			rowT := OT.Transpose(colT, i)
			h := OT.Remapping(prf.FixedKeyAES.Hash(rowT), l)
			x[j] = make([]byte, xByteLen)
			if r[i] {
				misc.BytesXorBytes(x[j], y[(2*i)*xByteLen:(2*i+1)*xByteLen], h)
			} else {
				misc.BytesXorBytes(x[j], y[(2*i+1)*xByteLen:(2*i+2)*xByteLen], h)
			}
		}
	}
	for i := m / batch * batch; i < m; i++ {
		rowT := OT.Transpose(colT, i)
		h := OT.Remapping(prf.FixedKeyAES.Hash(rowT), l)
		x[i] = make([]byte, xByteLen)
		y := net.Recv().Data
		if r[i] {
			misc.BytesXorBytes(x[i], y[xByteLen:], h)
		} else {
			misc.BytesXorBytes(x[i], y[:xByteLen], h)
		}
	}
	return x
}
