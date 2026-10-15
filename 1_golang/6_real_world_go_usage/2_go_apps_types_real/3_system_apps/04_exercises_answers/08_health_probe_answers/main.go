// Упражнение 08: HTTP health probe с таймаутом (Решение).
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func probe(url string, timeout time.Duration) bool {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func main() {
	fmt.Println("--- Проверка 08_health_probe (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			w.WriteHeader(204)
		case "/fail":
			w.WriteHeader(503)
		case "/slow":
			select {
			case <-time.After(2 * time.Second):
			case <-r.Context().Done():
			}
		}
	}))
	defer srv.Close()

	check(probe(srv.URL+"/ok", time.Second), "204 → здоров")
	check(!probe(srv.URL+"/fail", time.Second), "503 → нездоров")
	start := time.Now()
	check(!probe(srv.URL+"/slow", 100*time.Millisecond), "медленный ответ → нездоров")
	check(time.Since(start) < time.Second, "таймаут клиента сработал")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
