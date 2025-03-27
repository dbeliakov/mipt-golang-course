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

type Downloader struct {
	client *http.Client
}

func NewDownloader(client *http.Client) *Downloader {
	return &Downloader{
		client: client,
	}
}

func (d *Downloader) Download(urls ...string) (map[string][]byte, []error) {
	panic("implement me")
}
