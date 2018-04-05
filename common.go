package cstruct

import "errors"

var (
	ErrNil = errors.New("cstruct: Marshal called with nil")
)

type IStruct interface {
}
