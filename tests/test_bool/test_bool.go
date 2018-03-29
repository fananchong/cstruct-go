package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 bool) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.Bool.PackLE(val0)
	b := c.Bool.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.Bool.UnpackLE(l)
	val1b := c.Bool.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test(true)
	test(false)
}