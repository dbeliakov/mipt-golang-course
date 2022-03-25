package middleware

import (
	"net/http"
)

func Limit(l Limiter) func(http.Handler) http.Handler {
	// TODO: Implement me
	return nil
}

type Limiter interface {
	TryAcquire() bool
	Release()
}

type MutexLimiter struct {
	// TODO: Implement me
}

func NewMutexLimiter(count int) *MutexLimiter {
	// TODO: Implement me
	return nil
}

func (l *MutexLimiter) TryAcquire() bool {
	// TODO: Implement me
	return false
}

func (l *MutexLimiter) Release() {
	// TODO: Implement me
}

type ChanLimiter struct {
	// TODO: Implement me
}

func NewChanLimiter(count int) *ChanLimiter {
	// TODO: Implement me
	return nil
}

func (l *ChanLimiter) TryAcquire() bool {
	// TODO: Implement me
	return false
}

func (l *ChanLimiter) Release() {
	// TODO: Implement me
}
