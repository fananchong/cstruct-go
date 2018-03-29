package cstruct

import (
	"reflect"

	c "github.com/fananchong/cstruct-go/datatypes"
	"github.com/fatih/structs"
)

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
				panic("unknow type!")
			}
		}
	}
	return ret
}

func UnpackLE([]byte, interface{}) error {

}

func UnpackBE([]byte, interface{}) error {

}
