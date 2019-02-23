package mytest

import (
	"fmt"
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct111 struct {
	F5 int32
}

func Test_MY2NIL1(t *testing.T) {
	//var a interface{} = (*mystruct111)(nil)
	//a := &mystruct111{}
	a := []int{1, 2, 3}
	bufL, err0 := cstruct.Marshal(a)
	if err0 != nil {
		fmt.Println(err0)
		t.Error("出错啦！#0")
		return
	}
	b := &mystruct111{}
	if err := cstruct.Unmarshal(bufL, b); err != nil {
		fmt.Println(err)
		t.Error("出错啦！#1")
		return
	}
}
