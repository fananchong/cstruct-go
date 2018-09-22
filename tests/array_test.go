package mytest

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct5 struct {
	F32 [5]int8
	F33 [6]byte
	F34 [7]uint8
	F35 [4]bool
	F36 [5]int16
	F37 [5]uint16
}

func Test_LE1(t *testing.T) {
	a := &mystruct5{}
	a.F32 = [5]int8{0, -128, 2, 127, 4}
	a.F33 = [6]byte{'h', 'e', 'l', 'l', 'o', '1'}
	a.F34 = [7]uint8{0, 1, 2, 3, 255, 5, 6}
	a.F35 = [4]bool{true, false, false, true}
	a.F36 = [5]int16{1, -1, 0, 32767, -32768}
	a.F37 = [5]uint16{0, 1, 2, 32767, 65535}
	test1(t, a)
}

func test1(t *testing.T, a *mystruct5) {
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
}
