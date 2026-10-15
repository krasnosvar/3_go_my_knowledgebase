// Упражнение 01: authn-middleware по Bearer-токену.
//
// Задача: реализуйте auth(tokens map[string]string) middleware:
//   - заголовок Authorization должен иметь вид "Bearer <token>";
//   - нет заголовка / неверный формат / неизвестный токен → 401 и next НЕ вызывается;
//   - известный токен → в context кладётся userID (tokens[token]) по ключу userKey{},
//     затем вызывается next.
//
// Запуск: go run ./03_exercises/01_bearer_token
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type userKey struct{}

var _ = strings.CutPrefix
var _ = context.WithValue

func auth(tokens map[string]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: разберите заголовок, найдите пользователя, иначе 401.
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		})
	}
}

func main() {
	fmt.Println("--- Проверка 01_bearer_token ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	calls := 0
	h := auth(map[string]string{"t-1": "u-42"})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		uid, _ := r.Context().Value(userKey{}).(string)
		w.Write([]byte(uid))
	}))
	do := func(header string) *httptest.ResponseRecorder {
		req := httptest.NewRequest("GET", "/", nil)
		if header != "" {
			req.Header.Set("Authorization", header)
		}
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		return rec
	}

	rec := do("Bearer t-1")
	check(rec.Code == 200 && rec.Body.String() == "u-42", "валидный токен → 200, userID в context")
	check(do("").Code == 401, "без заголовка → 401")
	check(do("Basic t-1").Code == 401, "не Bearer → 401")
	check(do("Bearer nope").Code == 401, "неизвестный токен → 401")
	check(calls == 1, "next вызван только для валидного токена")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
