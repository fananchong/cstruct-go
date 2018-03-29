package datatypes

type _binary struct {
}

func (this *_binary) PackLE(val []byte) []byte {
	len := len(val)
	buf1 := UInt16.PackLE(uint16(len))
	buf2 := val
	return append(buf1, buf2...)
}
func (this *_binary) PackBE(val []byte) []byte {
	len := len(val)
	buf1 := UInt16.PackBE(uint16(len))
	buf2 := val
	return append(buf1, buf2...)
}
func (this *_binary) UnpackLE(buf []byte) []byte {
	len := UInt16.UnpackLE(buf)
	return buf[UInt16.Size() : UInt16.Size()+int(len)]
}
func (this *_binary) UnpackBE(buf []byte) []byte {
	len := UInt16.UnpackBE(buf)
	return buf[UInt16.Size() : UInt16.Size()+int(len)]
}
func (this *_binary) SizeLE(buf []byte) int {
	len := UInt16.UnpackLE(buf)
	return UInt16.Size() + int(len)
}
func (this *_binary) SizeBE(buf []byte) int {
	len := UInt16.UnpackBE(buf)
	return UInt16.Size() + int(len)
}

var Binary = &_binary{}
