package main

import (
	"fmt"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct1 struct {
	F1  bool   `c:"bool"`
	F6  int8   `c:"int8"`
	F7  int16  `c:"int16"`
	F9  int32  `c:"int32"`
	F11 int64  `c:"int64"`
	F12 uint8  `c:"uint8"`
	F13 uint16 `c:"uint16"`
	F15 uint32 `c:"uint32"`
	F17 uint64 `c:"uint64"`
}

func main() {
	a := &mystruct1{}
	a.F1 = true
	a.F6 = -128
	a.F7 = -32768
	a.F9 = -2147483648
	a.F11 = -9223372036854775808
	a.F12 = 255
	a.F13 = 32767
	a.F15 = 4294967295
	a.F17 = 18446744073709551615
	test1(a, cstruct.LE)

	fmt.Println("\n\n")

	a = &mystruct1{}
	a.F1 = false
	a.F6 = 127
	a.F7 = 32767
	a.F9 = 2147483647
	a.F11 = 9223372036854775807
	a.F12 = 1
	a.F13 = 1
	a.F15 = 1
	a.F17 = 1
	test1(a, cstruct.LE)

	fmt.Println("\n\n")

	a = &mystruct1{}
	a.F1 = true
	a.F6 = -128
	a.F7 = -32768
	a.F9 = -2147483648
	a.F11 = -9223372036854775808
	a.F12 = 255
	a.F13 = 32767
	a.F15 = 4294967295
	a.F17 = 18446744073709551615
	test1(a, cstruct.BE)

	fmt.Println("\n\n")

	a = &mystruct1{}
	a.F1 = false
	a.F6 = 127
	a.F7 = 32767
	a.F9 = 2147483647
	a.F11 = 9223372036854775807
	a.F12 = 1
	a.F13 = 1
	a.F15 = 1
	a.F17 = 1
	test1(a, cstruct.BE)

}

func test1(a *mystruct1, order cstruct.ByteOrder) {
	cstruct.CurrentByteOrder = order
	buf_l, _ := cstruct.Marshal(a)
	fmt.Println("Buf(l):", buf_l)
	b := &mystruct1{}
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("NewVal(l):", *b)

	fmt.Println("b.F1 =", b.F1)
	fmt.Println("b.F6 =", b.F6)
	fmt.Println("b.F7 =", b.F7)
	fmt.Println("b.F9 =", b.F9)
	fmt.Println("b.F11 =", b.F11)
	fmt.Println("b.F12 =", b.F12)
	fmt.Println("b.F13 =", b.F13)
	fmt.Println("b.F15 =", b.F15)
	fmt.Println("b.F17 =", b.F17)
}
