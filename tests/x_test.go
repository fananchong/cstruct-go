package mytest

import (
	"fmt"
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct44 struct {
	F37 uint16
	F38 [5]int32
}

type mystruct4 struct {
	F6 int8
	F7 int16
}

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
	S1  mystruct44
	S2  mystruct44
	S3  mystruct44
	F15 uint32
	F17 uint64
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
	F32 [5]int8
	F33 [6]byte
	F34 [7]uint8
	F35 [4]bool
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
	a.F13 = 65535
	a.F15 = 4294967295
	a.F17 = 18446744073709551615
	a.F18 = []bool{false, true, true, false, true}
	a.F19 = []int8{1, -1, 0, 127, -128}
	a.F20 = []uint8{0, 1, 2, 127, 255}
	a.F21 = []int16{1, -1, 0, 32767, -32768}
	a.F22 = []uint16{0, 1, 2, 32767, 65535}
	a.F23 = []int32{1, -1, 0, 2147483647, -2147483648}
	a.F24 = []uint32{0, 1, 2, 2147483647, 4294967295}
	a.F25 = []int64{1, -1, 0, 9223372036854775807, -9223372036854775808}
	a.F26 = []uint64{0, 1, 2, 9223372036854775807, 18446744073709551615}
	a.F27 = []float32{0.98, -1, 0, -0.98}
	a.F28 = []float64{0, 999888888.777, 2, -999888888.777}
	a.F29 = []string{"hello", "", "world", "", "123"}
	a.F30 = []*mystruct4{&mystruct4{1, 2}, &mystruct4{10, 20}, &mystruct4{100, 200}, &mystruct4{11, 21}}
	a.F31 = [][]byte{[]byte("hello1"), []byte{}, []byte("world1"), []byte("hello2"), []byte{}, []byte("world2")}
	a.F32 = [5]int8{0, 1, 2, 3, 4}
	a.F33 = [6]byte{'h', 'e', 'l', 'l', 'o', '1'}
	a.F34 = [7]uint8{0, 1, 2, 3, 4, 5, 6}
	a.F35 = [4]bool{true, false, false, true}

	a.S0.F3 = 988.07
	a.S0.F4 = "world1"
	a.S0.S1.F5 = []byte("world1")

	a.S1.F37 = 65535
	a.S1.F38 = [5]int32{1, -1, 0, 2147483647, -2147483648}
	a.S2.F37 = 1
	a.S2.F38 = [5]int32{1, -1, 0, 1, -2}
	a.S3.F37 = 65535
	a.S3.F38 = [5]int32{3, -3, 0, 3, -3}

	test1(t, a)
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
	a.F21 = []int16{}
	a.F22 = []uint16{}
	a.F23 = []int32{}
	a.F24 = []uint32{}
	a.F25 = []int64{}
	a.F26 = []uint64{}
	a.F27 = []float32{}
	a.F28 = []float64{}
	a.F29 = []string{}
	a.F30 = []*mystruct4{}
	a.F31 = [][]byte{}

	a.S0.F3 = 988.07
	a.S0.F4 = "world2"
	a.S0.S1.F5 = []byte("world2")

	test1(t, a)
}

func test1(t *testing.T, a *mystruct1) {
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

	if len(a.F21) != len(b.F21) {
		t.Error("出错啦！#24")
		return
	}
	for i := 0; i < len(a.F21); i++ {
		if a.F21[i] != b.F21[i] {
			t.Error("出错啦！#24")
			return
		}
	}
	t.Log(b.F21)

	if len(a.F22) != len(b.F22) {
		t.Error("出错啦！#25")
		return
	}
	for i := 0; i < len(a.F22); i++ {
		if a.F22[i] != b.F22[i] {
			t.Error("出错啦！#25")
			return
		}
	}
	t.Log(b.F22)

	if len(a.F23) != len(b.F23) {
		t.Error("出错啦！#26")
		return
	}
	for i := 0; i < len(a.F23); i++ {
		if a.F23[i] != b.F23[i] {
			t.Error("出错啦！#26")
			return
		}
	}
	t.Log(b.F23)

	if len(a.F24) != len(b.F24) {
		t.Error("出错啦！#27")
		return
	}
	for i := 0; i < len(a.F24); i++ {
		if a.F24[i] != b.F24[i] {
			t.Error("出错啦！#27")
			return
		}
	}
	t.Log(b.F24)

	if len(a.F25) != len(b.F25) {
		t.Error("出错啦！#28")
		return
	}
	for i := 0; i < len(a.F25); i++ {
		if a.F25[i] != b.F25[i] {
			t.Error("出错啦！#28")
			return
		}
	}
	t.Log(b.F25)

	if len(a.F26) != len(b.F26) {
		t.Error("出错啦！#29")
		return
	}
	for i := 0; i < len(a.F26); i++ {
		if a.F26[i] != b.F26[i] {
			t.Error("出错啦！#29")
			return
		}
	}
	t.Log(b.F26)

	if len(a.F27) != len(b.F27) {
		t.Error("出错啦！#30")
		return
	}
	for i := 0; i < len(a.F27); i++ {
		if a.F27[i] != b.F27[i] {
			t.Error("出错啦！#30")
			return
		}
	}
	t.Log(b.F27)

	if len(a.F28) != len(b.F28) {
		t.Error("出错啦！#31")
		return
	}
	for i := 0; i < len(a.F28); i++ {
		if a.F28[i] != b.F28[i] {
			t.Error("出错啦！#31")
			return
		}
	}
	t.Log(b.F28)

	if len(a.F29) != len(b.F29) {
		t.Error("出错啦！#32")
		return
	}
	for i := 0; i < len(a.F29); i++ {
		if a.F29[i] != b.F29[i] {
			t.Error("出错啦！#32")
			return
		}
	}
	t.Log(b.F29)

	if len(a.F30) != len(b.F30) {
		t.Error("出错啦！#33")
		return
	}
	for i := 0; i < len(a.F30); i++ {
		if a.F30[i].F6 != b.F30[i].F6 {
			t.Error("出错啦！#33")
			return
		}
		if a.F30[i].F7 != b.F30[i].F7 {
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
	t.Log(b.F32)
	t.Log(b.F33)
	t.Log(b.F34)
	t.Log(b.F35)
	t.Log(b.S1)
	t.Log(b.S2)
	t.Log(b.S3)
}
