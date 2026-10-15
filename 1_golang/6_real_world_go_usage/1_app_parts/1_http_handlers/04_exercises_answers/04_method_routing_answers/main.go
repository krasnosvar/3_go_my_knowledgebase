// Упражнение 04: маршрутизация по методу и path-параметру (Решение).
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item %s", r.PathValue("id"))
	})
	mux.HandleFunc("DELETE /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	return mux
}

func main() {
	fmt.Println("--- Проверка 04_method_routing (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	mux := newMux()
	do := func(method, path string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(method, path, nil))
		return rec
	}

	rec := do(http.MethodGet, "/items/42")
	check(rec.Code == 200 && rec.Body.String() == "item 42", "GET /items/42 → 200 \"item 42\"")
	check(do(http.MethodDelete, "/items/42").Code == 204, "DELETE /items/42 → 204")
	check(do(http.MethodPost, "/items/42").Code == 405, "POST /items/42 → 405")
	check(do(http.MethodGet, "/other").Code == 404, "GET /other → 404")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
