package httpfetch2

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestFetchAll(t *testing.T) {
	var mux = chi.NewMux()
	mux.HandleFunc("/{status}", handle)
	var s = httptest.NewServer(mux)

	var ctx = context.Background()

	c := s.Client()
	c.Timeout = 500 * time.Millisecond
	requests, requestsCh := makeReqs(s.URL, true)

	defer goleak.VerifyNone(t, goleak.IgnoreCurrent())
	defer s.Close()

	resultsCh := FetchAll(ctx, c, requestsCh)
	var results []Result
	for r := range resultsCh {
		results = append(results, r)
	}

	require.Len(t, results, len(requests))

	sort.Slice(results, func(i, j int) bool {
		return results[i].StatusCode < results[j].StatusCode
	})
	for _, res := range results[:3] {
		assert.Error(t, res.Error)
	}
	results = results[3:]
	assert.Equal(t, 200, results[0].StatusCode)
	assert.Equal(t, 201, results[1].StatusCode)
	assert.Equal(t, 404, results[2].StatusCode)
	assert.Equal(t, 409, results[3].StatusCode)
}

func TestFetchAll_Cancel(t *testing.T) {
	var mux = chi.NewMux()
	mux.HandleFunc("/{status}", handle)
	var s = httptest.NewServer(mux)

	var ctx, cancel = context.WithCancel(context.Background())

	c := s.Client()
	c.Timeout = 500 * time.Millisecond
	_, requestsCh := makeReqs(s.URL, false)

	defer goleak.VerifyNone(t, goleak.IgnoreCurrent())
	defer s.Close()
	defer c.CloseIdleConnections()

	resultsCh := FetchAll(ctx, c, requestsCh)
	var results []Result
	for i := 0; i < 4; i++ {
		r, ok := <-resultsCh
		require.True(t, ok)
		results = append(results, r)
	}
	cancel()

	waiter := time.After(100 * time.Millisecond)
Loop:
	for {
		select {
		case _, ok := <-resultsCh:
			if !ok {
				break Loop
			}
		case <-waiter:
			require.FailNow(t, "Channel is still not closed")
		}
	}
	<-waiter // Disable goroutine
}

func handle(rw http.ResponseWriter, req *http.Request) {
	status := chi.URLParam(req, "status")
	statusCode, err := strconv.Atoi(status)
	if err != nil {
		time.Sleep(2000 * time.Millisecond)
		return
	}
	rw.WriteHeader(statusCode)
}

func makeReqs(baseURL string, closeCh bool) ([]Request, <-chan Request) {
	var res []Request
	for _, end := range []string{"200", "201", "404", "301", "409", "invalid"} {
		res = append(res, Request{
			Method: http.MethodGet,
			URL:    fmt.Sprintf("%s/%s", baseURL, end),
		})
	}
	res = append(res, Request{
		Method: http.MethodGet,
		URL:    "a b c",
	})

	resCh := make(chan Request)
	go func() {
		for _, r := range res {
			resCh <- r
		}
		if closeCh {
			close(resCh)
		}
	}()
	return res, resCh
}
