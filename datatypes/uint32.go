package datatypes

type _uint32 struct {
}

func (this *_uint32) PackLE(val uint32) []byte {
	v0 := byte(val & 0xFF)
	v1 := byte((val >> 8) & 0xFF)
	v2 := byte((val >> 16) & 0xFF)
	v3 := byte((val >> 24) & 0xFF)
	return []byte{v0, v1, v2, v3}
}
func (this *_uint32) PackBE(val uint32) []byte {
	v3 := byte(val & 0xFF)
	v2 := byte((val >> 8) & 0xFF)
	v1 := byte((val >> 16) & 0xFF)
	v0 := byte((val >> 24) & 0xFF)
	return []byte{v0, v1, v2, v3}
}
func (this *_uint32) UnpackLE(buf []byte) uint32 {
	v0 := uint32(buf[0])
	v1 := uint32(buf[1])
	v2 := uint32(buf[2])
	v3 := uint32(buf[3])
	return (v3 << 24) + (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint32) UnpackBE(buf []byte) uint32 {
	v3 := uint32(buf[0])
	v2 := uint32(buf[1])
	v1 := uint32(buf[2])
	v0 := uint32(buf[3])
	return (v3 << 24) + (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint32) Size() int {
	return 4
}

var UInt32 = &_uint32{}
