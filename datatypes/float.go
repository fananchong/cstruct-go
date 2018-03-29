package datatypes

import "math"

type _float struct {
}

func (this *_float) PackLE(val float32) []byte {
	u := math.Float32bits(val)
	return UInt32.PackLE(u)
}
func (this *_float) PackBE(val float32) []byte {
	u := math.Float32bits(val)
	return UInt32.PackBE(u)
}
func (this *_float) UnpackLE(buf []byte) float32 {
	u := UInt32.UnpackLE(buf)
	return math.Float32frombits(u)
}
func (this *_float) UnpackBE(buf []byte) float32 {
	u := UInt32.UnpackBE(buf)
	return math.Float32frombits(u)
}
func (this *_float) Size() int {
	return 4
}

var Float = &_float{}
