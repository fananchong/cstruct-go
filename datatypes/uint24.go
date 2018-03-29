package datatypes

type _uint24 struct {
}

func (this *_uint24) PackLE(val uint32) []byte {
	v0 := byte(val & 0xFF)
	v1 := byte((val >> 8) & 0xFF)
	v2 := byte((val >> 16) & 0xFF)
	return []byte{v0, v1, v2}
}
func (this *_uint24) PackBE(val uint32) []byte {
	v2 := byte(val & 0xFF)
	v1 := byte((val >> 8) & 0xFF)
	v0 := byte((val >> 16) & 0xFF)
	return []byte{v0, v1, v2}
}
func (this *_uint24) UnpackLE(buf []byte) uint32 {
	v0 := uint32(buf[0])
	v1 := uint32(buf[1])
	v2 := uint32(buf[2])
	return (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint24) UnpackBE(buf []byte) uint32 {
	v2 := uint32(buf[0])
	v1 := uint32(buf[1])
	v0 := uint32(buf[2])
	return (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint24) Size() int {
	return 3
}

var UInt24 = &_uint24{}
