// Упражнение 02: recovery middleware (Решение).
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func recovery(panics *int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					*panics++
					http.Error(w, "internal error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	fmt.Println("--- Проверка 02_recovery (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var panics int
	mw := recovery(&panics)
	boom := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { panic("boom") }))
	ok := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("fine")) }))

	run := func(h http.Handler) (rec *httptest.ResponseRecorder, crashed bool) {
		defer func() {
			if recover() != nil {
				crashed = true
			}
		}()
		rec = httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
		return rec, false
	}

	rec, crashed := run(boom)
	check(!crashed, "паника не выходит за пределы middleware")
	if !crashed {
		check(rec.Code == 500 && strings.Contains(rec.Body.String(), "internal error"), "паника → 500 internal error")
	}
	check(panics == 1, "счётчик паник = 1")
	rec, _ = run(ok)
	check(rec != nil && rec.Code == 200 && rec.Body.String() == "fine", "обычный ответ проходит без изменений")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
