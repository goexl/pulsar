package core

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/goexl/gox"
)

type Pid struct {
	id pulsar.MessageID
	_  gox.CannotCopy
}

func NewPid(id pulsar.MessageID) *Pid {
	return &Pid{
		id: id,
	}
}

func (p *Pid) Serialize() []byte {
	return p.id.Serialize()
}

func (p *Pid) Ledger() int64 {
	return p.id.LedgerID()
}

func (p *Pid) Entry() int64 {
	return p.id.EntryID()
}

func (p *Pid) Batch() int32 {
	return p.id.BatchIdx()
}

func (p *Pid) Partition() int32 {
	return p.id.PartitionIdx()
}

func (p *Pid) Size() int32 {
	return p.id.BatchSize()
}

func (p *Pid) String() string {
	return p.id.String()
}
