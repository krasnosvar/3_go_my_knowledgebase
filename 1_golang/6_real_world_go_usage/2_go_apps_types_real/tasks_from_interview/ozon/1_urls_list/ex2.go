// Задание:
// 1. Последовательно выполнить HTTP-запросы для списка ссылок.
// Если код ответа "200 OK" -> вывести "address url - ok".
// Если код ответа отличается от "200 OK" или произошла ошибка -> вывести "address url - not ok".
//
// 2. Модифицировать программу так, чтобы запросы выполнялись в горутинах,
// а печать результатов происходила в основном потоке через канал результатов.

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	url string
	ok  bool
}

func main() {
	urls := []string{
		"http://ya.ru",
		"http://google.com",
		"http://somsite.com",
		"http://non-existent.domain.tld",
		"http://ya.ru",
		"http://ya.ru",
		"http://ya.ru",
		"http://zeeb.com",
	}

	jobs := make(chan string)
	results := make(chan result)

	var wg sync.WaitGroup
	workerPool(3, jobs, results, &wg)

	go func() {
		defer close(jobs)
		for _, url := range urls {
			jobs <- url
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		status := "not ok"
		if res.ok {
			status = "ok"
		}
		fmt.Printf("address %s - %s\n", res.url, status)
	}
}

func checkURL(url string) bool {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func workerPool(workers int, jobs <-chan string, results chan<- result, wg *sync.WaitGroup) {
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				results <- result{
					url: url,
					ok:  checkURL(url),
				}
			}
		}()
	}
}
