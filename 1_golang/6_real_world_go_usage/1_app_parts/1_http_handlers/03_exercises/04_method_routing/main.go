// Упражнение 04: маршрутизация по методу и path-параметру (Go 1.22+ ServeMux).
//
// Задача: в newMux зарегистрируйте маршруты стандартного http.ServeMux:
//   - "GET /items/{id}" → отвечает 200 и телом "item <id>" (r.PathValue("id"));
//   - "DELETE /items/{id}" → отвечает 204 без тела.
//
// Остальное мультиплексор сделает сам: POST на /items/1 → 405, /other → 404.
//
// Запуск: go run ./03_exercises/04_method_routing
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: зарегистрируйте "GET /items/{id}" и "DELETE /items/{id}".
	return mux
}

func main() {
	fmt.Println("--- Проверка 04_method_routing ---")

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
