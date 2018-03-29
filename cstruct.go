package cstruct

import (
	"errors"
	"fmt"
	"reflect"

	c "github.com/fananchong/cstruct-go/datatypes"
	"github.com/fatih/structs"
)

type ByteOrder int

const (
	_  ByteOrder = iota //0
	LE                  //1
	BE                  //2
)

const (
	Tag = "c"
)

const (
	CTypeBool   = "bool"
	CTypeInt8   = "int8"
	CTypeInt16  = "int16"
	CTypeInt24  = "int24"
	CTypeInt32  = "int32"
	CTypeInt40  = "int40"
	CTypeInt64  = "int64"
	CTypeUInt8  = "uint8"
	CTypeUInt16 = "uint16"
	CTypeUInt24 = "uint24"
	CTypeUInt32 = "uint32"
	CTypeUInt40 = "uint40"
	CTypeUInt64 = "uint64"
	CTypeFloat  = "float"
	CTypeDouble = "double"
	CTypeString = "string"
	CTypeBinary = "binary"
)

func PackLE(obj interface{}) []byte {
	var ret []byte
	for _, field := range structs.Fields(obj) {
		ret = append(ret, packField(field, LE)...)
	}
	return ret
}

func PackBE(obj interface{}) []byte {
	var ret []byte
	for _, field := range structs.Fields(obj) {
		ret = append(ret, packField(field, BE)...)
	}
	return ret
}

func packField(field *structs.Field, order ByteOrder) []byte {
	var ret []byte
	switch field.Kind() {
	case reflect.Struct:
		for _, f := range field.Fields() {
			ret = append(ret, packField(f, order)...)
		}
	default:
		ctype := field.Tag(Tag)
		if ctype != "" {
			switch ctype {
			case CTypeBool:
				if order == LE {
					ret = append(ret, c.Bool.PackLE(field.Value().(bool))...)
				} else {
					ret = append(ret, c.Bool.PackBE(field.Value().(bool))...)
				}
			default:
				panic(fmt.Sprintf("unknow type! type = %s", ctype))
			}
		}
	}
	return ret
}

func UnpackLE(buf []byte, obj interface{}) error {
	pos := 0
	for _, field := range structs.Fields(obj) {
		if err := unpackField(field, buf, &pos, LE); err != nil {
			return err
		}
	}
	return nil
}

func UnpackBE(buf []byte, obj interface{}) error {
	pos := 0
	for _, field := range structs.Fields(obj) {
		if err := unpackField(field, buf, &pos, BE); err != nil {
			return err
		}
	}
	return nil
}

func unpackField(field *structs.Field, buf []byte, pos *int, order ByteOrder) error {
	switch field.Kind() {
	case reflect.Struct:
		for _, f := range field.Fields() {
			if err := unpackField(f, buf, pos, order); err != nil {
				return err
			}
		}
	default:
		ctype := field.Tag(Tag)
		if ctype != "" {
			switch ctype {
			case CTypeBool:
				if order == LE {
					field.Set(c.Bool.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Bool.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Bool.Size()
			default:
				return errors.New(fmt.Sprintf("unknow type! type = %s", ctype))
			}
		}
	}
	return nil
}
