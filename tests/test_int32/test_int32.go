package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 int32) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.Int32.PackLE(val0)
	b := c.Int32.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.Int32.UnpackLE(l)
	val1b := c.Int32.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test(-2147483648)
	test(-2147483647)
	test(-2147483646)
	test(-8388609)
	test(-8388608)
	test(-8388607)
	test(-8388606)
	test(-32769)
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
	test(32768)
	test(32769)
	test(8388606)
	test(8388607)
	test(8388608)
	test(8388609)
	test(2147483646)
	test(2147483647)
}
