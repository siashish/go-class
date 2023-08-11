package main

import (
	"fmt"
	"math/cmplx"
)

// bool

// string

// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr

// -128 to 127
// 0 to 255

// byte // alias for uint8 // ASCII
// rune // alias for int32 // UNICODE

// float32 float64

// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T value:%v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value:%v\n", z, z)
}
