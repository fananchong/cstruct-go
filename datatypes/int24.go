package datatypes

type _int24 struct {
}

func (this *_int24) PackLE(val int32) []byte {
	if val < 0 {
		return UInt24.PackLE(uint32(0x7FFFFF - val))
	} else {
		return UInt24.PackLE(uint32(val))
	}
}
func (this *_int24) PackBE(val int32) []byte {
	if val < 0 {
		return UInt24.PackBE(uint32(0x7FFFFF - val))
	} else {
		return UInt24.PackBE(uint32(val))
	}
}
func (this *_int24) UnpackLE(buf []byte) int32 {
	v2 := int32(buf[2])
	if (v2 & 0x80) == 0x80 {
		return int32(0x7FFFFF - UInt24.UnpackLE(buf))
	} else {
		return int32(UInt24.UnpackLE(buf))
	}
}
func (this *_int24) UnpackBE(buf []byte) int32 {
	v2 := int32(buf[0])
	if (v2 & 0x80) == 0x80 {
		return int32(0x7FFFFF - UInt24.UnpackBE(buf))
	} else {
		return int32(UInt24.UnpackBE(buf))
	}
}
func (this *_int24) Size() int {
	return 3
}

var Int24 = &_int24{}
