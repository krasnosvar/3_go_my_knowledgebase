// Упражнение 02: JSON-хендлер net/http.
//
// Задача: реализуйте userHandler(w, r):
//   - метод не GET → 405 Method Not Allowed и заголовок Allow: GET;
//   - иначе → 200, заголовок Content-Type: application/json и тело
//     {"id":1,"name":"Ann"} (сериализуйте структуру через json.NewEncoder(w)).
//
// Порядок важен: заголовки ставятся ДО WriteHeader/первой записи тела.
// Проверка — через httptest.NewRecorder, без реального сервера.
//
// Запуск: go run ./03_exercises/02_json_handler/main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = json.NewEncoder // держим импорт живым в заготовке

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// userHandler отдаёт пользователя в JSON.
func userHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: проверка r.Method, заголовки, json.NewEncoder(w).Encode(user{1, "Ann"}).
}

func main() {
	fmt.Println("--- Проверка 02_json_handler ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	rec := httptest.NewRecorder()
	userHandler(rec, httptest.NewRequest(http.MethodGet, "/user", nil))
	check(rec.Code == http.StatusOK, fmt.Sprintf("GET: код %d, ожидался 200", rec.Code))
	check(rec.Header().Get("Content-Type") == "application/json", fmt.Sprintf("Content-Type = %q", rec.Header().Get("Content-Type")))
	check(strings.TrimSpace(rec.Body.String()) == `{"id":1,"name":"Ann"}`, fmt.Sprintf("тело %q", rec.Body.String()))

	rec = httptest.NewRecorder()
	userHandler(rec, httptest.NewRequest(http.MethodPost, "/user", nil))
	check(rec.Code == http.StatusMethodNotAllowed, fmt.Sprintf("POST: код %d, ожидался 405", rec.Code))
	check(rec.Header().Get("Allow") == http.MethodGet, fmt.Sprintf("POST: заголовок Allow = %q", rec.Header().Get("Allow")))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
