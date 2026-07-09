// Упражнение 03: readiness-хендлер, зависящий от флага готовности.
//
// Задача: реализовать readinessHandler(ready) — вернуть http.Handler,
// который отвечает:
//   - 200 OK и телом "ready", если ready.Load() == true;
//   - 503 Service Unavailable и телом "not ready", если false.
//
// Хэндлер должен читать флаг ПРИ КАЖДОМ запросе (динамически), а не
// один раз при создании. Проверка поднимает сервер через httptest,
// дёргает эндпоинт до и после переключения флага.
//
// Реализуйте readinessHandler.
//
// Запуск: go run ./03_exercises/03_readiness_handler/main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
)

// readinessHandler возвращает хендлер, чей ответ зависит от флага ready.
func readinessHandler(ready *atomic.Bool) http.Handler {
	// TODO: вернуть http.HandlerFunc, который на каждый запрос
	// читает ready.Load() и отвечает 200/"ready" либо 503/"not ready".
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func main() {
	fmt.Println("--- Проверка 03_readiness_handler ---")

	var ready atomic.Bool // на старте false
	srv := httptest.NewServer(readinessHandler(&ready))
	defer srv.Close()

	// get возвращает (статус, тело).
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

	// До готовности -> 503 "not ready".
	if code, body := get(); code != http.StatusServiceUnavailable || body != "not ready" {
		allOK = false
		fmt.Printf("FAIL: до готовности: статус=%d, тело=%q\n", code, body)
	}

	// Переключаем готовность.
	ready.Store(true)

	// После -> 200 "ready".
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
