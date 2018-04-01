package cstruct

import (
	"fmt"
	"os"
	"reflect"
	"sync"
)

// StructProperties

type StructProperties struct {
	Prop  []*Properties
	stype reflect.Type
}

var (
	propertiesMu  sync.RWMutex
	propertiesMap = make(map[reflect.Type]*StructProperties)
)

func GetProperties(t reflect.Type) *StructProperties {
	if t.Kind() != reflect.Struct {
		panic("cstruct: type must have kind struct")
	}

	propertiesMu.RLock()
	sprop, ok := propertiesMap[t]
	propertiesMu.RUnlock()
	if ok {
		return sprop
	}

	propertiesMu.Lock()
	sprop = getPropertiesLocked(t)
	propertiesMu.Unlock()
	return sprop
}

func getPropertiesLocked(t reflect.Type) *StructProperties {
	if prop, ok := propertiesMap[t]; ok {
		return prop
	}

	prop := new(StructProperties)
	propertiesMap[t] = prop
	prop.Prop = make([]*Properties, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		p := new(Properties)
		name := f.Name
		p.init(f.Type, name, f.Tag.Get("c"), &f, false)
		prop.Prop[i] = p
		if p.enc == nil {
			fmt.Fprintln(os.Stderr, "cstruct: no encoder for", f.Name, f.Type.String(), "[GetProperties]")
		}
	}

	return prop
}

// Properties

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
			p.enc = (*Buffer).enc_bool
			p.dec = (*Buffer).dec_bool
			p.size = (*Buffer).size_bool
		case CTypeInt8:
			p.enc = (*Buffer).enc_uint8
			p.dec = (*Buffer).dec_uint8
			p.size = (*Buffer).size_uint8
		case CTypeUInt8:
			p.enc = (*Buffer).enc_uint8
			p.dec = (*Buffer).dec_uint8
			p.size = (*Buffer).size_uint8
		case CTypeInt16:
			if CurrentByteOrder == LE {
				p.enc = (*Buffer).enc_uint16le
				p.dec = (*Buffer).dec_uint16le
				p.size = (*Buffer).size_uint16
			} else {
				p.enc = (*Buffer).enc_uint16be
				p.dec = (*Buffer).dec_uint16be
				p.size = (*Buffer).size_uint16
			}
		case CTypeUInt16:
			if CurrentByteOrder == LE {
				p.enc = (*Buffer).enc_uint16le
				p.dec = (*Buffer).dec_uint16le
				p.size = (*Buffer).size_uint16
			} else {
				p.enc = (*Buffer).enc_uint16be
				p.dec = (*Buffer).dec_uint16be
				p.size = (*Buffer).size_uint16
			}
		default:
			panic(fmt.Sprintf("unknow type! type = %s", p.tag))
		}
	}
}
