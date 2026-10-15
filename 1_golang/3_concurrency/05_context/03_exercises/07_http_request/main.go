// Упражнение 07: HTTP-запрос с контекстом.
//
// Задача: fetchStatus(ctx, client, url) создаёт запрос через
// http.NewRequestWithContext(ctx, http.MethodGet, url, nil), выполняет его
// client.Do(req) и возвращает StatusCode (не забудьте закрыть resp.Body).
// Ошибки создания/выполнения запроса возвращайте как есть.
//
// Тест поднимает локальный httptest-сервер: быстрый запрос должен вернуть
// 200, а запрос с таймаутом контекста 20ms к медленному обработчику —
// ошибку context.DeadlineExceeded.
//
// Запуск: go run ./03_exercises/07_http_request/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

// fetchStatus выполняет GET с контекстом и возвращает код ответа.
func fetchStatus(ctx context.Context, client *http.Client, url string) (int, error) {
	// TODO: NewRequestWithContext, client.Do, defer resp.Body.Close().
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 07_http_request ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/slow" {
			select {
			case <-time.After(time.Second):
			case <-r.Context().Done(): // клиент ушёл — заканчиваем сразу
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	code, err := fetchStatus(context.Background(), srv.Client(), srv.URL+"/fast")
	check(err == nil && code == http.StatusOK, fmt.Sprintf("быстрый запрос: (%d, %v), ожидалось (200, nil)", code, err))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err = fetchStatus(ctx, srv.Client(), srv.URL+"/slow")
	check(errors.Is(err, context.DeadlineExceeded), fmt.Sprintf("медленный запрос: %v, ожидалось context.DeadlineExceeded", err))
	check(time.Since(start) < 500*time.Millisecond, "запрос не прервался по контексту")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
