// Упражнение 08: health- и readiness-эндпоинты.
//
// Задача: реализуйте newMux(ready), возвращающую http.Handler с маршрутами:
//   - GET /healthz → всегда 200 и тело "ok" (liveness: процесс жив);
//   - GET /readyz  → 200 "ready", если ready() == true,
//     иначе 503 Service Unavailable и тело "not ready"
//     (readiness: готов ли принимать трафик —
//     например, прогрет кэш и есть соединение с БД);
//   - остальные пути → 404 (это поведение ServeMux по умолчанию).
//
// Запуск: go run ./03_exercises/08_health_endpoint/main.go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

// newMux собирает роутер с /healthz и /readyz.
func newMux(ready func() bool) http.Handler {
	// TODO: http.NewServeMux(), HandleFunc для /healthz и /readyz.
	return http.NewServeMux()
}

func main() {
	fmt.Println("--- Проверка 08_health_endpoint ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	isReady := false
	h := newMux(func() bool { return isReady })

	do := func(path string) (int, string) {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		return rec.Code, strings.TrimSpace(rec.Body.String())
	}

	code, body := do("/healthz")
	check(code == 200 && body == "ok", fmt.Sprintf("/healthz: %d %q", code, body))

	code, body = do("/readyz")
	check(code == 503 && body == "not ready", fmt.Sprintf("/readyz до готовности: %d %q, ожидалось 503 \"not ready\"", code, body))

	isReady = true
	code, body = do("/readyz")
	check(code == 200 && body == "ready", fmt.Sprintf("/readyz после готовности: %d %q", code, body))

	code, _ = do("/unknown")
	check(code == 404, fmt.Sprintf("/unknown: ожидался 404, получено %d", code))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
