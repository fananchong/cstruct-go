package mytest

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct8 struct {
	F1 int32
}

type mystruct7 struct {
	F70 []*mystruct8
}

func Test_LE7(t *testing.T) {
	cstruct.OptionSliceIgnoreNil = true
	{
		a := &mystruct7{}
		a.F70 = []*mystruct8{&mystruct8{1}, nil, &mystruct8{3}, nil}
		test7(t, a)
	}
	{
		a := &mystruct7{}
		a.F70 = []*mystruct8{}
		test7(t, a)
	}
	{
		a := &mystruct7{}
		a.F70 = []*mystruct8{nil, nil}
		test7(t, a)
	}
	{
		a := &mystruct7{}
		a.F70 = []*mystruct8{&mystruct8{1}, &mystruct8{3}}
		test7(t, a)
	}
}

func test7(t *testing.T, a *mystruct7) {
	buf_l, _ := cstruct.Marshal(a)
	b := &mystruct7{}
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		t.Log(err)
		t.Error("出错啦！")
		return
	}
	t.Log(b)
}
