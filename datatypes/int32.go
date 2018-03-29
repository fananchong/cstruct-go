package datatypes

type _int32 struct {
}

func (this *_int32) PackLE(val int32) []byte {
	if val < 0 {
		return UInt32.PackLE(uint32(0x7FFFFFFF - val))
	} else {
		return UInt32.PackLE(uint32(val))
	}
}
func (this *_int32) PackBE(val int32) []byte {
	if val < 0 {
		return UInt32.PackBE(uint32(0x7FFFFFFF - val))
	} else {
		return UInt32.PackBE(uint32(val))
	}
}
func (this *_int32) UnpackLE(buf []byte) int32 {
	v3 := int32(buf[3])
	if (v3 & 0x80) == 0x80 {
		return int32(0x7FFFFFFF - UInt32.UnpackLE(buf))
	} else {
		return int32(UInt32.UnpackLE(buf))
	}
}
func (this *_int32) UnpackBE(buf []byte) int32 {
	v3 := int32(buf[0])
	if (v3 & 0x80) == 0x80 {
		return int32(0x7FFFFFFF - UInt32.UnpackBE(buf))
	} else {
		return int32(UInt32.UnpackBE(buf))
	}
}
func (this *_int32) Size() int {
	return 4
}

var Int32 = &_int32{}
