package mytest

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct6 struct {
	F37 uint16
	F38 [5]int32
}

type mystruct5 struct {
	F32 [5]int8
	F33 [6]byte
	F34 [7]uint8
	F35 [4]bool
	F36 [5]int16
	F37 [5]uint16
	F38 [5]int32
	F39 [5]uint32
	F40 [5]float32
	F41 [5]int64
	F42 [5]uint64
	F43 [5]float64
	F44 [3]mystruct6
}

func Test_LE111(t *testing.T) {
	a := &mystruct5{}
	a.F32 = [5]int8{0, -128, 2, 127, 4}
	a.F33 = [6]byte{'h', 'e', 'l', 'l', 'o', '1'}
	a.F34 = [7]uint8{0, 1, 2, 3, 255, 5, 6}
	a.F35 = [4]bool{true, false, false, true}
	a.F36 = [5]int16{1, -1, 0, 32767, -32768}
	a.F37 = [5]uint16{0, 1, 2, 32767, 65535}
	a.F38 = [5]int32{1, -1, 0, 2147483647, -2147483648}
	a.F39 = [5]uint32{0, 1, 2, 2147483647, 4294967295}
	a.F40 = [5]float32{0.98, -1, 0, -0.98, 999.9}
	a.F41 = [5]int64{1, -1, 0, 9223372036854775807, -9223372036854775808}
	a.F42 = [5]uint64{0, 1, 2, 9223372036854775807, 18446744073709551615}
	a.F43 = [5]float64{0, 999888888.777, 2, -999888888.777, 99999999.99}

	b1 := mystruct6{}
	b1.F37 = 65535
	b1.F38 = [5]int32{1, -1, 0, 2147483647, -2147483648}
	b2 := mystruct6{}
	b2.F37 = 1
	b2.F38 = [5]int32{1, -1, 0, 1, -2}
	b3 := mystruct6{}
	b3.F37 = 65535
	b3.F38 = [5]int32{3, -3, 0, 3, -3}
	a.F44 = [3]mystruct6{b1, b2, b3}

	test111(t, a)
}

func test111(t *testing.T, a *mystruct5) {
	buf_l, _ := cstruct.Marshal(a)
	b := &mystruct5{}
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		t.Log(err)
		t.Error("出错啦！")
		return
	}

	t.Log(b.F32)
	t.Log(b.F33)
	t.Log(b.F34)
	t.Log(b.F35)
	t.Log(b.F36)
	t.Log(b.F37)
	t.Log(b.F38)
	t.Log(b.F39)
	t.Log(b.F40)
	t.Log(b.F41)
	t.Log(b.F42)
	t.Log(b.F43)
	t.Log(b.F44)
}
