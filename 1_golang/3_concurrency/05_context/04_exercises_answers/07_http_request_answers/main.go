// Упражнение 07: HTTP-запрос с контекстом (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func fetchStatus(ctx context.Context, client *http.Client, url string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func main() {
	fmt.Println("--- Проверка 07_http_request (Решение) ---")

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
