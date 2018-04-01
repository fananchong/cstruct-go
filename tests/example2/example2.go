package main

import (
	"fmt"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct1 struct {
	F1  bool   `c:"bool"`
	F6  int8   `c:"int8"`
	F7  int16  `c:"int16"`
	F12 uint8  `c:"uint8"`
	F13 uint16 `c:"uint16"`
}

func main() {
	a := &mystruct1{}
	a.F1 = true
	a.F6 = -128
	a.F7 = -32768
	a.F12 = 255
	a.F13 = 32767
	test1(a, cstruct.LE)

	fmt.Println("\n\n")

	a = &mystruct1{}
	a.F1 = false
	a.F6 = 127
	a.F7 = 32767
	a.F12 = 1
	a.F13 = 1
	test1(a, cstruct.LE)

	fmt.Println("\n\n")

	a = &mystruct1{}
	a.F1 = true
	a.F6 = -128
	a.F7 = -32768
	a.F12 = 255
	a.F13 = 32767
	test1(a, cstruct.BE)

	fmt.Println("\n\n")

	a = &mystruct1{}
	a.F1 = false
	a.F6 = 127
	a.F7 = 32767
	a.F12 = 1
	a.F13 = 1
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
	fmt.Println("b.F12 =", b.F12)
	fmt.Println("b.F13 =", b.F13)
}
