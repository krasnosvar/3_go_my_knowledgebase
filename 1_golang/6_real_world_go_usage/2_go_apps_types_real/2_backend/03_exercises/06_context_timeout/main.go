// Упражнение 06: таймаут на вызов зависимости.
//
// Задача: реализуйте callWithTimeout(ctx, d, fn) (string, error):
//   - создать дочерний контекст context.WithTimeout(ctx, d) и обязательно cancel;
//   - вызвать fn(ctx) и вернуть её результат;
//
// И handler: при errors.Is(err, context.DeadlineExceeded) → 504,
// при другой ошибке → 502, иначе 200 и результат.
// fn в тесте «медленная зависимость»: ждёт либо результата, либо ctx.Done().
//
// Запуск: go run ./03_exercises/06_context_timeout
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
	// TODO: context.WithTimeout + defer cancel + fn(ctx).
	return "", nil
}

func handler(dep func(context.Context) (string, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := callWithTimeout(r.Context(), 50*time.Millisecond, dep)
		// TODO: 504 для DeadlineExceeded, 502 для прочих ошибок.
		_ = err
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
	fmt.Println("--- Проверка 06_context_timeout ---")

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
