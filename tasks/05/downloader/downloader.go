package downloader

type File struct {
	Name    string
	Content []byte
}

type DownloadResult struct {
	File  *File
	Error error
}

type Downloader struct {
}

func NewDownloader(limit int) *Downloader {
	return &Downloader{}
}

func (d *Downloader) ResultsChan() <-chan DownloadResult {
	ch := make(chan DownloadResult)
	close(ch)
	return ch
}

func (d *Downloader) Download(url string) {
}

func (d *Downloader) Shutdown() {
}

var downloadFile = func(url string) (File, error) {
	// NOTE: Для решения задачи можно не имплементировать, в тестах все равно подменяется
	return File{}, nil
}
