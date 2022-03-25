package middleware

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

const count = 3

func TestMutexLimiter(t *testing.T) {
	var l = NewMutexLimiter(count)
	for i := 0; i < count; i++ {
		require.True(t, l.TryAcquire())
	}
	require.False(t, l.TryAcquire())
	l.Release()
	require.True(t, l.TryAcquire())
	require.False(t, l.TryAcquire())
}

func TestChanLimiter(t *testing.T) {
	var l = NewChanLimiter(count)
	for i := 0; i < count; i++ {
		require.True(t, l.TryAcquire())
	}
	require.False(t, l.TryAcquire())
	l.Release()
	require.True(t, l.TryAcquire())
	require.False(t, l.TryAcquire())
}

func TestLimit(t *testing.T) {
	var r = chi.NewRouter()
	r.Use(Limit(NewMutexLimiter(0)))
	r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	s := httptest.NewServer(r)
	defer s.Close()
	c := s.Client()
	resp, err := c.Get(s.URL)
	require.NoError(t, err)
	defer func() {
		_ = resp.Body.Close()
	}()
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)

	r = chi.NewRouter()
	r.Use(Limit(NewChanLimiter(1)))
	r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	s = httptest.NewServer(r)
	defer s.Close()
	c = s.Client()
	resp, err = c.Get(s.URL)
	require.NoError(t, err)
	defer func() {
		_ = resp.Body.Close()
	}()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func BenchmarkChanLimiter(b *testing.B) {
	var l = NewChanLimiter(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.TryAcquire()
		l.Release()
	}
}

func BenchmarkMutexLimiter(b *testing.B) {
	var l = NewMutexLimiter(count)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.TryAcquire()
		l.Release()
	}
}

func TestMutexLimiter_Race(t *testing.T) {
	var wg sync.WaitGroup
	var l = NewMutexLimiter(count)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				if l.TryAcquire() {
					l.Release()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestChanLimiter_Race(t *testing.T) {
	var wg sync.WaitGroup
	var l = NewChanLimiter(count)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				if l.TryAcquire() {
					l.Release()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
