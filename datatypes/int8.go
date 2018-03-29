package datatypes

type _int8 struct {
}

func (this *_int8) PackLE(val int8) []byte {
	return []byte{byte(val)}
}
func (this *_int8) PackBE(val int8) []byte {
	return this.PackLE(val)
}
func (this *_int8) UnpackLE(buf []byte) int8 {
	return int8(buf[0])
}
func (this *_int8) UnpackBE(buf []byte) int8 {
	return this.UnpackLE(buf)
}
func (this *_int8) Size() int {
	return 1
}

var Int8 = &_int8{}
