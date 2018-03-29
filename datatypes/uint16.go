package datatypes

type _uint16 struct {
}

func (this *_uint16) PackLE(val uint16) []byte {
	v0 := byte(val & 0xFF)
	v1 := byte((val >> 8) & 0xFF)
	return []byte{v0, v1}
}
func (this *_uint16) PackBE(val uint16) []byte {
	v1 := byte(val & 0xFF)
	v0 := byte((val >> 8) & 0xFF)
	return []byte{v0, v1}
}
func (this *_uint16) UnpackLE(buf []byte) uint16 {
	v0 := uint16(buf[0])
	v1 := uint16(buf[1])
	return (v1 << 8) + (v0)
}
func (this *_uint16) UnpackBE(buf []byte) uint16 {
	v1 := uint16(buf[0])
	v0 := uint16(buf[1])
	return (v1 << 8) + (v0)
}
func (this *_uint16) Size() int {
	return 2
}

var UInt16 = &_uint16{}
