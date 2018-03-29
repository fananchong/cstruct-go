package datatypes

type t_bool struct {
}

func (this *t_bool) PackLE(val bool) []byte {
	if val {
		return []byte(1 & 0xff)
	} else {
		return []byte(0 & 0xff)
	}
}
func (this *t_bool) PackBE(val bool) []byte {
	return this.PackLE(val)
}
func (this *t_bool) UnpackLE(buf []byte) bool {
	return buf[0] != 0
}
func (this *t_bool) UnpackBE(buf []byte) bool {
	return this.UnpackBE(buf)

}
func (this *t_bool) Size() int {
	return 1
}
