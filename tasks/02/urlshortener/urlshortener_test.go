package urlshortener

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

func TestURLShortener(t *testing.T) {
	var srv = NewShortener("")

	r := chi.NewMux()
	r.Put("/", srv.HandleSave)
	r.Get("/{key}", srv.HandleExpand)
	s := httptest.NewServer(r)
	// Fix for tests
	srv.addr = s.URL

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	reqY, _ := http.NewRequest(http.MethodPut, s.URL+"?u=" + url.QueryEscape("https://yandex.ru"), nil)
	reqG, _ := http.NewRequest(http.MethodPut, s.URL+"?u=" + url.QueryEscape("https://google.com"), nil)

	resp, err := client.Do(reqY)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	yLink, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	_ = resp.Body.Close()

	resp, err = client.Do(reqG)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	gLink, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	_ = resp.Body.Close()

	require.NotEqual(t, yLink, gLink)

	reqY, _ = http.NewRequest(http.MethodGet, string(yLink), nil)
	reqG, _ = http.NewRequest(http.MethodGet, string(gLink), nil)

	resp, err = client.Do(reqY)
	require.NoError(t, err)
	defer func() {
		_ = resp.Body.Close()
	}()
	require.Equal(t, http.StatusMovedPermanently, resp.StatusCode)
	require.Equal(t, resp.Header.Get("Location"), "https://yandex.ru")

	resp, err = client.Do(reqG)
	require.NoError(t, err)
	defer func() {
		_ = resp.Body.Close()
	}()
	require.Equal(t, http.StatusMovedPermanently, resp.StatusCode)
	require.Equal(t, resp.Header.Get("Location"), "https://google.com")
}

func TestURLShortener_BadRequest(t *testing.T) {
	var srv = NewShortener("")

	r := chi.NewMux()
	r.Put("/", srv.HandleSave)
	s := httptest.NewServer(r)
	// Fix for tests
	srv.addr = s.URL

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	reqY, _ := http.NewRequest(http.MethodPut, s.URL+"?u="+"/a b", nil)

	resp, err := client.Do(reqY)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	_ = resp.Body.Close()
}

func TestURLShortener_NotFound(t *testing.T) {
	var srv = NewShortener("")

	r := chi.NewMux()
	r.Put("/{key}", srv.HandleExpand)
	s := httptest.NewServer(r)
	// Fix for tests
	srv.addr = s.URL

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	reqY, _ := http.NewRequest(http.MethodPut, s.URL+"/someKey", nil)

	resp, err := client.Do(reqY)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	_ = resp.Body.Close()
}