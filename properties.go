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

type Properties struct {
	Name string
	tag  int

	// 或者是字段
	field field
	enc   encoder
	dec   decoder

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
	switch s {
	case CTypeBool:
		p.tag = 1
	case CTypeInt8:
		p.tag = 2
	case CTypeUInt8:
		p.tag = 2
	case CTypeInt16:
		p.tag = 3
	case CTypeUInt16:
		p.tag = 4
	case CTypeInt32:
		p.tag = 4
	case CTypeUInt32:
		p.tag = 4
	case CTypeInt64:
		p.tag = 5
	case CTypeUInt64:
		p.tag = 5
	case CTypeFloat:
		p.tag = 4
	case CTypeDouble:
		p.tag = 5
	case CTypeString:
		p.tag = 6
	case CTypeBinary:
		p.tag = 7
	default:
		panic(fmt.Sprintf("unknow type! type = %s", s))
	}
}

func (p *Properties) setEncAndDec(typ reflect.Type, f *reflect.StructField, lockGetProp bool) {
	p.enc = nil
	p.dec = nil

	if typ.Kind() == reflect.Struct {
		if lockGetProp {
			p.sprop = GetProperties(typ)
		} else {
			p.sprop = getPropertiesLocked(typ)
		}
	} else {
		switch p.tag {
		case 1:
			p.enc = (*Buffer).enc_bool
			p.dec = (*Buffer).dec_bool
		case 2:
			p.enc = (*Buffer).enc_uint8
			p.dec = (*Buffer).dec_uint8
		case 3:
			if CurrentByteOrder == LE {
				p.enc = (*Buffer).enc_uint16le
				p.dec = (*Buffer).dec_uint16le
			} else {
				p.enc = (*Buffer).enc_uint16be
				p.dec = (*Buffer).dec_uint16be
			}
		case 4:
			if CurrentByteOrder == LE {
				p.enc = (*Buffer).enc_uint32le
				p.dec = (*Buffer).dec_uint32le
			} else {
				p.enc = (*Buffer).enc_uint32be
				p.dec = (*Buffer).dec_uint32be
			}
		case 5:
			if CurrentByteOrder == LE {
				p.enc = (*Buffer).enc_uint64le
				p.dec = (*Buffer).dec_uint64le
			} else {
				p.enc = (*Buffer).enc_uint64be
				p.dec = (*Buffer).dec_uint64be
			}
		case 6:
			p.enc = (*Buffer).enc_string
			p.dec = (*Buffer).dec_string
		case 7:
			p.enc = (*Buffer).enc_binary
			p.dec = (*Buffer).dec_binary
		default:
			panic(fmt.Sprintf("unknow type! type = %d", p.tag))
		}
	}
}
