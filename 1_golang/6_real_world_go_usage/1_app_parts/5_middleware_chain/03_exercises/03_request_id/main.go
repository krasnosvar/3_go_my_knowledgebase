// Упражнение 03: request ID через заголовок и context.
//
// Задача: реализуйте requestID(gen func() string) middleware:
//   - взять X-Request-ID из входящего запроса; если пусто — сгенерировать gen();
//   - положить ID в context запроса (ключ ctxKey{}) и передать дальше r.WithContext(ctx);
//   - выставить заголовок ответа X-Request-ID.
//
// Функция idFrom(ctx) уже написана — handler читает ID через неё.
//
// Запуск: go run ./03_exercises/03_request_id
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type ctxKey struct{}

func idFrom(ctx context.Context) string {
	id, _ := ctx.Value(ctxKey{}).(string)
	return id
}

func requestID(gen func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: определите id, положите в context и в заголовок ответа.
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	fmt.Println("--- Проверка 03_request_id ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var seen string
	h := requestID(func() string { return "gen-1" })(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = idFrom(r.Context())
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	check(seen == "gen-1", "без заголовка handler видит сгенерированный ID")
	check(rec.Header().Get("X-Request-ID") == "gen-1", "ID выставлен в заголовок ответа")

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Request-ID", "abc")
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	check(seen == "abc" && rec.Header().Get("X-Request-ID") == "abc", "входящий X-Request-ID сохраняется")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
