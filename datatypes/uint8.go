package datatypes

type _uint8 struct {
}

func (this *_uint8) PackLE(val uint8) []byte {
	return []byte{byte(val)}
}
func (this *_uint8) PackBE(val uint8) []byte {
	return this.PackLE(val)
}
func (this *_uint8) UnpackLE(buf []byte) uint8 {
	return uint8(buf[0])
}
func (this *_uint8) UnpackBE(buf []byte) uint8 {
	return this.UnpackLE(buf)
}
func (this *_uint8) Size() int {
	return 1
}

var UInt8 = &_uint8{}
