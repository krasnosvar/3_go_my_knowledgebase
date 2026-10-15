// Упражнение 04: цепочка logging + auth с прерыванием.
//
// Задача:
//   - logging(log *[]string): после next записать в log строку "<METHOD> <path>";
//   - auth: если заголовок X-API-Key != "secret" → 401 и next НЕ вызывать;
//   - в main цепочка собирается как logging(auth(h)), поэтому даже
//     отклонённые запросы попадают в лог.
//
// Запуск: go run ./03_exercises/04_middleware_chain
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func logging(log *[]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			// TODO: допишите запись "<METHOD> <path>" в *log.
		})
	}
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: проверьте X-API-Key, при ошибке 401 и return.
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("--- Проверка 04_middleware_chain ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var log []string
	hits := 0
	h := logging(&log)(auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { hits++ })))

	req := httptest.NewRequest("GET", "/orders", nil)
	req.Header.Set("X-API-Key", "secret")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	check(rec.Code == 200 && hits == 1, "с ключом запрос доходит до handler'а")

	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("DELETE", "/orders", nil))
	check(rec.Code == 401 && hits == 1, "без ключа → 401, handler не вызван")
	check(fmt.Sprint(log) == "[GET /orders DELETE /orders]", "оба запроса в логе: "+fmt.Sprint(log))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
