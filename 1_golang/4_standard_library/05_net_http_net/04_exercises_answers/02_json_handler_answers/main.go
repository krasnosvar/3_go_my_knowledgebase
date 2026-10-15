// Упражнение 02: JSON-хендлер net/http (Решение).
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// userHandler отдаёт пользователя в JSON.
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user{ID: 1, Name: "Ann"}); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("--- Проверка 02_json_handler (Решение) ---")

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
