// Упражнение 06: разбор URL на части (Решение).
package main

import (
	"fmt"
	"net/url"
)

type parts struct {
	scheme string
	host   string
	port   string
	path   string
	q      string
}

func parseURL(raw string) (parts, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return parts{}, err
	}
	return parts{
		scheme: u.Scheme,
		host:   u.Hostname(),
		port:   u.Port(),
		path:   u.Path,
		q:      u.Query().Get("q"),
	}, nil
}

func main() {
	fmt.Println("--- Проверка 06_parse_url (Решение) ---")

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
