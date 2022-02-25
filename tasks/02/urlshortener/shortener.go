package urlshortener

import (
	"net/http"
)

type URLShortener struct {
	// TODO: More fields?
	addr string
}

func NewShortener(addr string) *URLShortener {
	return &URLShortener{
		// TODO: Initialization for more fields?
		addr: addr,
	}
}

func (s *URLShortener) HandleSave(rw http.ResponseWriter, req *http.Request) {
	// TODO: Implement me
}

func (s *URLShortener) HandleExpand(rw http.ResponseWriter, req *http.Request) {
	// TODO: Implement me
}