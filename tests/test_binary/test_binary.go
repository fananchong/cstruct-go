package main

import (
	"fmt"

	c "github.com/fananchong/cstruct-go/datatypes"
)

func test(val0 []byte) {
	fmt.Println("test raw val:", val0, "==========================")

	l := c.Binary.PackLE(val0)
	b := c.Binary.PackBE(val0)

	fmt.Println("Buf(l):", l)
	fmt.Println("Buf(b):", b)

	val1l := c.Binary.UnpackLE(l)
	val1b := c.Binary.UnpackBE(b)

	fmt.Println("NewVal(l):", val1l)
	fmt.Println("NewVal(b):", val1b)

	fmt.Println("")
}

func main() {
	test([]byte{1, 2, 3, 4, 5})
	test([]byte("hello"))
	test([]byte(""))
	test([]byte("world"))
}
