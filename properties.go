package cstruct

import (
	"reflect"
	"sync"
)

// StructProperties

type StructProperties struct {
	Prop []*Properties
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
		p.init(f.Type, name, "", &f)
		prop.Prop[i] = p
	}

	return prop
}

// Properties

type encoder func(p *Buffer, prop *Properties, base structPointer) error
type decoder func(p *Buffer, prop *Properties, base structPointer) error

type Properties struct {
	Name  string
	field field
	enc   encoder
	dec   decoder
	stype reflect.Type
	sprop *StructProperties
}

func (p *Properties) init(typ reflect.Type, name, tag string, f *reflect.StructField) {
	p.Name = name
	if f != nil {
		p.field = toField(f)
	}
	p.setEncAndDec(typ, f)
}

func (p *Properties) setEncAndDec(typ reflect.Type, f *reflect.StructField) {
	p.enc = nil
	p.dec = nil

	switch typ.Kind() {
	case reflect.Bool: // bool
		p.enc = (*Buffer).enc_bool
		p.dec = (*Buffer).dec_bool
	case reflect.Int8, reflect.Uint8: // int8 uint8
		p.enc = (*Buffer).enc_uint8
		p.dec = (*Buffer).dec_uint8
	case reflect.Int16, reflect.Uint16: // int16 uint16
		p.enc = (*Buffer).enc_uint16
		p.dec = (*Buffer).dec_uint16
	case reflect.Int32, reflect.Uint32, reflect.Float32: // int32 uint32 float32
		if CurrentByteOrder == LE {
			p.enc = (*Buffer).enc_uint32le
			p.dec = (*Buffer).dec_uint32le
		} else {
			p.enc = (*Buffer).enc_uint32be
			p.dec = (*Buffer).dec_uint32be
		}
	case reflect.Int64, reflect.Uint64, reflect.Float64: // int64 uint64 float64
		if CurrentByteOrder == LE {
			p.enc = (*Buffer).enc_uint64le
			p.dec = (*Buffer).dec_uint64le
		} else {
			p.enc = (*Buffer).enc_uint64be
			p.dec = (*Buffer).dec_uint64be
		}
	case reflect.String: // string
		p.enc = (*Buffer).enc_string
		p.dec = (*Buffer).dec_string
	case reflect.Ptr: // struct ptr
		if t2 := typ.Elem(); t2.Kind() == reflect.Struct {
			p.stype = t2
			p.sprop = getPropertiesLocked(p.stype)
			p.enc = (*Buffer).enc_substruct_ptr
			p.dec = (*Buffer).dec_substruct_ptr
		} else {
			panic("cstruct: unknow type. field name =" + f.Name)
		}
	case reflect.Struct: // struct
		p.stype = typ
		p.sprop = getPropertiesLocked(p.stype)
		p.enc = (*Buffer).enc_substruct
		p.dec = (*Buffer).dec_substruct
	case reflect.Slice:
		switch t2 := typ.Elem(); t2.Kind() {
		case reflect.Uint8: // []byte
			p.enc = (*Buffer).enc_slice_byte
			p.dec = (*Buffer).dec_slice_byte
		case reflect.Bool: // []bool
			p.enc = (*Buffer).enc_slice_bool
			p.dec = (*Buffer).dec_slice_bool
		}
	default:
		panic("cstruct: unknow type. field name = " + f.Name)
	}
}
