package datatypes

type _int40 struct {
}

func (this *_int40) PackLE(val int64) []byte {
	if val < 0 {
		return UInt40.PackLE(uint64(0x7FFFFFFFFF - val))
	} else {
		return UInt40.PackLE(uint64(val))
	}
}
func (this *_int40) PackBE(val int64) []byte {
	if val < 0 {
		return UInt40.PackBE(uint64(0x7FFFFFFFFF - val))
	} else {
		return UInt40.PackBE(uint64(val))
	}
}
func (this *_int40) UnpackLE(buf []byte) int64 {
	v4 := int64(buf[4])
	if (v4 & 0x80) == 0x80 {
		return int64(0x7FFFFFFFFF - UInt40.UnpackLE(buf))
	} else {
		return int64(UInt40.UnpackLE(buf))
	}
}
func (this *_int40) UnpackBE(buf []byte) int64 {
	v4 := int64(buf[0])
	if (v4 & 0x80) == 0x80 {
		return int64(0x7FFFFFFFFF - UInt40.UnpackBE(buf))
	} else {
		return int64(UInt40.UnpackBE(buf))
	}
}
func (this *_int40) Size() int {
	return 5
}

var Int40 = &_int40{}
