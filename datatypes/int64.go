package datatypes

type _int64 struct {
}

func (this *_int64) PackLE(val int64) []byte {
	if val < 0 {
		return UInt64.PackLE(uint64(0x7FFFFFFFFFFFFFFF - val))
	} else {
		return UInt64.PackLE(uint64(val))
	}
}
func (this *_int64) PackBE(val int64) []byte {
	if val < 0 {
		return UInt64.PackBE(uint64(0x7FFFFFFFFFFFFFFF - val))
	} else {
		return UInt64.PackBE(uint64(val))
	}
}
func (this *_int64) UnpackLE(buf []byte) int64 {
	v7 := int64(buf[7])
	if (v7 & 0x80) == 0x80 {
		return int64(0x7FFFFFFFFFFFFFFF - UInt64.UnpackLE(buf))
	} else {
		return int64(UInt64.UnpackLE(buf))
	}
}
func (this *_int64) UnpackBE(buf []byte) int64 {
	v7 := int64(buf[0])
	if (v7 & 0x80) == 0x80 {
		return int64(0x7FFFFFFFFFFFFFFF - UInt64.UnpackBE(buf))
	} else {
		return int64(UInt64.UnpackBE(buf))
	}
}
func (this *_int64) Size() int {
	return 8
}

var Int64 = &_int64{}
