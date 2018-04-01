package benchmarks

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
	"github.com/golang/protobuf/proto"
)

type mystruct2 struct {
	F3 float64 `c:"double"`
	F4 string  `c:"string"`
}

type mystruct1 struct {
	F1  bool    `c:"bool"`
	F2  float32 `c:"float"`
	F3  float64 `c:"double"`
	F4  string  `c:"string"`
	F5  []byte  `c:"binary"`
	F6  int8    `c:"int8"`
	F7  int16   `c:"int16"`
	F9  int32   `c:"int32"`
	F11 int64   `c:"int64"`
	F12 uint8   `c:"uint8"`
	F13 uint16  `c:"uint16"`
	F15 uint32  `c:"uint32"`
	F17 uint64  `c:"uint64"`
	S0  mystruct2
}

func Benchmark_CStructGO(b *testing.B) {
	a0 := &mystruct1{}
	a0.F1 = true
	a0.F2 = 0.98
	a0.F3 = 999.777
	a0.F4 = "hellohellohellohellohellohellohellohellohellohello"
	a0.F5 = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a0.F6 = 126
	a0.F7 = 32766
	a0.F9 = -2147483646
	a0.F11 = -9223372036854774807
	a0.F12 = 254
	a0.F13 = 65534
	a0.F15 = 4294967295
	a0.F17 = 18446744073709551614
	a0.S0.F3 = 988.07
	a0.S0.F4 = "world1world1world1world1world1world1world1world1world1"

	for i := 0; i < b.N; i++ {
		buf_l, _ := cstruct.Marshal(a0)
		a1 := &mystruct1{}
		cstruct.Unmarshal(buf_l, a1)
	}
}

func Benchmark_Protobuf(b *testing.B) {
	a0 := &Myproto1{}
	a0.F1 = true
	a0.F2 = 0.98
	a0.F3 = 999.777
	a0.F4 = "hellohellohellohellohellohellohellohellohellohello"
	a0.F5 = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a0.F6 = 126
	a0.F7 = 32766
	a0.F9 = -2147483646
	a0.F11 = -9223372036854774807
	a0.F12 = 254
	a0.F13 = 65534
	a0.F15 = 4294967295
	a0.F17 = 18446744073709551614
	a0.S0 = &Myproto2{}
	a0.S0.F3 = 988.07
	a0.S0.F4 = "world1world1world1world1world1world1world1world1world1"

	for i := 0; i < b.N; i++ {
		buf_l, _ := proto.Marshal(a0)
		a1 := &Myproto1{}
		proto.Unmarshal(buf_l, a1)
	}
}
