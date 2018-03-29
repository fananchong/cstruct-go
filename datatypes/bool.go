package datatypes

type _bool struct {
}

func (this *_bool) PackLE(val bool) []byte {
	if val {
		return []byte{1 & 0xff}
	} else {
		return []byte{0 & 0xff}
	}
}
func (this *_bool) PackBE(val bool) []byte {
	return this.PackLE(val)
}
func (this *_bool) UnpackLE(buf []byte) bool {
	return buf[0] != 0
}
func (this *_bool) UnpackBE(buf []byte) bool {
	return this.UnpackLE(buf)

}
func (this *_bool) Size() int {
	return 1
}

var Bool = &_bool{}
