package functional

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"unsafe"
)

func WithoutLastElem[E any](s []E) ([]E, bool) {
	if len(s) == 0 {
		return s, false
	}
	return s[:len(s)-1], true
}

func FileLine() string {
	_, file, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("[%s, %d]", filepath.Base(file), line)
	return str
}

func ExistValue[K comparable, V comparable](m map[K]V, v V) (K, bool) {
	var z K
	if m == nil {
		return z, false
	}
	for i, val := range m {
		if val == v {
			return i, true
		}
	}
	return z, false
}

func FilterByValue[K comparable, V any](m map[K]V, f func(V) bool) map[K]V {
	mm := make(map[K]V)
	for i, v := range m {
		mm[i] = v
	}
	for k, v := range m {
		if !f(v) {
			delete(mm, k)
		}
	}
	return mm
}
func FilterByKey[K comparable, V any](m map[K]V, f func(K) bool) map[K]V {
	mm := make(map[K]V)
	for i, v := range m {
		mm[i] = v
	}
	for k := range m {
		if !f(k) {
			delete(mm, k)
		}
	}
	return mm
}

type Eface struct {
	Typ, Val unsafe.Pointer
}

func PointerOf(x any) unsafe.Pointer {
	return (*Eface)(unsafe.Pointer(&x)).Val
}
func Kind[T any](x T) reflect.Kind {
	return reflect.ValueOf(x).Kind()
}

func Type2Width(x reflect.Type) int {
	switch x.Kind() {
	case reflect.Float32:
		return 31
	case reflect.Float64:
		return 63
	case reflect.Bool:
		return 1
	default:
		return x.Bits()
	}
}

func Append[T any](sl ...any) []T {
	var temp []T = nil
	var p T
	var s []T
	ptyp := reflect.TypeOf(p)
	styp := reflect.TypeOf(s)
	for i, v := range sl {
		if typ := reflect.TypeOf(v); typ == ptyp {
			temp = append(temp, v.(T))
		} else if typ == styp {
			temp = append(temp, v.([]T)...)
		} else {
			panic(fmt.Sprintf("#Lean# func Append[T any](sl ...any) []T # sl[i] type error %v", reflect.TypeOf(sl[i])))
		}
	}
	return temp
}
