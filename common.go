package cstruct

import "errors"

var (
	ErrNil = errors.New("cstruct: Marshal called with nil")
)

type IStruct interface {
}

type ByteOrder int

const (
	_  ByteOrder = iota //0
	LE                  //1
	BE                  //2
)

var CurrentByteOrder = LE
