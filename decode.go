package cstruct

const (
	boolPoolSize   = 16
	uint32PoolSize = 8
	uint64PoolSize = 4
)

func Unmarshal(buf []byte, obj IStruct) error {
	return NewBuffer(buf).Unmarshal(obj)
}
