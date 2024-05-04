package downloader

import (
	"errors"
	"go.uber.org/goleak"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestDownloader(t *testing.T) {
	defer goleak.VerifyNone(t)

	limit := 2
	downloader := NewDownloader(limit)

	// Заменяем функцию downloadFile для контрольного тестирования
	originalDownloadFile := downloadFile
	defer func() { downloadFile = originalDownloadFile }()
	downloadFile = func(url string) (File, error) {
		if url == "http://example.com/bad" {
			return File{}, errors.New("bad request")
		}
		return File{Name: "file1", Content: []byte("content")}, nil
	}

	// Тест скачивания файлов
	urls := []string{"http://example.com/good", "http://example.com/good", "http://example.com/bad", "http://example.com/good"}

	for _, url := range urls {
		downloader.Download(url)
	}

	// Создаем слайс для учета результатов
	results := make([]DownloadResult, 0)

	// Запускаем в горутине чтение из канала
	finished := make(chan struct{})
	go func() {
		for result := range downloader.ResultsChan() {
			results = append(results, result)
		}
		close(finished)
	}()

	// Даем некоторое время на скачивание файлов
	time.Sleep(500 * time.Millisecond)

	// Вызываем shutdown и ждем его завершения и закрытия канала
	downloader.Shutdown()
	<-finished

	// Убеждаемся, что результаты соответствуют ожиданиям
	if len(results) != 4 {
		t.Errorf("Expected 4 results, got %d", len(results))
	}

	// Проверяем, что возврат ошибок работает корректно
	foundError := false
	for _, result := range results {
		if result.Error != nil {
			foundError = true
			break
		}
	}

	if !foundError {
		t.Error("Expected an error for bad request but did not receive one")
	}
}

func TestDownloaderNoRaceCondition(t *testing.T) {
	defer goleak.VerifyNone(t)

	limit := 5
	downloader := NewDownloader(limit)
	downloadFile = func(url string) (File, error) {
		time.Sleep(10 * time.Millisecond) // Имитация загрузки файла
		return File{Name: "file", Content: []byte("data")}, nil
	}

	go func() {
		for _ = range downloader.ResultsChan() {
		}
	}()

	// Скачиваем несколько файлов в горутинах
	url := "http://example.com/good"
	count := 100
	for i := 0; i < count; i++ {
		go downloader.Download(url)
	}

	// Ждем завершения всех операций
	time.Sleep(1 * time.Second)
	downloader.Shutdown()
}

func TestDownloaderConcurrencyLimit(t *testing.T) {
	defer goleak.VerifyNone(t)

	limit := 3
	downloader := NewDownloader(limit)
	var running int32

	go func() {
		for _ = range downloader.ResultsChan() {
		}
	}()

	downloadFile = func(url string) (File, error) {
		atomic.AddInt32(&running, 1)
		current := atomic.LoadInt32(&running)
		if current > int32(limit) {
			t.Errorf("Concurrency limit exceeded: %d", current)
		}
		time.Sleep(50 * time.Millisecond) // Имитация загрузки файла
		atomic.AddInt32(&running, -1)
		return File{Name: "file", Content: []byte("content")}, nil
	}

	urls := []string{"http://example.com/1", "http://example.com/2", "http://example.com/3", "http://example.com/4", "http://example.com/5"}
	for _, url := range urls {
		downloader.Download(url)
	}

	// Даем необходимое время для выполнения
	time.Sleep(time.Second)
	downloader.Shutdown()

	// Убедимся, что лимит не был превышен
	if peak := atomic.LoadInt32(&running); peak > int32(limit) {
		t.Errorf("Concurrency limit exceeded: %d", peak)
	}
}

func TestDownloaderShutdownPreventsFurtherDownloads(t *testing.T) {
	defer goleak.VerifyNone(t)

	limit := 2 // Устанавливаем лимит в 2 для наблюдения за ожиданием загрузок
	downloader := NewDownloader(limit)

	go func() {
		for _ = range downloader.ResultsChan() {
		}
	}()

	countDownloadsStarted := int32(0)
	downloadFile = func(url string) (File, error) {
		atomic.AddInt32(&countDownloadsStarted, 1)
		time.Sleep(500 * time.Millisecond) // Имитация продолжительной загрузки
		return File{Name: "file", Content: []byte("content")}, nil
	}

	// Стартуем загрузки до предела лимита
	downloader.Download("http://example.com/1")
	downloader.Download("http://example.com/2")
	downloader.Download("http://example.com/3")

	runtime.Gosched()
	time.Sleep(100 * time.Millisecond) // Даем время горутинам стартовать

	// Вызываем Shutdown
	go downloader.Shutdown()
	time.Sleep(20 * time.Millisecond) // Пауза для того чтобы Shutdown принял прерывание

	// Пытаемся начать новые загрузки после вызова Shutdown
	downloader.Download("http://example.com/3")
	downloader.Download("http://example.com/4")

	time.Sleep(150 * time.Millisecond) // Даем время для завершения всех текущих загрузок

	// Проверяем, что новые загрузки не начинались
	if count := atomic.LoadInt32(&countDownloadsStarted); count != 2 {
		t.Errorf("Expected only 2 downloads to start, but started %d", count)
	}
}
