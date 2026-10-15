// Упражнение 03: request ID через заголовок и context (Решение).
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
			id := r.Header.Get("X-Request-ID")
			if id == "" {
				id = gen()
			}
			w.Header().Set("X-Request-ID", id)
			r = r.WithContext(context.WithValue(r.Context(), ctxKey{}, id))
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	fmt.Println("--- Проверка 03_request_id (Решение) ---")

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
