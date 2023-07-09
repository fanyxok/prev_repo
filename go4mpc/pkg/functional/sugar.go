package functional

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Conditional[T any](c bool, x func() T, y func() T) T {
	if c {
		return x()
	} else {
		return y()
	}
}

func LastElem[E any](s []E) (E, bool) {
	// len(([]E)nil) is defined as Zero
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}

func DebugInfo() string {
	_, file, line, b := runtime.Caller(1)
	if b {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	} else {
		return "DebugInfo Can't Recover"
	}
}

type NamedSlice[T any] func(T, T) T

func NewNamedSlice[T any](x []struct {
	name  string
	value T
}) (func(string) T, func(string, T)) {
	//struc := reflect.StructOf([]reflect.StructField{})
	m := make(map[string]T)
	for _, v := range x {
		m[v.name] = v.value
	}
	getter := func(name string) T {
		return m[name]
	}
	setter := func(name string, val T) {
		m[name] = val
	}
	return getter, setter
}

func (ct *NamedSlice[T]) Get(name string) T {
	var x T
	return x
}

func (ct *NamedSlice[T]) Set(name string, v T) {
}
