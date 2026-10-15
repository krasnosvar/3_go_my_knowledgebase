// Упражнение 01: маппинг ошибок домена в HTTP-статусы.
//
// Задача: реализуйте statusFor(err), которая выбирает HTTP-код ответа:
//   - nil                → http.StatusOK (200);
//   - ErrNotFound        → http.StatusNotFound (404);
//   - ErrInvalidInput    → http.StatusBadRequest (400);
//   - ErrConflict        → http.StatusConflict (409);
//   - любая другая       → http.StatusInternalServerError (500).
//
// Ошибки могут прийти обёрнутыми (fmt.Errorf("...: %w", ErrNotFound)),
// поэтому сравнивайте через errors.Is, а не через ==.
// Используйте константы http.StatusXxx, а не «магические» числа.
//
// Запуск: go run ./03_exercises/01_status_codes/main.go
package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrConflict     = errors.New("conflict")
)

// statusFor возвращает HTTP-статус для ошибки сервисного слоя.
func statusFor(err error) int {
	// TODO: switch с errors.Is для каждой доменной ошибки.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_status_codes ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	cases := []struct {
		name string
		err  error
		want int
	}{
		{"nil", nil, http.StatusOK},
		{"not found", ErrNotFound, http.StatusNotFound},
		{"wrapped not found", fmt.Errorf("get user 42: %w", ErrNotFound), http.StatusNotFound},
		{"invalid", fmt.Errorf("validate: %w", ErrInvalidInput), http.StatusBadRequest},
		{"conflict", ErrConflict, http.StatusConflict},
		{"unknown", errors.New("db is down"), http.StatusInternalServerError},
	}
	for _, c := range cases {
		got := statusFor(c.err)
		check(got == c.want, fmt.Sprintf("%s: ожидалось %d, получено %d", c.name, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
