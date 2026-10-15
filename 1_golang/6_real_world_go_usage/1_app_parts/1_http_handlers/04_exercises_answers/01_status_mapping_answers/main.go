// Упражнение 01: маппинг доменных ошибок в HTTP-статусы (Решение).
package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errNotFound = errors.New("not found")
	errInvalid  = errors.New("invalid input")
	errConflict = errors.New("conflict")
)

// statusFor переводит ошибку сервисного слоя в HTTP-статус.
func statusFor(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, errNotFound):
		return http.StatusNotFound
	case errors.Is(err, errInvalid):
		return http.StatusBadRequest
	case errors.Is(err, errConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func main() {
	fmt.Println("--- Проверка 01_status_mapping (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(statusFor(nil) == 200, "nil → 200")
	check(statusFor(fmt.Errorf("user 7: %w", errNotFound)) == 404, "обёрнутый errNotFound → 404")
	check(statusFor(fmt.Errorf("email: %w", errInvalid)) == 400, "обёрнутый errInvalid → 400")
	check(statusFor(errConflict) == 409, "errConflict → 409")
	check(statusFor(errors.New("db down")) == 500, "неизвестная ошибка → 500")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
