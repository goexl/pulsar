package message

type Id interface {
	Serialize() []byte

	Ledger() int64

	Entry() int64

	Batch() int32

	Partition() int32

	Size() int32

	String() string
}
