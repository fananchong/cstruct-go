package mytest

import (
	"fmt"
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct11 struct {
	F5  []byte
	S0  *mystruct2
	F18 []bool
	F19 []int8
	F20 []uint8
	F21 []int16
	F22 []uint16
	F23 []int32
	F24 []uint32
	F25 []int64
	F26 []uint64
	F27 []float32
	F28 []float64
	F29 []string
	F30 []*mystruct4
	F31 [][]byte
}

func Test_NIL1(t *testing.T) {
	a := &mystruct11{}
	buf_l, _ := cstruct.Marshal(a)
	b := &mystruct11{}
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		fmt.Println(err)
		t.Error("出错啦！#0")
		return
	}
}

func Test_NIL2(t *testing.T) {
	a := &mystruct11{}
	a.S0 = &mystruct2{}
	a.S0.F3 = 988.07
	a.S0.F4 = "world1"
	a.S0.S1.F5 = []byte("world1")
	a.F30 = []*mystruct4{&mystruct4{}, nil, &mystruct4{}}
	a.F31 = [][]byte{[]byte("hello1"), []byte{}, nil, nil, []byte{}, []byte("world2")}

	buf_l, _ := cstruct.Marshal(a)
	b := &mystruct11{}
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		fmt.Println(err)
		t.Error("出错啦！#0")
		return
	}

	if a.S0.F3 != b.S0.F3 {
		t.Error("出错啦！#18")
		return
	}
	t.Log(b.S0.F3)
	if a.S0.F4 != b.S0.F4 {
		t.Error("出错啦！#19")
		return
	}
	t.Log(b.S0.F4)
	if len(a.S0.S1.F5) != len(b.S0.S1.F5) {
		t.Error("出错啦！#20")
		return
	}
	t.Log(b.S0.S1.F5)
	for i := 0; i < len(a.S0.S1.F5); i++ {
		if a.S0.S1.F5[i] != b.S0.S1.F5[i] {
			t.Error("出错啦！#20")
			return
		}
	}

	if len(a.F30) != len(b.F30) {
		t.Error("出错啦！#33")
		return
	}
	for i := 0; i < len(a.F30); i++ {
		if a.F30[i] != nil && a.F30[i].F6 != b.F30[i].F6 {
			t.Error("出错啦！#33")
			return
		}
		if a.F30[i] != nil && a.F30[i].F7 != b.F30[i].F7 {
			t.Error("出错啦！#33")
			return
		}
	}
	t.Log(b.F30)

	if len(a.F31) != len(b.F31) {
		t.Error("出错啦！#34")
		return
	}
	for i := 0; i < len(a.F31); i++ {
		if string(a.F31[i]) != string(b.F31[i]) {
			t.Error("出错啦！#34")
			return
		}
	}
	t.Log(b.F31)
}
