package cstruct

import (
	"fmt"
	"reflect"
)

type encoder func(p *Buffer, prop *Properties, base structPointer) error
type decoder func(p *Buffer, prop *Properties, base structPointer) error
type sizer func(p *Buffer, prop *Properties, base structPointer) int

type Properties struct {
	Name string
	tag  string

	// 或者是字段
	field field
	enc   encoder
	dec   decoder
	size  sizer

	// 或者是嵌套结构体
	sprop *StructProperties
}

func (p *Properties) Init(typ reflect.Type, name, tag string, f *reflect.StructField) {
	p.init(typ, name, tag, f, true)
}

func (p *Properties) init(typ reflect.Type, name, tag string, f *reflect.StructField, lockGetProp bool) {
	p.Name = name
	if f != nil {
		p.field = toField(f)
	}
	p.Parse(tag)
	p.setEncAndDec(typ, f, lockGetProp)
}

func (p *Properties) Parse(s string) {
	p.tag = s
}

func (p *Properties) setEncAndDec(typ reflect.Type, f *reflect.StructField, lockGetProp bool) {
	p.enc = nil
	p.dec = nil
	p.size = nil

	if typ.Kind() == reflect.Struct {
		if lockGetProp {
			p.sprop = GetProperties(typ)
		} else {
			p.sprop = getPropertiesLocked(typ)
		}
	} else {
		switch p.tag {
		case CTypeBool:
			if CurrentByteOrder == LE {
				p.enc = enc_bool
				p.size = size_bool
			} else {
				//TODO:
			}

		default:
			panic(fmt.Sprintf("unknow type! type = %s", p.tag))
		}
	}
}
