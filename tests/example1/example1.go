package main

import (
	"fmt"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct1 struct {
	F0 bool `c:"bool"`
	//	F1 int8      `c:int8`
	//	F2 float32   `c:float`
	//	F3 float64   `c:double`
	//	F4 string    `c:string`
	//	F5 []byte    `c:binary`
	//	F6 mystruct2 `c:struct`
}

func main() {

	// le
	a := &mystruct1{}
	a.F0 = true
	buf_l := cstruct.PackLE(a)
	fmt.Println("Buf(l):", buf_l)
	b := &mystruct1{}
	if err := cstruct.UnpackLE(buf_l, b); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("NewVal(l):", *b)

	// be
	c := &mystruct1{}
	c.F0 = false
	buf_b := cstruct.PackBE(c)
	fmt.Println("Buf(b):", buf_b)
	d := &mystruct1{}
	if err := cstruct.UnpackBE(buf_b, d); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("NewVal(b):", *d)
}
