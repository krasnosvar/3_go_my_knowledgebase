// Упражнение 06: таймаут на вызов зависимости (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func callWithTimeout(ctx context.Context, d time.Duration, fn func(context.Context) (string, error)) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	return fn(ctx)
}

func handler(dep func(context.Context) (string, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := callWithTimeout(r.Context(), 50*time.Millisecond, dep)
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			http.Error(w, "upstream timeout", http.StatusGatewayTimeout)
			return
		case err != nil:
			http.Error(w, "upstream error", http.StatusBadGateway)
			return
		}
		fmt.Fprint(w, res)
	}
}

func dep(delay time.Duration, fail bool) func(context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		select {
		case <-time.After(delay):
			if fail {
				return "", errors.New("503 from upstream")
			}
			return "data", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func main() {
	fmt.Println("--- Проверка 06_context_timeout (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	serve := func(h http.HandlerFunc) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		h(rec, httptest.NewRequest("GET", "/", nil))
		return rec
	}

	rec := serve(handler(dep(time.Millisecond, false)))
	check(rec.Code == 200 && rec.Body.String() == "data", "быстрая зависимость → 200 data")
	start := time.Now()
	rec = serve(handler(dep(2*time.Second, false)))
	check(rec.Code == 504, "медленная зависимость → 504")
	check(time.Since(start) < time.Second, "ответ не ждёт медленную зависимость")
	check(serve(handler(dep(time.Millisecond, true))).Code == 502, "ошибка зависимости → 502")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
