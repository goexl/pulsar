package serializer

var (
	_ Serializer[int] = (*Json[int])(nil)
)

type Json[T any] struct{}

func NewJson[T any]() *Json[T] {
	return new(Json[T])
}

func (j *Json[T]) Encode(from T) (to []byte, err error) {
	return
}

func (j *Json[T]) Decode(from []byte, to T) (err error) {
	return
}
