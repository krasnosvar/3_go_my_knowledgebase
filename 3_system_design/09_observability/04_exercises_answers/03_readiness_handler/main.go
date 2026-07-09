// Упражнение 03: readiness-хендлер, зависящий от флага готовности (Решение).
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
)

func readinessHandler(ready *atomic.Bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if ready.Load() {
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, "ready")
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = io.WriteString(w, "not ready")
	})
}

func main() {
	fmt.Println("--- Проверка 03_readiness_handler (Решение) ---")

	var ready atomic.Bool
	srv := httptest.NewServer(readinessHandler(&ready))
	defer srv.Close()

	get := func() (int, string) {
		resp, err := http.Get(srv.URL + "/readyz")
		if err != nil {
			log.Fatalf("GET: %v", err)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return resp.StatusCode, string(body)
	}

	allOK := true

	if code, body := get(); code != http.StatusServiceUnavailable || body != "not ready" {
		allOK = false
		fmt.Printf("FAIL: до готовности: статус=%d, тело=%q\n", code, body)
	}

	ready.Store(true)

	if code, body := get(); code != http.StatusOK || body != "ready" {
		allOK = false
		fmt.Printf("FAIL: после готовности: статус=%d, тело=%q\n", code, body)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
