// Упражнение 08: liveness и readiness (Решение).
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
)

func healthMux(checks map[string]func() error) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		var failed []string
		for name, c := range checks {
			if c() != nil {
				failed = append(failed, name)
			}
		}
		if len(failed) > 0 {
			sort.Strings(failed)
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, strings.Join(failed, ","))
			return
		}
		fmt.Fprint(w, "ready")
	})
	return mux
}

func main() {
	fmt.Println("--- Проверка 08_healthcheck (Решение) ---")

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
