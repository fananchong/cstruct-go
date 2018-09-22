package benchmarks

import (
	"fmt"
	"reflect"
	"testing"
)

type TestTypeA struct {
	a int
	b float64
}

func Benchmark_sprintf(b *testing.B) {
	a := TestTypeA{}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%T", a)
	}
}

func Benchmark_reflect(b *testing.B) {
	a := TestTypeA{}
	for i := 0; i < b.N; i++ {
		_ = reflect.TypeOf(a).String()
	}
}
