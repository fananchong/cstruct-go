package datatypes

type _string struct {
}

func (this *_string) PackLE(val string) []byte {
	len := len(val)
	buf1 := UInt16.PackLE(uint16(len))
	buf2 := []byte(val)
	return append(buf1, buf2...)
}
func (this *_string) PackBE(val string) []byte {
	len := len(val)
	buf1 := UInt16.PackBE(uint16(len))
	buf2 := []byte(val)
	return append(buf1, buf2...)
}
func (this *_string) UnpackLE(buf []byte) string {
	len := UInt16.UnpackLE(buf)
	return string(buf[UInt16.Size() : UInt16.Size()+int(len)])
}
func (this *_string) UnpackBE(buf []byte) string {
	len := UInt16.UnpackBE(buf)
	return string(buf[UInt16.Size() : UInt16.Size()+int(len)])
}
func (this *_string) SizeLE(buf []byte) int {
	len := UInt16.UnpackLE(buf)
	return UInt16.Size() + int(len)
}
func (this *_string) SizeBE(buf []byte) int {
	len := UInt16.UnpackBE(buf)
	return UInt16.Size() + int(len)
}

var String = &_string{}
