package mytest

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct_sint12 struct {
	F1 int8
}

type mystruct_sint10 struct {
	F2 []*mystruct_sint12
	F1 int32
	F3 []*mystruct_sint12
}

type mystruct_sint9 struct {
	F90 []*mystruct_sint10
}

func Test_LE_sint9(t *testing.T) {
	cstruct.OptionSliceIgnoreNil = true
	b1 := []*mystruct_sint12{&mystruct_sint12{1}, &mystruct_sint12{2}, &mystruct_sint12{3}, &mystruct_sint12{4}}
	b2 := []*mystruct_sint12{}
	b3 := []*mystruct_sint12{&mystruct_sint12{1}, &mystruct_sint12{2}}
	{
		a := &mystruct_sint9{}
		a.F90 = []*mystruct_sint10{&mystruct_sint10{nil, 1, nil}, &mystruct_sint10{nil, 2, nil}, &mystruct_sint10{b3, 3, b1}}
		test7(t, a)
	}
	{
		a := &mystruct_sint9{}
		a.F90 = []*mystruct_sint10{&mystruct_sint10{b1, 1, b3}}
		test7(t, a)
	}
	{
		a := &mystruct_sint9{}
		a.F90 = []*mystruct_sint10{}
		test7(t, a)
	}
	{
		a := &mystruct_sint9{}
		a.F90 = []*mystruct_sint10{&mystruct_sint10{b1, 1, b2}}
		test7(t, a)

	}
	{
		a := &mystruct_sint9{}
		a.F90 = []*mystruct_sint10{&mystruct_sint10{b2, 1, b1}}
		test7(t, a)
	}
	{
		a := &mystruct_sint9{}
		a.F90 = []*mystruct_sint10{&mystruct_sint10{b1, 1, b2}, &mystruct_sint10{b2, 1, b1}, &mystruct_sint10{b2, 1, b1}, &mystruct_sint10{b2, 1, b1}}
		test7(t, a)

	}
}

func test7(t *testing.T, a *mystruct_sint9) {
	buf_l, _ := cstruct.Marshal(a)
	b := &mystruct_sint9{}
	t.Log(buf_l)
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		t.Log(err)
		t.Error("出错啦！")
		return
	}
	t.Log(b)
}
