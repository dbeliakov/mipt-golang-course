package datatransfer

import (
	"context"
	"errors"
)

var ErrProducerFinished = errors.New("producer finished")

type ProducerStorage interface {
	ReadNext() (any, error)
}
type ConsumerStorage interface {
	WriteMany([]any) error
}

type Transfer struct {
	writersCount int
	chunkSize    int
	// your code here
}

func NewTransfer(writersCount, chunkSize int) *Transfer {
	return &Transfer{
		writersCount: writersCount,
		chunkSize:    chunkSize,
		// your code here
	}
}

func (t *Transfer) TransferData(ctx context.Context, p ProducerStorage, c ConsumerStorage) error {
	panic("implement me")
}
