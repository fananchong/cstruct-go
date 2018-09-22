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
	t := ""
	for i := 0; i < b.N; i++ {
		t = fmt.Sprintf("%T", a)
	}
	b.Log(t)
}

func Benchmark_reflect(b *testing.B) {
	a := TestTypeA{}
	t := ""
	for i := 0; i < b.N; i++ {
		t = reflect.TypeOf(a).String()
	}
	b.Log(t)
}
