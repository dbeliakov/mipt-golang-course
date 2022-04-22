package httpfetch

import (
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

func FetchAll(c *http.Client, requests []Request) []Result {
	return nil
}
