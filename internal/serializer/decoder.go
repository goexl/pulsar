package serializer

type Decoder[T any] interface {
	Decode(from []byte, to T) (err error)
}
