// Упражнение 03: handler, возвращающий error (Решение).
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
	err := h(w, r)
	if err == nil {
		return
	}
	var he *httpError
	if errors.As(err, &he) {
		http.Error(w, he.Msg, he.Code)
		return
	}
	loggedErrs = append(loggedErrs, err)
	http.Error(w, "internal error", http.StatusInternalServerError)
}

func main() {
	fmt.Println("--- Проверка 03_error_handling (Решение) ---")

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
