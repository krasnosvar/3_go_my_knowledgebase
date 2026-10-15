// Упражнение 01: маппинг ошибок домена в HTTP-статусы (Решение).
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
	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest
	case errors.Is(err, ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func main() {
	fmt.Println("--- Проверка 01_status_codes (Решение) ---")

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
