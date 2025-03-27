package downloader

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockTransport struct {
	responses map[string]*http.Response
	errors    map[string]error
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err, ok := m.errors[req.URL.String()]; ok {
		return nil, err
	}
	if resp, ok := m.responses[req.URL.String()]; ok {
		return resp, nil
	}
	return nil, errors.New("unexpected URL")
}

func newMockClient(responses map[string]*http.Response, errors map[string]error) *http.Client {
	return &http.Client{
		Transport: &mockTransport{
			responses: responses,
			errors:    errors,
		},
	}
}

func TestDownloader_Download_Success(t *testing.T) {
	responses := map[string]*http.Response{
		"http://success.com": {
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("test content")),
		},
	}
	client := newMockClient(responses, nil)
	downloader := &Downloader{client: client}

	results, errs := downloader.Download("http://success.com")

	require.Empty(t, errs)
	require.Len(t, results, 1)
	assert.Equal(t, []byte("test content"), results["http://success.com"])
}

func TestDownloader_Download_NetworkError(t *testing.T) {
	mockErr := errors.New("connection failed")
	errors := map[string]error{
		"http://error.com": mockErr,
	}
	client := newMockClient(nil, errors)
	downloader := &Downloader{client: client}

	results, errs := downloader.Download("http://error.com")

	require.Empty(t, results)
	require.Len(t, errs, 1)
	assert.IsType(t, &DownloadError{}, errs[0])
	assert.Equal(t, "http://error.com", errs[0].(*DownloadError).URL())
	assert.Contains(t, errs[0].Error(), "connection failed")
}

func TestDownloader_Download_ReadError(t *testing.T) {
	responses := map[string]*http.Response{
		"http://badbody.com": {
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(&failingReader{}),
		},
	}
	client := newMockClient(responses, nil)
	downloader := &Downloader{client: client}

	results, errs := downloader.Download("http://badbody.com")

	require.Empty(t, results)
	require.Len(t, errs, 1)
	assert.IsType(t, &DownloadError{}, errs[0])
	assert.Equal(t, "http://badbody.com", errs[0].(*DownloadError).URL())
	assert.Contains(t, errs[0].Error(), "read error")
}

func TestDownloader_Download_MixedResults(t *testing.T) {
	responses := map[string]*http.Response{
		"http://good.com": {
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("good content")),
		},
	}
	errors := map[string]error{
		"http://bad.com": errors.New("failed"),
	}
	client := newMockClient(responses, errors)
	downloader := &Downloader{client: client}

	results, errs := downloader.Download("http://good.com", "http://bad.com")

	require.Len(t, results, 1)
	assert.Equal(t, []byte("good content"), results["http://good.com"])
	require.Len(t, errs, 1)
	assert.IsType(t, &DownloadError{}, errs[0])
	assert.Equal(t, "http://bad.com", errs[0].(*DownloadError).URL())
}

func TestDownloader_Download_EmptyURLs(t *testing.T) {
	client := newMockClient(nil, nil)
	downloader := &Downloader{client: client}

	results, errs := downloader.Download()

	assert.Empty(t, results)
	assert.Empty(t, errs)
}

func TestDownloader_Download_NonOKStatusCode(t *testing.T) {
	responses := map[string]*http.Response{
		"http://notfound.com": {
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(bytes.NewBufferString("not found")),
		},
	}
	client := newMockClient(responses, nil)
	downloader := &Downloader{client: client}

	results, errs := downloader.Download("http://notfound.com")

	assert.Empty(t, errs)
	require.Contains(t, results, "http://notfound.com")
	assert.Equal(t, results["http://notfound.com"], []byte("not found"))
}

type failingReader struct{}

func (f *failingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func TestDownloadError(t *testing.T) {
	mockErr := errors.New("mock error")
	de := &DownloadError{
		url: "http://test.com",
		err: mockErr,
	}

	assert.Equal(t, "http://test.com", de.URL())
	assert.Equal(t, `download "http://test.com": mock error`, de.Error())
	assert.Equal(t, mockErr, de.Unwrap())
}
