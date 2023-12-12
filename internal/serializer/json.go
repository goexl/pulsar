package serializer

import (
	"encoding/json"
)

var _ Serializer[int] = (*Json[int])(nil)

type Json[T any] struct{}

func NewJson[T any]() *Json[T] {
	return new(Json[T])
}

func (j *Json[T]) Encode(from T) ([]byte, error) {
	return json.Marshal(from)
}

func (j *Json[T]) Decode(from []byte, to T) (err error) {
	return json.Unmarshal(from, to)
}
