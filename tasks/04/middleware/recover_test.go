package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

func TestRecover(t *testing.T) {
	var router = chi.NewRouter()

	var b bytes.Buffer
	logger := log.New(&b, "", 0)

	router.Use(Recover(logger))
	router.Get("/ok", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte("OK"))
	})
	router.Get("/404", http.NotFound)
	router.Get("/panic", func(rw http.ResponseWriter, req *http.Request) {
		param := req.URL.Query().Get("m")
		panic(param)
	})

	s := httptest.NewServer(router)
	defer s.Close()

	c := s.Client()

	testCases := []struct {
		URL     string
		Code    int
		Message string
		Log     string
	}{
		{
			URL:     s.URL + "/ok",
			Code:    http.StatusOK,
			Message: "OK",
			Log:     "",
		},
		{
			URL:     s.URL + "/404",
			Code:    http.StatusNotFound,
			Message: "404 page not found\n",
			Log:     "",
		},
		{
			URL:     s.URL + "/panic?m=hello",
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error\n",
			Log:     "[ERROR] Panic caught: hello\n",
		},
		{
			URL:     s.URL + "/panic?m=world",
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error\n",
			Log:     "[ERROR] Panic caught: world\n",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b.Reset()
			code, body := fetch(t, c, tc.URL)
			require.Equal(t, tc.Code, code)
			require.Equal(t, tc.Message, body)
			require.Equal(t, tc.Log, b.String())
		})
	}
}

func fetch(t *testing.T, c *http.Client, u string) (int, string) {
	resp, err := c.Get(u)
	require.NoError(t, err)
	defer func() {
		_ = resp.Body.Close()
	}()
	data, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp.StatusCode, string(data)
}
