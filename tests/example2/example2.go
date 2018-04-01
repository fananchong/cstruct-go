package main

import (
	"fmt"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct1 struct {
	F1  bool  `c:"bool"`
	F6  int8  `c:"int8"`
	F12 uint8 `c:"uint8"`
}

func main() {

	// le
	a := &mystruct1{}
	a.F1 = true
	a.F6 = -128
	a.F12 = 255

	buf_l, _ := cstruct.Marshal(a)
	fmt.Println("Buf(l):", buf_l)
	b := &mystruct1{}
	if err := cstruct.UnpackLE(buf_l, b); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("NewVal(l):", *b)

	fmt.Println("b.F1 =", b.F1)
	fmt.Println("b.F6 =", b.F6)
	fmt.Println("b.F12 =", b.F12)
}
