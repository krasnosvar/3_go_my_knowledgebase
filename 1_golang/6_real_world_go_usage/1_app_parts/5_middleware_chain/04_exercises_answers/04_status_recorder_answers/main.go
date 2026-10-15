// Упражнение 04: middleware, который видит статус ответа (Решение).
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func metrics(counts map[int]int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(rec, r)
			counts[rec.status]++
		})
	}
}

func main() {
	fmt.Println("--- Проверка 04_status_recorder (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	counts := map[int]int{}
	h := metrics(counts)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/missing" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("ok"))
	}))
	for _, p := range []string{"/", "/", "/missing"} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("GET", p, nil))
		if p == "/missing" {
			check(rec.Code == 404, "клиент по-прежнему получает 404")
		}
	}
	check(counts[200] == 2, "два ответа 200 (статус по умолчанию)")
	check(counts[404] == 1, "один ответ 404")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
