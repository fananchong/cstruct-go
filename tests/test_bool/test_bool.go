package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 bool) {
	fmt.Println("test val0:", val0, "==========================")

	l := c.Bool.PackLE(val0)
	b := c.Bool.PackBE(val0)

	fmt.Println("l:", l)
	fmt.Println("b:", b)

	val1l := c.Bool.UnpackLE(l)
	val1b := c.Bool.UnpackBE(b)

	fmt.Println("val1l:", val1l)
	fmt.Println("val1b:", val1b)

	fmt.Println("")
}

func main() {
	test(true)
	test(false)
}
