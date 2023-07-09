package test

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"reflect"
	"s3l/mpcfgo/internal/misc"
	"s3l/mpcfgo/internal/ot"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReflect(t *testing.T) {
	fmt.Println(reflect.ValueOf(float32(123)).Type().Bits())
	fmt.Println(reflect.ValueOf(float64(123)).Type().Bits())
}

func TestGCD(t *testing.T) {
	fmt.Println("GCD", misc.GCD(1234567, 1356))
}

func TestFindGenerator(t *testing.T) {
	q, _ := new(big.Int).SetString("18446744073709551616", 10)
	r := new(big.Int).SetUint64(rand.Uint64())

	for i := 1; i < 100; i++ {
		gcd := new(big.Int).GCD(nil, nil, q, r)
		if gcd.Int64() == 1 {
			fmt.Println(r)
		}
		r = new(big.Int).Rand(rand.New(rand.NewSource(0)), q)

	}

}

func TestGenerate(t *testing.T) {
	in, err := os.OpenFile("OT_testcase.txt", os.O_WRONLY|os.O_CREATE, 0666)
	letItCrash(err)

	for i := 0; i < 100000; i++ {
		r := make([]byte, 8)
		rand.Read(r)
		fmt.Fprintln(in, r)
		rand.Read(r)
		fmt.Fprintln(in, r)
		fmt.Fprintln(in, strconv.FormatBool(misc.Bool()))
	}
	defer in.Close()

}

func TestXor(t *testing.T) {
	bytes0 := make([]byte, 8)
	bytes1 := make([]byte, 8)

	var ns int64
	for i := 0; i < 100000; i++ {
		rand.Read(bytes0)
		rand.Read(bytes1)
		t1 := time.Now()
		//misc.BytesXorBytes(bytes0, bytes1)
		t := time.Since(t1)
		ns += int64(t / time.Nanosecond)
	}
	println(ns)
	var nd int64
	for i := 0; i < 100000; i++ {
		rand.Read(bytes0)
		rand.Read(bytes1)
		t1 := time.Now()
		new(big.Int).Xor(new(big.Int).SetBytes(bytes0), new(big.Int).SetBytes(bytes1))
		t := time.Since(t1)
		nd += int64(t / time.Nanosecond)
	}
	println(nd)

}

func TestBB(t *testing.T) {
	bools := make([]bool, 32)
	b := misc.BoolsToBytes(bools)
	fmt.Println(b)

	bl := misc.BytesToBools(b)
	fmt.Println(bl)
	assert.EqualValues(t, bools, bl)
}

func TestTranspose(t *testing.T) {
	b := make([][]byte, 128)
	for i := range b {
		b[i] = make([]byte, 13)
		for j := range b[i] {
			b[i][j] = byte(j % 8)
		}
	}
	fmt.Println(len(b), len(b[0]))
	c := ot.Transpose(b, 103)
	fmt.Println(c)
}
