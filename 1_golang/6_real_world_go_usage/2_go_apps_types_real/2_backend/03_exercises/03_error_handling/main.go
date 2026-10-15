// Упражнение 03: handler, возвращающий error.
//
// Вместо того чтобы в каждом handler'е писать http.Error, handler'ы
// возвращают error, а адаптер один раз решает, как её отдать.
//
// Задача: реализуйте (h appHandler) ServeHTTP:
//   - вызвать h(w, r); nil → ничего не делать (handler уже ответил);
//   - если ошибка (в т.ч. обёрнутая) — *httpError (errors.As) → его Code и Msg;
//   - любая другая → 500 и "internal error" (детали клиенту не раскрываем),
//     а саму ошибку добавить в loggedErrs (имитация лога).
//
// Запуск: go run ./03_exercises/03_error_handling
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type httpError struct {
	Code int
	Msg  string
}

func (e *httpError) Error() string { return e.Msg }

var loggedErrs []error

type appHandler func(http.ResponseWriter, *http.Request) error

func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: вызовите h и переведите ошибку в ответ.
	_ = h(w, r)
}

func main() {
	fmt.Println("--- Проверка 03_error_handling ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	serve := func(h appHandler) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
		return rec
	}

	rec := serve(func(w http.ResponseWriter, r *http.Request) error { w.Write([]byte("ok")); return nil })
	check(rec.Code == 200 && rec.Body.String() == "ok", "nil → ответ handler'а")
	rec = serve(func(w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("load: %w", &httpError{Code: 404, Msg: "no such order"})
	})
	check(rec.Code == 404 && strings.Contains(rec.Body.String(), "no such order"), "обёрнутый httpError → 404")
	rec = serve(func(w http.ResponseWriter, r *http.Request) error { return errors.New("pq: password=secret") })
	check(rec.Code == 500 && !strings.Contains(rec.Body.String(), "secret"), "прочая ошибка → 500 без деталей")
	check(len(loggedErrs) == 1, "внутренняя ошибка залогирована")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
