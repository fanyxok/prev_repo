package main

type sbool bool
type sint8 int8
type sint16 int16
type sint32 int32
type sint64 int64
type sfloat32 float32
type sfloat64 float64

var b1 func(int) sbool
var i8 func(int) sint8
var i16 func(int) sint16
var i32 func(int) sint32
var i64 func(int) sint64
var f32 func(int) sfloat32
var f64 func(int) sfloat64

var b1n func(int, int) []sbool
var i8n func(int, int) []sint8
var i16n func(int, int) []sint16
var i32n func(int, int) []sint32
var i64n func(int, int) []sint64
var f32n func(int, int) []sfloat32
var f64n func(int, int) []sfloat64

var Addi8 func(x sint8, y int8) sint8
var Addi16 func(x sint16, y int16) sint16
var Addi32 func(x sint32, y int32) sint32
var Addi64 func(x sint64, y int64) sint64
var Addf32 func(x sfloat32, y float32) sfloat32
var Addf64 func(x sfloat64, y float64) sfloat64

var Subi8 func(x sint8, y int8) sint8
var Subi16 func(x sint16, y int16) sint16
var Subi32 func(x sint32, y int32) sint32
var Subi64 func(x sint64, y int64) sint64
var Subf32 func(x sfloat32, y float32) sfloat32
var Subf64 func(x sfloat64, y float64) sfloat64

var Muli8 func(x sint8, y int8) sint8
var Muli16 func(x sint16, y int16) sint16
var Muli32 func(x sint32, y int32) sint32
var Muli64 func(x sint64, y int64) sint64
var Mulf32 func(x sfloat32, y float32) sfloat32
var Mulf64 func(x sfloat64, y float64) sfloat64

var Divi8 func(x sint8, y int8) sint8
var Divi16 func(x sint16, y int16) sint16
var Divi32 func(x sint32, y int32) sint32
var Divi64 func(x sint64, y int64) sint64
var Divf32 func(x sfloat32, y float32) sfloat32
var Divf64 func(x sfloat64, y float64) sfloat64

var Shri8 func(l sint8, m int) sint8
var Shri16 func(l sint16, m int) sint16
var Shri32 func(l sint32, m int) sint32
var Shri64 func(l sint64, m int) sint64

var Shli8 func(l sint8, m int) sint8
var Shli16 func(l sint16, m int) sint16
var Shli32 func(l sint32, m int) sint32
var Shli64 func(l sint64, m int) sint64

var Gti8 func(x sint8, y int8) sint8
var Gti16 func(x sint16, y int16) sint16
var Gti32 func(x sint32, y int32) sint32
var Gti64 func(x sint64, y int64) sint64
var Gtf32 func(x sfloat32, y float32) sfloat32
var Gtf64 func(x sfloat64, y float64) sfloat64

var Lti8 func(l sint8, m int8) sint8
var Lti16 func(l sint16, m int16) sint16
var Lti32 func(l sint32, m int32) sint32
var Lti64 func(l sint64, m int64) sint64
var Ltf32 func(x sfloat32, y float32) sfloat32
var Ltf64 func(x sfloat64, y float64) sfloat64

var Eqb1 func(l sbool, m bool) sbool
var Eqi8 func(l sint8, m int8) sint8
var Eqi16 func(l sint16, m int16) sint16
var Eqi32 func(l sint32, m int32) sint32
var Eqi64 func(l sint64, m int64) sint64
var Eqf32 func(x sfloat32, y float32) sfloat32
var Eqf64 func(x sfloat64, y float64) sfloat64

var Andb1 func(l sbool, m bool) sbool
var Orb1 func(l sbool, m bool) sbool

var Muxb1 func(c sbool, x sbool, y sbool) sbool
var Muxi8 func(c sbool, x sint8, y sint8) sbool
var Muxi16 func(c sbool, x sint16, y sint16) sbool
var Muxi32 func(c sbool, x sint32, y sint32) sbool
var Muxi64 func(c sbool, x sint64, y sint64) sbool
var Muxf32 func(c sbool, x sfloat32, y sfloat32) sfloat32
var Muxf64 func(c sbool, x sfloat64, y sfloat64) sfloat64
