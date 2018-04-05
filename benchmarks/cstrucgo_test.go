package benchmarks

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
	"github.com/golang/protobuf/proto"
)

type mystruct2 struct {
	F3 float64
	F4 string
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
	F15 uint32
	F17 uint64
	S0  mystruct2
	F29 []string
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
	a0.F29 = []string{"hello", "", "world", "", "123"}

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
	a0.F29 = []string{"hello", "", "world", "", "123"}

	for i := 0; i < b.N; i++ {
		buf_l, _ := proto.Marshal(a0)
		a1 := &Myproto1{}
		proto.Unmarshal(buf_l, a1)
	}
}

func Benchmark_GoGoProtobuf(b *testing.B) {
	a0 := &Myproto3{}
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
	a0.S0 = &Myproto4{}
	a0.S0.F3 = 988.07
	a0.S0.F4 = "world1world1world1world1world1world1world1world1world1"
	a0.F29 = []string{"hello", "", "world", "", "123"}

	for i := 0; i < b.N; i++ {
		buf_l, _ := proto.Marshal(a0)
		a1 := &Myproto3{}
		proto.Unmarshal(buf_l, a1)
	}
}
