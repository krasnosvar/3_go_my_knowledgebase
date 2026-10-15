// Упражнение 03: handler создания ресурса с JSON-телом (Решение).
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type user struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func createUser(nextID *int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in user
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Email == "" {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		*nextID++
		in.ID = *nextID
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(in)
	}
}

func main() {
	fmt.Println("--- Проверка 03_json_create (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var id int
	h := createUser(&id)
	do := func(body string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		h(rec, httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body)))
		return rec
	}

	rec := do(`{"email":"a@b.c"}`)
	var got user
	_ = json.Unmarshal(rec.Body.Bytes(), &got)
	check(rec.Code == 201, "валидное тело → 201")
	check(rec.Header().Get("Content-Type") == "application/json", "Content-Type: application/json")
	check(got.ID == 1 && got.Email == "a@b.c", "тело ответа {id:1, email:a@b.c}")
	check(do(`{bad json`).Code == 400, "битый JSON → 400")
	check(do(`{"email":""}`).Code == 400, "пустой email → 400")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
