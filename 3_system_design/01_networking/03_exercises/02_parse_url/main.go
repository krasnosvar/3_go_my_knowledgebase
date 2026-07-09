// Упражнение 02: разбор URL на части.
//
// Задача: с помощью net/url разобрать строку URL и вернуть её части:
// схему (scheme), хост без порта (hostname), порт (port), путь (path)
// и значение query-параметра "q".
//
// Пример: "https://example.com:8443/search?q=go&lang=ru"
//
//	scheme="https", host="example.com", port="8443",
//	path="/search", query q="go"
//
// Запуск: go run ./03_exercises/02_parse_url/main.go
package main

import (
	"fmt"
	"net/url"
)

// parts описывает извлечённые из URL составляющие.
type parts struct {
	scheme string
	host   string // hostname без порта
	port   string
	path   string
	q      string // значение query-параметра "q"
}

// Держим импорт net/url живым в заготовке.
var _ = url.Parse

// parseURL разбирает raw и заполняет parts.
// При ошибке разбора возвращает (parts{}, err).
func parseURL(raw string) (parts, error) {
	// TODO: используйте url.Parse, затем u.Scheme, u.Hostname(),
	// u.Port(), u.Path и u.Query().Get("q").
	return parts{}, nil
}

func main() {
	fmt.Println("--- Проверка 02_parse_url ---")

	got, err := parseURL("https://example.com:8443/search?q=go&lang=ru")
	if err != nil {
		fmt.Printf("FAIL: неожиданная ошибка: %v\n", err)
		fmt.Println("FAIL")
		return
	}

	want := parts{
		scheme: "https",
		host:   "example.com",
		port:   "8443",
		path:   "/search",
		q:      "go",
	}

	if got == want {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: got=%+v, want=%+v\n", got, want)
		fmt.Println("FAIL")
	}
}
