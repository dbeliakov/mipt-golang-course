package httpfetch

import (
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
	defer goleak.VerifyNone(t)

	var mux = chi.NewMux()
	mux.HandleFunc("/{status}", handle)
	var s = httptest.NewServer(mux)
	defer s.Close()

	c := s.Client()
	c.Timeout = 500 * time.Millisecond
	requests := makeReqs(s.URL)
	results := FetchAll(c, requests)
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

func handle(rw http.ResponseWriter, req *http.Request) {
	status := chi.URLParam(req, "status")
	statusCode, err := strconv.Atoi(status)
	if err != nil {
		time.Sleep(1000 * time.Millisecond)
		return
	}
	rw.WriteHeader(statusCode)
}

func makeReqs(baseURL string) []Request {
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
	return res
}
