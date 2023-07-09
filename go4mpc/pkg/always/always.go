package always

import "fmt"

type Number interface {
	bool | string | int8 | int16 | int | int32 | int64 | uint8 | uint16 | uint | uint32 | uint64 | float32 | float64
}

func Eq[T Number](expected T, actual T) bool {
	if expected != actual {
		fmt.Println("###Panic###\nNot equal: \n"+
			"expected:", expected, "\n"+
			"actual  :", actual, "\n------")
		panic("Not equal")
	}
	return true
}

func NotEq[T Number](expected T, actual T) bool {
	if expected == actual {
		fmt.Println("Equal: \n"+
			"expected not:", expected, "\n"+
			"actual      :", actual)
		panic("Equal")
	}
	return true
}

func Nil(err error) bool {
	if err != nil {
		panic("error is not nil:" + err.Error())
	}
	return true
}

func NotNil(x any) {
	if x == nil {
		panic("Expected not Nil, find Nil")
	}
}
