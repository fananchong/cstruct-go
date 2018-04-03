package mytest

import (
	"fmt"
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct3 struct {
	F5 []byte
}

type mystruct2 struct {
	F3 float64
	F4 string
	S1 mystruct3
}

type mystruct1 struct {
	F1  bool
	F2  float32
	F3  float64
	F4  string
	F5  []byte
	F6  int8
	F7  int16
	F9  int32
	F11 int64
	F12 uint8
	F13 uint16
	S0  *mystruct2
	F15 uint32
	F17 uint64
	F18 []bool
	F19 []int8
	F20 []uint8
}

func Test_LE1(t *testing.T) {
	a := &mystruct1{S0: &mystruct2{}}
	a.F1 = true
	a.F2 = 0.98
	a.F3 = 999888888.777
	a.F4 = "hello1hello2hello3hello4hello5hello6hello7hello8hello9hello0"
	a.F5 = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a.F6 = -128
	a.F7 = -32768
	a.F9 = -2147483648
	a.F11 = -9223372036854775808
	a.F12 = 255
	a.F13 = 32767
	a.F15 = 4294967295
	a.F17 = 18446744073709551615
	a.F18 = []bool{false, true, true, false, true}
	a.F19 = []int8{1, -1, 0, 127, -128}
	a.F20 = []uint8{0, 1, 2, 127, 255}

	a.S0.F3 = 988.07
	a.S0.F4 = "world1"
	a.S0.S1.F5 = []byte("world1")

	test1(t, a, cstruct.LE)
}

func Test_LE2(t *testing.T) {
	a := &mystruct1{S0: &mystruct2{}}
	a.F1 = false
	a.F2 = -0.98
	a.F3 = -999888888.777
	a.F4 = ""
	a.F5 = []byte{}
	a.F6 = 127
	a.F7 = 32767
	a.F9 = 2147483647
	a.F11 = 9223372036854775807
	a.F12 = 1
	a.F13 = 1
	a.F15 = 1
	a.F17 = 1
	a.F18 = []bool{}
	a.F19 = []int8{}
	a.F20 = []uint8{}

	a.S0.F3 = 988.07
	a.S0.F4 = "world2"
	a.S0.S1.F5 = []byte("world2")

	test1(t, a, cstruct.LE)
}

func Test_BE1(t *testing.T) {
	a := &mystruct1{S0: &mystruct2{}}
	a.F1 = true
	a.F2 = 0.98
	a.F3 = 999888888.777
	a.F4 = "hello1hello2hello3hello4hello5hello6hello7hello8hello9hello0"
	a.F5 = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a.F6 = -128
	a.F7 = -32768
	a.F9 = -2147483648
	a.F11 = -9223372036854775808
	a.F12 = 255
	a.F13 = 32767
	a.F15 = 4294967295
	a.F17 = 18446744073709551615
	a.F18 = []bool{false, true, true, false, true}
	a.F19 = []int8{1, -1, 0, 127, -128}
	a.F20 = []uint8{0, 1, 2, 127, 255}

	a.S0.F3 = 988.07
	a.S0.F4 = "world3"
	a.S0.S1.F5 = []byte("world3")

	test1(t, a, cstruct.BE)
}

func Test_BE2(t *testing.T) {
	a := &mystruct1{S0: &mystruct2{}}
	a.F1 = false
	a.F2 = -0.98
	a.F3 = -999888888.777
	a.F4 = ""
	a.F5 = []byte{}
	a.F6 = 127
	a.F7 = 32767
	a.F9 = 2147483647
	a.F11 = 9223372036854775807
	a.F12 = 1
	a.F13 = 1
	a.F15 = 1
	a.F17 = 1
	a.F18 = []bool{}
	a.F19 = []int8{}
	a.F20 = []uint8{}

	a.S0.F3 = 988.07
	a.S0.F4 = "world4"
	a.S0.S1.F5 = []byte("world4")

	test1(t, a, cstruct.BE)
}

func test1(t *testing.T, a *mystruct1, order cstruct.ByteOrder) {
	cstruct.CurrentByteOrder = order
	buf_l, _ := cstruct.Marshal(a)
	b := &mystruct1{S0: &mystruct2{}}
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		fmt.Println(err)
		t.Error("出错啦！#0")
		return
	}
	if a.F1 != b.F1 {
		t.Error("出错啦！#1")
		return
	}
	t.Log(b.F1)
	if a.F2 != b.F2 {
		t.Error("出错啦！#2")
		return
	}
	t.Log(b.F2)
	if a.F3 != b.F3 {
		t.Error("出错啦！#3")
		return
	}
	t.Log(b.F3)
	if a.F4 != b.F4 {
		t.Error("出错啦！#4")
		return
	}
	t.Log(b.F4)
	if len(a.F5) != len(b.F5) {
		t.Error("出错啦！#5")
		return
	}
	t.Log(b.F5)
	for i := 0; i < len(a.F5); i++ {
		if a.F5[i] != a.F5[i] {
			t.Error("出错啦！#5")
			return
		}
	}
	if a.F6 != b.F6 {
		t.Error("出错啦！#6")
		return
	}
	t.Log(b.F6)
	if a.F7 != b.F7 {
		t.Error("出错啦！#7")
		return
	}
	t.Log(b.F7)
	if a.F9 != b.F9 {
		t.Error("出错啦！#9")
		return
	}
	t.Log(b.F9)
	if a.F11 != b.F11 {
		t.Error("出错啦！#11")
		return
	}
	t.Log(b.F11)
	if a.F12 != b.F12 {
		t.Error("出错啦！#12")
		return
	}
	t.Log(b.F12)
	if a.F13 != b.F13 {
		t.Error("出错啦！#13")
		return
	}
	t.Log(b.F13)
	if a.F15 != b.F15 {
		t.Error("出错啦！#15")
		return
	}
	t.Log(b.F15)
	if a.F17 != b.F17 {
		t.Error("出错啦！#17")
		return
	}
	t.Log(b.F17)
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

	if len(a.F18) != len(b.F18) {
		t.Error("出错啦！#21")
		return
	}
	for i := 0; i < len(a.F18); i++ {
		if a.F18[i] != b.F18[i] {
			t.Error("出错啦！#21")
			return
		}
	}
	t.Log(b.F18)

	if len(a.F19) != len(b.F19) {
		t.Error("出错啦！#22")
		return
	}
	for i := 0; i < len(a.F19); i++ {
		if a.F19[i] != b.F19[i] {
			t.Error("出错啦！#22")
			return
		}
	}
	t.Log(b.F19)

	if len(a.F20) != len(b.F20) {
		t.Error("出错啦！#23")
		return
	}
	for i := 0; i < len(a.F20); i++ {
		if a.F20[i] != b.F20[i] {
			t.Error("出错啦！#23")
			return
		}
	}
	t.Log(b.F20)
}
