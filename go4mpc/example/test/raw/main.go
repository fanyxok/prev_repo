package main

const Const0 = 100
const Const1 = 101
const Const2 = Const0 + Const1

func constp(a, b sint32) sint32 {
	var c sint32 = a + b
	for i := 0; i < Const2-Const0; i = i + 1 {
		var d sint32 = c + a
		c = d + b
	}
	return c
}

func expand0(a, b sint32) sint32 {
	var c sint32 = a + b
	for i := 0; i < Const0; i = i + 1 {
		var d sint32 = c + a
		c = d + b
	}
	return c
}

func expand1(ax, bx []sint32) sint32 {
	var c sint32 = 1
	c = c + ax[0]
	c = c + bx[0]
	for i := 1; i < Const0; i = i + 1 {
		var d sint32 = c + ax[i]
		c = d + bx[i]
	}
	return c
}

func expand2(ax, bx []sint32) []sint32 {
	var c []sint32 = make([]sint32, Const0)
	for i := 0; i < Const0; i = i + 1 {
		c[i] = ax[i] + bx[i]
	}
	return c
}

func nestfor(ax, bx []sint32) sint32 {
	var c sint32 = 1
	for i := 0; i < Const0; i = i + 1 {
		for j := 0; j < Const0; j = j + 1 {
			var d = c + ax[i]
			c = d + bx[j]
		}
	}
	return c
}

func slice(ax []sint32) []sint32 {
	var d []sint32 = ax[:Const0]
	return d
}

func fncall(ax []sint32) []sint32 {
	var c []sint32 = slice(ax)
	for i := 0; i < Const0; i = i + 1 {
		c = slice(c)
	}
	return c
}
