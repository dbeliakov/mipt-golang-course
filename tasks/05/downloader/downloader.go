package downloader

import (
	"fmt"
	"net/http"
)

type DownloadError struct {
	err error
	url string
}

func (d *DownloadError) URL() string {
	return d.url
}

func (d *DownloadError) Error() string {
	return fmt.Sprintf("download %q: %s", d.url, d.err.Error())
}

func (d *DownloadError) Unwrap() error {
	return d.err
}

// Downloader loads URLs over HTTP. MaxConcurrent limits in-flight requests (worker pool size);
// 0 means one goroutine per URL (no cap).
type Downloader struct {
	client        *http.Client
	maxConcurrent int
}

func NewDownloader(client *http.Client) *Downloader {
	return &Downloader{
		client: client,
	}
}

func NewDownloaderWithConcurrency(client *http.Client, maxConcurrent int) *Downloader {
	return &Downloader{
		client:        client,
		maxConcurrent: maxConcurrent,
	}
}

func (d *Downloader) Download(urls ...string) (map[string][]byte, []error) {
	panic("implement me")
}
