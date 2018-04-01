package cstruct

func Unmarshal(buf []byte, obj IStruct) error {
	return NewBuffer(buf).Unmarshal(obj)
}
