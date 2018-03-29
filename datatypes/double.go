package datatypes

import "math"

type _double struct {
}

func (this *_double) PackLE(val float64) []byte {
	u := math.Float64bits(val)
	return UInt64.PackLE(u)
}
func (this *_double) PackBE(val float64) []byte {
	u := math.Float64bits(val)
	return UInt64.PackBE(u)
}
func (this *_double) UnpackLE(buf []byte) float64 {
	u := UInt64.UnpackLE(buf)
	return math.Float64frombits(u)
}
func (this *_double) UnpackBE(buf []byte) float64 {
	u := UInt64.UnpackBE(buf)
	return math.Float64frombits(u)
}
func (this *_double) Size() int {
	return 8
}

var Double = &_double{}
