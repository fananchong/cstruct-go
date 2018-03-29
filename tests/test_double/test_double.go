package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 float64) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.Double.PackLE(val0)
	b := c.Double.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.Double.UnpackLE(l)
	val1b := c.Double.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test(0.333)
	test(1.0)
	test(2)
	test(99.998)
	test(990000000.998)
}
