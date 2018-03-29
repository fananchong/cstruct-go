package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 string) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.String.PackLE(val0)
	b := c.String.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.String.UnpackLE(l)
	val1b := c.String.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test("hello")
	test("")
	test("world")
}
