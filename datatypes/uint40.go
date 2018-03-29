package datatypes

type _uint40 struct {
}

func (this *_uint40) PackLE(val uint64) []byte {
	v0 := byte(val & 0xFF)
	v1 := byte((val >> 8) & 0xFF)
	v2 := byte((val >> 16) & 0xFF)
	v3 := byte((val >> 24) & 0xFF)
	v4 := byte((val >> 32) & 0xFF)
	return []byte{v0, v1, v2, v3, v4}
}
func (this *_uint40) PackBE(val uint64) []byte {
	v4 := byte(val & 0xFF)
	v3 := byte((val >> 8) & 0xFF)
	v2 := byte((val >> 16) & 0xFF)
	v1 := byte((val >> 24) & 0xFF)
	v0 := byte((val >> 32) & 0xFF)
	return []byte{v0, v1, v2, v3, v4}
}
func (this *_uint40) UnpackLE(buf []byte) uint64 {
	v0 := uint64(buf[0])
	v1 := uint64(buf[1])
	v2 := uint64(buf[2])
	v3 := uint64(buf[3])
	v4 := uint64(buf[4])
	return (v4 << 32) + (v3 << 24) + (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint40) UnpackBE(buf []byte) uint64 {
	v4 := uint64(buf[0])
	v3 := uint64(buf[1])
	v2 := uint64(buf[2])
	v1 := uint64(buf[3])
	v0 := uint64(buf[4])
	return (v4 << 32) + (v3 << 24) + (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint40) Size() int {
	return 5
}

var UInt40 = &_uint40{}
