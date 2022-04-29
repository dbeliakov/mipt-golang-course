package httpfetch2

import (
	"context"
	"net/http"
)

type Request struct {
	Method string
	URL    string
	Body   []byte
}

type Result struct {
	StatusCode int
	Error      error
}

func FetchAll(ctx context.Context, c *http.Client, requests <-chan Request) <-chan Result {
	сh := make(chan Result)
	close(сh)
	return сh
}
