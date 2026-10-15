// Упражнение 08: liveness и readiness.
//
// Задача: в healthMux(checks map[string]func() error) зарегистрируйте:
//   - "GET /healthz" → всегда 200 "ok" (процесс жив);
//   - "GET /readyz"  → выполнить все проверки зависимостей; если все nil →
//     200 "ready"; иначе 503 и в теле имена упавших проверок через запятую
//     в алфавитном порядке (sort.Strings).
//
// Liveness не должен зависеть от БД: иначе kubelet будет перезапускать
// под при каждой проблеме базы.
//
// Запуск: go run ./03_exercises/08_healthcheck
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
)

var _ = sort.Strings
var _ = strings.Join

func healthMux(checks map[string]func() error) *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: /healthz и /readyz.
	return mux
}

func main() {
	fmt.Println("--- Проверка 08_healthcheck ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	dbErr, cacheErr := error(nil), error(nil)
	mux := healthMux(map[string]func() error{
		"db":    func() error { return dbErr },
		"cache": func() error { return cacheErr },
	})
	get := func(p string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest("GET", p, nil))
		return rec
	}

	rec := get("/readyz")
	check(rec.Code == 200 && rec.Body.String() == "ready", "все зависимости ок → 200 ready")
	dbErr, cacheErr = errors.New("down"), errors.New("down")
	rec = get("/readyz")
	check(rec.Code == 503 && rec.Body.String() == "cache,db", "упали db и cache → 503 cache,db")
	rec = get("/healthz")
	check(rec.Code == 200 && rec.Body.String() == "ok", "liveness не зависит от БД → 200 ok")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
