package serializer

type Encoder[T any] interface {
	Encode(from T) (to []byte, err error)
}
