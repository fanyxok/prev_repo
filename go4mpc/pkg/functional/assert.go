package functional

import "fmt"

func Assert[T comparable](expected T, actual T) {
	if expected != actual {
		panic(fmt.Sprintf("%s expexted: %#v, actual: %#v\n", FileLine(), expected, actual))
	}
}
