package mytest

import (
	"testing"

	cstruct "github.com/fananchong/cstruct-go"
)

type mystruct12 struct {
	F1 int8
}

type mystruct10 struct {
	F2 []mystruct12
	F1 int32
	F3 []mystruct12
}

type mystruct9 struct {
	F90 []mystruct10
}

func Test_LE9(t *testing.T) {
	b1 := []mystruct12{mystruct12{1}, mystruct12{2}, mystruct12{3}, mystruct12{4}}
	b2 := []mystruct12{}
	b3 := []mystruct12{mystruct12{1}, mystruct12{2}}
	{
		a := &mystruct9{}
		a.F90 = []mystruct10{mystruct10{nil, 1, nil}, mystruct10{nil, 2, nil}, mystruct10{b3, 3, b1}}
		test9(t, a)
	}
	{
		a := &mystruct9{}
		a.F90 = []mystruct10{mystruct10{b1, 1, b3}}
		test9(t, a)
	}
	{
		a := &mystruct9{}
		a.F90 = []mystruct10{}
		test9(t, a)
	}
	{
		a := &mystruct9{}
		a.F90 = []mystruct10{mystruct10{b1, 1, b2}}
		test9(t, a)

	}
	{
		a := &mystruct9{}
		a.F90 = []mystruct10{mystruct10{b2, 1, b1}}
		test9(t, a)
	}
	{
		a := &mystruct9{}
		a.F90 = []mystruct10{mystruct10{b1, 1, b2}, mystruct10{b2, 1, b1}, mystruct10{b2, 1, b1}, mystruct10{b2, 1, b1}}
		test9(t, a)

	}
}

func test9(t *testing.T, a *mystruct9) {
	buf_l, _ := cstruct.Marshal(a)
	b := &mystruct9{}
	t.Log(buf_l)
	if err := cstruct.Unmarshal(buf_l, b); err != nil {
		t.Log(err)
		t.Error("出错啦！")
		return
	}
	t.Log(b)
}
