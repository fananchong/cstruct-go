package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 uint16) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.UInt16.PackLE(val0)
	b := c.UInt16.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.UInt16.UnpackLE(l)
	val1b := c.UInt16.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test(0)
	test(1)
	test(2)
	test(254)
	test(255)
	test(65534)
	test(65535)
}
