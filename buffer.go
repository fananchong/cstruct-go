package cstruct

import (
	"io"
	"reflect"
	"unsafe"
)

type Buffer struct {
	buf   []byte
	index int
}

func NewBuffer(e []byte) *Buffer {
	return &Buffer{buf: e}
}

// Marshal

func (p *Buffer) Marshal(obj IStruct) error {
	t, base, err := getbase(obj)
	if structPointer_IsNil(base) {
		return ErrNil
	}
	if err == nil {
		err = p.enc_struct(GetProperties(t.Elem()), base)
	}
	return err
}

func getbase(obj IStruct) (t reflect.Type, b structPointer, err error) {
	if obj == nil {
		err = ErrNil
		return
	}
	t = reflect.TypeOf(obj)
	value := reflect.ValueOf(obj)
	b = toStructPointer(value)
	return
}

func (o *Buffer) enc_struct(prop *StructProperties, base structPointer) error {
	for _, p := range prop.Prop {
		if p.enc != nil {
			if err := p.enc(o, p, base); err != nil {
				return err
			}
		}
	}
	return nil
}

// Unmarshal

func (p *Buffer) Unmarshal(obj IStruct) error {
	typ, base, err := getbase(obj)
	if err != nil {
		return err
	}

	return p.unmarshalType(typ.Elem(), GetProperties(typ.Elem()), base)
}

func (o *Buffer) unmarshalType(st reflect.Type, prop *StructProperties, base structPointer) error {
	for _, p := range prop.Prop {
		if p.dec != nil {
			if err := p.dec(o, p, base); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *Buffer) enc_bool(p *Properties, base structPointer) error {
	v := structPointer_BoolVal(base, p.field)
	if v == nil {
		return ErrNil
	}
	x := 0
	if *v {
		x = 1
	}
	o.buf = append(o.buf, uint8(x))
	return nil
}

func (o *Buffer) dec_bool(p *Properties, base structPointer) error {
	i := o.index + 1
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint8(o.buf[i-1])
	*structPointer_BoolVal(base, p.field) = (u != 0)
	return nil
}

func (o *Buffer) size_bool(prop *Properties, base structPointer) int {
	return 1
}

func (o *Buffer) enc_uint8(p *Properties, base structPointer) error {
	v := (*uint8)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	o.buf = append(o.buf, uint8(*v))
	return nil
}

func (o *Buffer) dec_uint8(p *Properties, base structPointer) error {
	i := o.index + 1
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint8(o.buf[i-1])
	*(*uint8)(unsafe.Pointer(uintptr(base) + uintptr(p.field))) = u
	return nil
}

func (o *Buffer) size_uint8(prop *Properties, base structPointer) int {
	return 1
}

func (o *Buffer) enc_uint16le(p *Properties, base structPointer) error {
	v := (*uint16)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	x := *v
	o.buf = append(o.buf, uint8(x), uint8(x>>8))
	return nil
}

func (o *Buffer) enc_uint16be(p *Properties, base structPointer) error {
	v := (*uint16)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	x := *v
	o.buf = append(o.buf, uint8(x>>8), uint8(x))
	return nil
}

func (o *Buffer) dec_uint16le(p *Properties, base structPointer) error {
	i := o.index + 2
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint16(o.buf[i-2])
	u |= uint16(o.buf[i-1]) << 8
	*(*uint16)(unsafe.Pointer(uintptr(base) + uintptr(p.field))) = u
	return nil
}

func (o *Buffer) dec_uint16be(p *Properties, base structPointer) error {
	i := o.index + 2
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint16(o.buf[i-2]) << 8
	u |= uint16(o.buf[i-1])
	*(*uint16)(unsafe.Pointer(uintptr(base) + uintptr(p.field))) = u
	return nil
}

func (o *Buffer) size_uint16(prop *Properties, base structPointer) int {
	return 2
}
