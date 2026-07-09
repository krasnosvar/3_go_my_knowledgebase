// Задание: последовательно выполнить HTTP-запросы для списка ссылок.
// Если код ответа "200 OK" → вывести "address url - ok".
// Если код ответа отличается от "200 OK" или произошла ошибка → вывести "address url - not ok".

package main

import (
	"fmt"
	"net/http"
)

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

	for _, url := range urls {
		checkURL(url)
	}
}

func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("address %v - not ok\n", url)
		return
	}
	defer resp.Body.Close() // defer теперь привязан к функции checkURL, а не main

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("address %v - ok\n", url)
	} else {
		fmt.Printf("address %v - not ok\n", url)
	}
}
