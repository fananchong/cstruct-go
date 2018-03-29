package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 int8) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.Int8.PackLE(val0)
	b := c.Int8.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.Int8.UnpackLE(l)
	val1b := c.Int8.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test(-128)
	test(-127)
	test(-126)
	test(-2)
	test(-1)
	test(0)
	test(1)
	test(2)
	test(126)
	test(127)
}
