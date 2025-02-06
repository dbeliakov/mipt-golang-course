package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	startTime := time.Now()

	dType := flag.String("type", "sync", "[sync|async]")
	dir := flag.String("dir", ".", "directory to download")
	flag.Parse()

	var (
		d   Downloader
		err error
	)
	switch *dType {
	case "sync":
		d, err = NewSyncDownloader(*dir)
		if err != nil {
			log.Fatal(err)
		}
	case "async":
		d, err = NewAsyncDownloader(*dir)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("invalid argument: type")
	}

	for _, url := range flag.Args() {
		fmt.Printf("Downloading %s\n", url)
		d.Download(url)
	}
	d.Wait()

	fmt.Printf("Completed in %s\n", time.Since(startTime))
}

type Downloader interface {
	Download(url string)
	Wait()
}

type SyncDownloader struct {
	dir        string
	httpClient *http.Client
}

var _ Downloader = (*SyncDownloader)(nil)

func NewSyncDownloader(dir string) (*SyncDownloader, error) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("create directory: %w", err)
	}
	return &SyncDownloader{
		dir:        dir,
		httpClient: &http.Client{},
	}, nil
}

func (d SyncDownloader) Download(url string) {
	err := copyContentToFile(d.httpClient, url, d.dir)
	if err != nil {
		log.Printf("download failed: %v", err)
	}
}

func (d SyncDownloader) Wait() {
	// NOTE: nothing to wait
}

type AsyncDownloader struct {
	dir        string
	httpClient *http.Client
	wg         sync.WaitGroup
}

var _ Downloader = (*AsyncDownloader)(nil)

func NewAsyncDownloader(dir string) (*AsyncDownloader, error) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("create directory: %w", err)
	}
	return &AsyncDownloader{
		dir:        dir,
		httpClient: &http.Client{},
	}, nil
}

func (d *AsyncDownloader) Download(url string) {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		err := copyContentToFile(d.httpClient, url, d.dir)
		if err != nil {
			log.Printf("download failed: %v", err)
		}
	}()
}

func (d *AsyncDownloader) Wait() {
	d.wg.Wait()
}

func copyContentToFile(client *http.Client, url string, dir string) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("get url content: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var out *os.File
	for i := 0; i < 5 && out == nil; i++ {
		fn := filepath.Join(dir, filepath.Base(url))
		if i > 0 {
			ext := filepath.Ext(url)
			fn = filepath.Join(dir, strings.TrimSuffix(filepath.Base(url), ext)+"-"+strconv.Itoa(i)+ext)
		}
		var err error
		out, err = os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				out = nil
				continue
			}
			return fmt.Errorf("create file: %w", err)
		}
		fmt.Printf("Created %s\n", fn)
	}
	if out == nil {
		return errors.New("filed to create file")
	}
	defer func() { _ = out.Close() }()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}
