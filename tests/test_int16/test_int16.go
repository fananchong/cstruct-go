package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 int16) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.Int16.PackLE(val0)
	b := c.Int16.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.Int16.UnpackLE(l)
	val1b := c.Int16.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test(-32768)
	test(-32767)
	test(-32766)
	test(-129)
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
	test(128)
	test(129)
	test(32766)
	test(32767)
}
