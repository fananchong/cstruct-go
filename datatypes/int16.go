package datatypes

type _int16 struct {
}

func (this *_int16) PackLE(val int16) []byte {
	if val < 0 {
		return UInt16.PackLE(uint16(0x7FFF - val))
	} else {
		return UInt16.PackLE(uint16(val))
	}
}
func (this *_int16) PackBE(val int16) []byte {
	if val < 0 {
		return UInt16.PackBE(uint16(0x7FFF - val))
	} else {
		return UInt16.PackBE(uint16(val))
	}
}
func (this *_int16) UnpackLE(buf []byte) int16 {
	//	v1 := int16(buf[1])
	//	if (v1 & 0x80) == 0x80 {
	//		return int16(0x7FFF - UInt16.UnpackLE(buf))
	//	} else {
	return int16(UInt16.UnpackLE(buf))
	//	}
}
func (this *_int16) UnpackBE(buf []byte) int16 {
	v1 := int16(buf[0])
	if (v1 & 0x80) == 0x80 {
		return int16(0x7FFF - UInt16.UnpackBE(buf))
	} else {
		return int16(UInt16.UnpackBE(buf))
	}
}
func (this *_int16) Size() int {
	return 2
}

var Int16 = &_int16{}
