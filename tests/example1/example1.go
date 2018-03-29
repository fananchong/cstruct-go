package main

import (
	"fmt"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct3 struct {
	F3 float64 `c:"double"`
	F4 string  `c:"string"`
}

type mystruct2 struct {
	F6  int8  `c:"int8"`
	F7  int16 `c:"int16"`
	F8  int32 `c:"int24"`
	F9  int32 `c:"int32"`
	F10 int64 `c:"int40"`
	F11 int64 `c:"int64"`
	S0  mystruct3
}

type mystruct1 struct {
	N0  uint32
	F1  bool    `c:"bool"`
	F2  float32 `c:"float"`
	F3  float64 `c:"double"`
	S0  mystruct2
	F4  string `c:"string"`
	F5  []byte `c:"binary"`
	F6  int8   `c:"int8"`
	F7  int16  `c:"int16"`
	F8  int32  `c:"int24"`
	F9  int32  `c:"int32"`
	F10 int64  `c:"int40"`
	F11 int64  `c:"int64"`
	N1  float32
	F12 uint8  `c:"uint8"`
	F13 uint16 `c:"uint16"`
	F14 uint32 `c:"uint24"`
	F15 uint32 `c:"uint32"`
	F16 uint64 `c:"uint40"`
	F17 uint64 `c:"uint64"`
	N2  bool
}

func main() {

	// le
	a := &mystruct1{}
	a.F1 = true
	a.F2 = 0.98
	a.F3 = 999.777
	a.F4 = "hello"
	a.F5 = []byte{1, 2, 3, 4}
	a.F6 = 126
	a.F7 = 32766
	a.F8 = -8388608
	a.F9 = -2147483646
	a.F10 = 549755813887
	a.F11 = -9223372036854774807
	a.F12 = 254
	a.F13 = 65534
	a.F14 = 16777214
	a.F15 = 4294967295
	a.F16 = 1099511627770
	a.F17 = 18446744073709551614
	a.S0.F6 = -126
	a.S0.F7 = -32766
	a.S0.F8 = 8388600
	a.S0.F9 = 2147483640
	a.S0.F10 = -549755813880
	a.S0.F11 = 9223372036854774800
	a.S0.S0.F3 = 100.9
	a.S0.S0.F4 = "world"

	buf_l := cstruct.PackLE(a)
	fmt.Println("Buf(l):", buf_l)
	b := &mystruct1{}
	if err := cstruct.UnpackLE(buf_l, b); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("NewVal(l):", *b)

	// be
	buf_b := cstruct.PackBE(a)
	fmt.Println("Buf(b):", buf_b)
	d := &mystruct1{}
	if err := cstruct.UnpackBE(buf_b, d); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("NewVal(b):", *d)

	fmt.Println("b.F1 =", b.F1)
	fmt.Println("b.F2 =", b.F2)
	fmt.Println("b.F3 =", b.F3)
	fmt.Println("b.F4 =", b.F4)
	fmt.Println("b.F5 =", b.F5)
	fmt.Println("b.F6 =", b.F6)
	fmt.Println("b.F7 =", b.F7)
	fmt.Println("b.F8 =", b.F8)
	fmt.Println("b.F9 =", b.F9)
	fmt.Println("b.F10 =", b.F10)
	fmt.Println("b.F11 =", b.F11)
	fmt.Println("b.F12 =", b.F12)
	fmt.Println("b.F13 =", b.F13)
	fmt.Println("b.F14 =", b.F14)
	fmt.Println("b.F15 =", b.F15)
	fmt.Println("b.F16 =", b.F16)
	fmt.Println("b.F17 =", b.F17)
	fmt.Println("b.F17 =", b.F17)

	fmt.Println("b.S0.F6 =", b.S0.F6)
	fmt.Println("b.S0.F7 =", b.S0.F7)
	fmt.Println("b.S0.F8 =", b.S0.F8)
	fmt.Println("b.S0.F9 =", b.S0.F9)
	fmt.Println("b.S0.F10 =", b.S0.F10)
	fmt.Println("b.S0.F11 =", b.S0.F11)
	fmt.Println("b.S0.S0.F3 =", b.S0.S0.F3)
	fmt.Println("b.S0.S0.F4 =", b.S0.S0.F4)
}
