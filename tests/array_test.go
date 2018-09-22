package mytest

import (
	"reflect"
	"testing"
	"unsafe"
)

func Test_Array(t *testing.T) {
	var a [5]int = [5]int{0, 1, 2, 3, 4}

	var data []int
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	sliceHeader.Cap = 5
	sliceHeader.Len = 5
	sliceHeader.Data = uintptr(unsafe.Pointer(&a))
	for i := 0; i < 5; i++ {
		t.Log(data[i])
	}
}
