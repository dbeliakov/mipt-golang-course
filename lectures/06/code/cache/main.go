package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
	"https://yandex.ru",
	"https://google.com",
	"https://yandex.ru",
}

func main() {
	b := GetBodies(urls)
	for _, body := range b {
		fmt.Println(body[:100])
	}
}

type Entry struct {
	value string
	ready <-chan struct{}
}

type Cache struct {
	values map[string]*Entry
	fn     func(string) string
	m      sync.Mutex
}

func NewCache(fn func(string) string) *Cache {
	return &Cache{
		values: make(map[string]*Entry),
		fn:     fn,
	}
}

func (c *Cache) Get(key string) string {
	c.m.Lock()
	v := c.values[key]
	if v != nil {
		c.m.Unlock()
		<-v.ready
		return v.value
	}
	ready := make(chan struct{})
	c.values[key] = &Entry{ready: ready}
	c.m.Unlock()

	newV := c.fn(key)

	c.m.Lock()
	c.values[key].value = newV
	c.m.Unlock()
	close(ready)

	return newV
}

func GetBodies(urls []string) []string {
	var res []string
	var cache = NewCache(func(u string) string {
		time.Sleep(time.Second)
		fmt.Println("NEW REQUEST", u)
		resp, err := http.Get(u)
		if err != nil {
			return "error"
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "error"
		}
		return string(data)
	})

	var results = make(chan string)
	for _, u := range urls {
		go func(u string) {
			results <- cache.Get(u)
		}(u)
	}
	for i := 0; i < len(urls); i++ {
		res = append(res, <-results)
	}
	return res
}
