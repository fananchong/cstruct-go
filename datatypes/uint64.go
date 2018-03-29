package datatypes

type _uint64 struct {
}

func (this *_uint64) PackLE(val uint64) []byte {
	v0 := byte(val & 0xFF)
	v1 := byte((val >> 8) & 0xFF)
	v2 := byte((val >> 16) & 0xFF)
	v3 := byte((val >> 24) & 0xFF)
	v4 := byte((val >> 32) & 0xFF)
	v5 := byte((val >> 40) & 0xFF)
	v6 := byte((val >> 48) & 0xFF)
	v7 := byte((val >> 56) & 0xFF)
	return []byte{v0, v1, v2, v3, v4, v5, v6, v7}
}
func (this *_uint64) PackBE(val uint64) []byte {
	v7 := byte(val & 0xFF)
	v6 := byte((val >> 8) & 0xFF)
	v5 := byte((val >> 16) & 0xFF)
	v4 := byte((val >> 24) & 0xFF)
	v3 := byte((val >> 32) & 0xFF)
	v2 := byte((val >> 40) & 0xFF)
	v1 := byte((val >> 48) & 0xFF)
	v0 := byte((val >> 56) & 0xFF)
	return []byte{v0, v1, v2, v3, v4, v5, v6, v7}
}
func (this *_uint64) UnpackLE(buf []byte) uint64 {
	v0 := uint64(buf[0])
	v1 := uint64(buf[1])
	v2 := uint64(buf[2])
	v3 := uint64(buf[3])
	v4 := uint64(buf[4])
	v5 := uint64(buf[5])
	v6 := uint64(buf[6])
	v7 := uint64(buf[7])
	return (v7 << 56) + (v6 << 48) + (v5 << 40) + (v4 << 32) + (v3 << 24) + (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint64) UnpackBE(buf []byte) uint64 {
	v7 := uint64(buf[0])
	v6 := uint64(buf[1])
	v5 := uint64(buf[2])
	v4 := uint64(buf[3])
	v3 := uint64(buf[4])
	v2 := uint64(buf[5])
	v1 := uint64(buf[6])
	v0 := uint64(buf[7])
	return (v7 << 56) + (v6 << 48) + (v5 << 40) + (v4 << 32) + (v3 << 24) + (v2 << 16) + (v1 << 8) + (v0)
}
func (this *_uint64) Size() int {
	return 8
}

var UInt64 = &_uint64{}
