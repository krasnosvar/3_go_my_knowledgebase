// Упражнение 08: HTTP health probe с таймаутом.
//
// Задача: реализуйте probe(url string, timeout time.Duration) bool:
//   - http.Client{Timeout: timeout} (у http.DefaultClient таймаута нет!);
//   - GET url; true только если нет ошибки и статус 2xx;
//   - тело ответа закрыть (defer resp.Body.Close()).
//
// Тест поднимает локальный httptest.Server (без внешней сети).
//
// Запуск: go run ./03_exercises/08_health_probe
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func probe(url string, timeout time.Duration) bool {
	// TODO: клиент с Timeout, GET, проверка статуса 2xx.
	return false
}

func main() {
	fmt.Println("--- Проверка 08_health_probe ---")

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
