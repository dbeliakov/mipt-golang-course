package datatransfer

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestProducer struct {
	items     []any
	pos       int
	mu        sync.Mutex
	readDelay time.Duration
	failAfter int
	failWith  error
}

func NewTestProducer(items ...any) *TestProducer {
	return &TestProducer{items: items}
}

func (p *TestProducer) ReadNext() (any, error) {
	if p.readDelay > 0 {
		time.Sleep(p.readDelay)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.failAfter > 0 {
		p.failAfter--
		if p.failAfter == 0 {
			return nil, p.failWith
		}
	}

	if p.pos >= len(p.items) {
		return nil, ErrProducerFinished
	}

	item := p.items[p.pos]
	p.pos++
	return item, nil
}

func (p *TestProducer) SetReadDelay(delay time.Duration) {
	p.readDelay = delay
}

func (p *TestProducer) SetFailAfter(n int, err error) {
	p.failAfter = n
	p.failWith = err
}

type TestConsumer struct {
	chunks     [][]any
	mu         sync.Mutex
	writeDelay time.Duration
	failAfter  int
	failWith   error
}

func NewTestConsumer() *TestConsumer {
	return &TestConsumer{}
}

func (c *TestConsumer) WriteMany(items []any) error {
	if c.writeDelay > 0 {
		time.Sleep(c.writeDelay)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.failAfter > 0 {
		c.failAfter--
		if c.failAfter == 0 {
			return c.failWith
		}
	}

	chunk := make([]any, len(items))
	copy(chunk, items)
	c.chunks = append(c.chunks, chunk)
	return nil
}

func (c *TestConsumer) Chunks() [][]any {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.chunks
}

func (c *TestConsumer) Data() []any {
	c.mu.Lock()
	defer c.mu.Unlock()

	var data []any
	for _, c := range c.chunks {
		data = append(data, c...)
	}

	return data
}

func (c *TestConsumer) SetWriteDelay(delay time.Duration) {
	c.writeDelay = delay
}

func (c *TestConsumer) SetFailAfter(n int, err error) {
	c.failAfter = n
	c.failWith = err
}

func TestTransferData_SuccessfulTransfer(t *testing.T) {
	t.Parallel()
	producer := NewTestProducer(1, 2, 3, 4, 5)
	consumer := NewTestConsumer()

	transfer := NewTransfer(2, 3)
	err := transfer.TransferData(context.Background(), producer, consumer)

	require.NoError(t, err)
	assert.ElementsMatch(t, []any{
		1, 2, 3, 4, 5,
	}, consumer.Data())
}

func TestTransferData_EmptyProducer(t *testing.T) {
	t.Parallel()
	producer := NewTestProducer()
	consumer := NewTestConsumer()

	transfer := NewTransfer(2, 3)
	err := transfer.TransferData(context.Background(), producer, consumer)

	require.NoError(t, err)
	assert.Empty(t, consumer.Chunks())
}

func TestTransferData_ProducerError(t *testing.T) {
	t.Parallel()
	producerErr := errors.New("producer error")
	producer := NewTestProducer(1, 2, 3)
	producer.SetFailAfter(2, producerErr)

	consumer := NewTestConsumer()

	transfer := NewTransfer(2, 3)
	err := transfer.TransferData(context.Background(), producer, consumer)

	require.Error(t, err)
	assert.ErrorIs(t, err, producerErr)
	assert.Len(t, consumer.Chunks(), 1)
}

func TestTransferData_ConsumerError(t *testing.T) {
	t.Parallel()
	consumerErr := errors.New("consumer error")
	producer := NewTestProducer(1, 2)
	consumer := NewTestConsumer()
	consumer.SetFailAfter(1, consumerErr)

	transfer := NewTransfer(1, 3)
	err := transfer.TransferData(context.Background(), producer, consumer)

	require.Error(t, err)
	assert.ErrorIs(t, err, consumerErr)
	assert.Empty(t, consumer.Chunks())
}

func TestTransferData_ContextCancel(t *testing.T) {
	t.Parallel()
	producer := NewTestProducer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	producer.SetReadDelay(50 * time.Millisecond)

	consumer := NewTestConsumer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	transfer := NewTransfer(2, 3)
	err := transfer.TransferData(ctx, producer, consumer)

	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
}

func TestTransferData_ConcurrentWriters(t *testing.T) {
	t.Parallel()
	items := make([]any, 100)
	for i := 0; i < 100; i++ {
		items[i] = i
	}

	producer := NewTestProducer(items...)
	consumer := NewTestConsumer()

	transfer := NewTransfer(4, 10)
	err := transfer.TransferData(context.Background(), producer, consumer)

	require.NoError(t, err)

	total := len(consumer.Data())
	assert.Equal(t, 100, total)
}

func TestTransferData_ProducerFinished(t *testing.T) {
	t.Parallel()
	producer := NewTestProducer(1, 2)
	consumer := NewTestConsumer()

	transfer := NewTransfer(1, 1)
	err := transfer.TransferData(context.Background(), producer, consumer)

	require.NoError(t, err)
	assert.Equal(t, [][]any{{1}, {2}}, consumer.Chunks())
}

func TestTransferData_LargeChunkSize(t *testing.T) {
	t.Parallel()
	producer := NewTestProducer(1, 2, 3, 4, 5)
	consumer := NewTestConsumer()

	transfer := NewTransfer(2, 10)
	err := transfer.TransferData(context.Background(), producer, consumer)

	require.NoError(t, err)
	assert.ElementsMatch(t, []any{1, 2, 3, 4, 5}, consumer.Data())
}
