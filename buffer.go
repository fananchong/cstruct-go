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

// bool
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
	v := structPointer_BoolVal(base, p.field)
	if v == nil {
		return ErrNil
	}
	*v = (u != 0)
	return nil
}

// uint8
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
	v := (*uint8)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	*v = u
	return nil
}

// uint16
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
	v := (*uint16)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	*v = u
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
	v := (*uint16)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	*v = u
	return nil
}

// uint32
func (o *Buffer) enc_uint32le(p *Properties, base structPointer) error {
	v := (*uint32)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	x := *v
	o.buf = append(o.buf, uint8(x), uint8(x>>8), uint8(x>>16), uint8(x>>24))
	return nil
}

func (o *Buffer) enc_uint32be(p *Properties, base structPointer) error {
	v := (*uint32)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	x := *v
	o.buf = append(o.buf, uint8(x>>24), uint8(x>>16), uint8(x>>8), uint8(x))
	return nil
}

func (o *Buffer) dec_uint32le(p *Properties, base structPointer) error {
	i := o.index + 4
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint32(o.buf[i-4])
	u |= uint32(o.buf[i-3]) << 8
	u |= uint32(o.buf[i-2]) << 16
	u |= uint32(o.buf[i-1]) << 24
	v := (*uint32)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	*v = u
	return nil
}

func (o *Buffer) dec_uint32be(p *Properties, base structPointer) error {
	i := o.index + 4
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint32(o.buf[i-4]) << 24
	u = uint32(o.buf[i-3]) << 16
	u = uint32(o.buf[i-2]) << 8
	u |= uint32(o.buf[i-1])
	v := (*uint32)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	*v = u
	return nil
}

// uint64
func (o *Buffer) enc_uint64le(p *Properties, base structPointer) error {
	v := (*uint64)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	x := *v
	o.buf = append(o.buf, uint8(x), uint8(x>>8), uint8(x>>16), uint8(x>>24), uint8(x>>32), uint8(x>>40), uint8(x>>48), uint8(x>>56))
	return nil
}

func (o *Buffer) enc_uint64be(p *Properties, base structPointer) error {
	v := (*uint64)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	x := *v
	o.buf = append(o.buf, uint8(x>>56), uint8(x>>48), uint8(x>>40), uint8(x>>32), uint8(x>>24), uint8(x>>16), uint8(x>>8), uint8(x))
	return nil
}

func (o *Buffer) dec_uint64le(p *Properties, base structPointer) error {
	i := o.index + 8
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint64(o.buf[i-8])
	u |= uint64(o.buf[i-7]) << 8
	u |= uint64(o.buf[i-6]) << 16
	u |= uint64(o.buf[i-5]) << 24
	u |= uint64(o.buf[i-4]) << 32
	u |= uint64(o.buf[i-3]) << 40
	u |= uint64(o.buf[i-2]) << 48
	u |= uint64(o.buf[i-1]) << 56
	v := (*uint64)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	*v = u
	return nil
}

func (o *Buffer) dec_uint64be(p *Properties, base structPointer) error {
	i := o.index + 8
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	u := uint64(o.buf[i-8]) << 56
	u = uint64(o.buf[i-7]) << 48
	u = uint64(o.buf[i-6]) << 40
	u = uint64(o.buf[i-5]) << 32
	u = uint64(o.buf[i-4]) << 24
	u = uint64(o.buf[i-3]) << 16
	u = uint64(o.buf[i-2]) << 8
	u |= uint64(o.buf[i-1])
	v := (*uint64)(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if v == nil {
		return ErrNil
	}
	*v = u
	return nil
}

// string
func (o *Buffer) enc_string(p *Properties, base structPointer) error {
	v := structPointer_StringVal(base, p.field)
	if v == nil {
		return ErrNil
	}
	x := len(*v)
	if CurrentByteOrder == LE {
		o.buf = append(o.buf, uint8(x), uint8(x>>8))
	} else {
		o.buf = append(o.buf, uint8(x>>8), uint8(x))
	}

	if x > 0 {
		o.buf = append(o.buf, (*v)...)
	}
	return nil
}

func (o *Buffer) dec_string(p *Properties, base structPointer) error {
	i := o.index + 2
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	var nb uint16 = 0
	if CurrentByteOrder == LE {
		nb = uint16(o.buf[i-2])
		nb |= uint16(o.buf[i-1]) << 8
	} else {
		nb = uint16(o.buf[i-2]) << 8
		nb |= uint16(o.buf[i-1])
	}

	end := o.index + int(nb)
	if end < o.index || end > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	buf := o.buf[o.index:end]
	o.index += int(nb)

	v := structPointer_StringVal(base, p.field)
	if v == nil {
		return ErrNil
	}
	*v = string(buf)
	return nil
}

// binary
func (o *Buffer) enc_binary(p *Properties, base structPointer) error {
	v := structPointer_Bytes(base, p.field)
	if v == nil {
		return ErrNil
	}
	x := len(*v)
	if CurrentByteOrder == LE {
		o.buf = append(o.buf, uint8(x), uint8(x>>8))
	} else {
		o.buf = append(o.buf, uint8(x>>8), uint8(x))
	}

	if x > 0 {
		o.buf = append(o.buf, (*v)...)
	}
	return nil
}

func (o *Buffer) dec_binary(p *Properties, base structPointer) error {
	i := o.index + 2
	if i < 0 || i > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	o.index = i
	var nb uint16 = 0
	if CurrentByteOrder == LE {
		nb = uint16(o.buf[i-2])
		nb |= uint16(o.buf[i-1]) << 8
	} else {
		nb = uint16(o.buf[i-2]) << 8
		nb |= uint16(o.buf[i-1])
	}

	end := o.index + int(nb)
	if end < o.index || end > len(o.buf) {
		return io.ErrUnexpectedEOF
	}
	buf := o.buf[o.index:end]
	o.index += int(nb)

	v := structPointer_Bytes(base, p.field)
	if v == nil {
		return ErrNil
	}
	*v = append(buf)
	return nil
}

// struct
func (o *Buffer) enc_substruct_ptr(p *Properties, base structPointer) error {
	b := structPointer_GetStructPointer(base, p.field)
	if structPointer_IsNil(b) {
		return ErrNil
	}
	return o.enc_struct(p.sprop, b)
}
func (o *Buffer) dec_substruct_ptr(p *Properties, base structPointer) error {
	bas := structPointer_GetStructPointer(base, p.field)
	if structPointer_IsNil(bas) {
		bas = toStructPointer(reflect.New(p.stype))
		structPointer_SetStructPointer(base, p.field, bas)
	}
	return o.unmarshalType(p.stype, p.sprop, bas)
}

func (o *Buffer) enc_substruct(p *Properties, base structPointer) error {
	b := structPointer(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if structPointer_IsNil(b) {
		return ErrNil
	}
	return o.enc_struct(p.sprop, b)
}
func (o *Buffer) dec_substruct(p *Properties, base structPointer) error {
	bas := structPointer(unsafe.Pointer(uintptr(base) + uintptr(p.field)))
	if structPointer_IsNil(bas) {
		return ErrNil
	}
	return o.unmarshalType(p.stype, p.sprop, bas)
}
