package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 uint64) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.UInt40.PackLE(val0)
	b := c.UInt40.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.UInt40.UnpackLE(l)
	val1b := c.UInt40.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test(0)
	test(1)
	test(2)
	test(255)
	test(256)
	test(257)
	test(65535)
	test(65536)
	test(65537)
	test(16777214)
	test(16777215)
	test(16777216)
	test(16777217)
	test(4294967294)
	test(4294967295)
	test(4294967296)
	test(4294967297)
	test(549755813885)
	test(549755813886)
	test(549755813887)
	test(1099511627774)
	test(1099511627775)
	test(1099511627776)
	test(1099511627777)
}
