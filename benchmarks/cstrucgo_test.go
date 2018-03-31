package benchmark1

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct1 struct {
	F1  bool    `c:"bool"`
	F2  float32 `c:"float"`
	F3  float64 `c:"double"`
	F4  string  `c:"string"`
	F5  []byte  `c:"binary"`
	F6  int8    `c:"int8"`
	F7  int16   `c:"int16"`
	F8  int32   `c:"int24"`
	F9  int32   `c:"int32"`
	F10 int64   `c:"int40"`
	F11 int64   `c:"int64"`
	F12 uint8   `c:"uint8"`
	F13 uint16  `c:"uint16"`
	F14 uint32  `c:"uint24"`
	F15 uint32  `c:"uint32"`
	F16 uint64  `c:"uint40"`
	F17 uint64  `c:"uint64"`
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
	a0.F8 = -8388608
	a0.F9 = -2147483646
	a0.F10 = 549755813887
	a0.F11 = -9223372036854774807
	a0.F12 = 254
	a0.F13 = 65534
	a0.F14 = 16777214
	a0.F15 = 4294967295
	a0.F16 = 1099511627770
	a0.F17 = 18446744073709551614

	for i := 0; i < b.N; i++ {
		buf_l := cstruct.PackLE(a0)
		a1 := &mystruct1{}
		cstruct.UnpackLE(buf_l, a1)
	}
}
