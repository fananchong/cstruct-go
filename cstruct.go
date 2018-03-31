package cstruct

import (
	"errors"
	"fmt"
	"reflect"

	c "github.com/fananchong/cstruct-go/datatypes"
	"github.com/fatih/structs"
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
			case CTypeInt8:
				if order == LE {
					ret = append(ret, c.Int8.PackLE(field.Value().(int8))...)
				} else {
					ret = append(ret, c.Int8.PackBE(field.Value().(int8))...)
				}
			case CTypeInt16:
				if order == LE {
					ret = append(ret, c.Int16.PackLE(field.Value().(int16))...)
				} else {
					ret = append(ret, c.Int16.PackBE(field.Value().(int16))...)
				}
			case CTypeInt24:
				if order == LE {
					ret = append(ret, c.Int24.PackLE(field.Value().(int32))...)
				} else {
					ret = append(ret, c.Int24.PackBE(field.Value().(int32))...)
				}
			case CTypeInt32:
				if order == LE {
					ret = append(ret, c.Int32.PackLE(field.Value().(int32))...)
				} else {
					ret = append(ret, c.Int32.PackBE(field.Value().(int32))...)
				}
			case CTypeInt40:
				if order == LE {
					ret = append(ret, c.Int40.PackLE(field.Value().(int64))...)
				} else {
					ret = append(ret, c.Int40.PackBE(field.Value().(int64))...)
				}
			case CTypeInt64:
				if order == LE {
					ret = append(ret, c.Int64.PackLE(field.Value().(int64))...)
				} else {
					ret = append(ret, c.Int64.PackBE(field.Value().(int64))...)
				}
			case CTypeUInt8:
				if order == LE {
					ret = append(ret, c.UInt8.PackLE(field.Value().(uint8))...)
				} else {
					ret = append(ret, c.UInt8.PackBE(field.Value().(uint8))...)
				}
			case CTypeUInt16:
				if order == LE {
					ret = append(ret, c.UInt16.PackLE(field.Value().(uint16))...)
				} else {
					ret = append(ret, c.UInt16.PackBE(field.Value().(uint16))...)
				}
			case CTypeUInt24:
				if order == LE {
					ret = append(ret, c.UInt24.PackLE(field.Value().(uint32))...)
				} else {
					ret = append(ret, c.UInt24.PackBE(field.Value().(uint32))...)
				}
			case CTypeUInt32:
				if order == LE {
					ret = append(ret, c.UInt32.PackLE(field.Value().(uint32))...)
				} else {
					ret = append(ret, c.UInt32.PackBE(field.Value().(uint32))...)
				}
			case CTypeUInt40:
				if order == LE {
					ret = append(ret, c.UInt40.PackLE(field.Value().(uint64))...)
				} else {
					ret = append(ret, c.UInt40.PackBE(field.Value().(uint64))...)
				}
			case CTypeUInt64:
				if order == LE {
					ret = append(ret, c.UInt64.PackLE(field.Value().(uint64))...)
				} else {
					ret = append(ret, c.UInt64.PackBE(field.Value().(uint64))...)
				}
			case CTypeFloat:
				if order == LE {
					ret = append(ret, c.Float.PackLE(field.Value().(float32))...)
				} else {
					ret = append(ret, c.Float.PackBE(field.Value().(float32))...)
				}
			case CTypeDouble:
				if order == LE {
					ret = append(ret, c.Double.PackLE(field.Value().(float64))...)
				} else {
					ret = append(ret, c.Double.PackBE(field.Value().(float64))...)
				}
			case CTypeString:
				if order == LE {
					ret = append(ret, c.String.PackLE(field.Value().(string))...)
				} else {
					ret = append(ret, c.String.PackBE(field.Value().(string))...)
				}
			case CTypeBinary:
				if order == LE {
					ret = append(ret, c.Binary.PackLE(field.Value().([]byte))...)
				} else {
					ret = append(ret, c.Binary.PackBE(field.Value().([]byte))...)
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
			case CTypeInt8:
				if order == LE {
					field.Set(c.Int8.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Int8.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Int8.Size()
			case CTypeInt16:
				if order == LE {
					field.Set(c.Int16.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Int16.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Int16.Size()
			case CTypeInt24:
				if order == LE {
					field.Set(c.Int24.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Int24.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Int24.Size()
			case CTypeInt32:
				if order == LE {
					field.Set(c.Int32.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Int32.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Int32.Size()
			case CTypeInt40:
				if order == LE {
					field.Set(c.Int40.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Int40.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Int40.Size()
			case CTypeInt64:
				if order == LE {
					field.Set(c.Int64.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Int64.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Int64.Size()
			case CTypeUInt8:
				if order == LE {
					field.Set(c.UInt8.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.UInt8.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.UInt8.Size()
			case CTypeUInt16:
				if order == LE {
					field.Set(c.UInt16.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.UInt16.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.UInt16.Size()
			case CTypeUInt24:
				if order == LE {
					field.Set(c.UInt24.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.UInt24.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.UInt24.Size()
			case CTypeUInt32:
				if order == LE {
					field.Set(c.UInt32.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.UInt32.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.UInt32.Size()
			case CTypeUInt40:
				if order == LE {
					field.Set(c.UInt40.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.UInt40.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.UInt40.Size()
			case CTypeUInt64:
				if order == LE {
					field.Set(c.UInt64.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.UInt64.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.UInt64.Size()
			case CTypeFloat:
				if order == LE {
					field.Set(c.Float.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Float.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Float.Size()
			case CTypeDouble:
				if order == LE {
					field.Set(c.Double.UnpackLE(buf[*pos:]))
				} else {
					field.Set(c.Double.UnpackBE(buf[*pos:]))
				}
				*pos = *pos + c.Double.Size()
			case CTypeString:
				if order == LE {
					field.Set(c.String.UnpackLE(buf[*pos:]))
					*pos = *pos + c.String.SizeLE(buf[*pos:])
				} else {
					field.Set(c.String.UnpackBE(buf[*pos:]))
					*pos = *pos + c.String.SizeBE(buf[*pos:])
				}
			case CTypeBinary:
				if order == LE {
					field.Set(c.Binary.UnpackLE(buf[*pos:]))
					*pos = *pos + c.Binary.SizeLE(buf[*pos:])
				} else {
					field.Set(c.Binary.UnpackBE(buf[*pos:]))
					*pos = *pos + c.Binary.SizeBE(buf[*pos:])
				}
			default:
				return errors.New(fmt.Sprintf("unknow type! type = %s", ctype))
			}
		}
	}
	return nil
}
