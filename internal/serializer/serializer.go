package serializer

type Serializer[T any] interface {
	Encoder[T]
	Decoder[T]
}
