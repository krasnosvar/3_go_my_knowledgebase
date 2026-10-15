// Упражнение 01: маппинг доменных ошибок в HTTP-статусы.
//
// Задача: реализуйте statusFor(err) так, чтобы:
//   - nil → 200;
//   - ошибка, обёрнутая вокруг errNotFound → 404;
//   - вокруг errInvalid → 400;
//   - вокруг errConflict → 409;
//   - любая другая → 500.
//
// Ошибки приходят обёрнутыми через fmt.Errorf("...: %w", err), поэтому
// сравнивайте через errors.Is, а не через ==.
//
// Запуск: go run ./03_exercises/01_status_mapping
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

var _ = http.StatusOK

// statusFor переводит ошибку сервисного слоя в HTTP-статус.
func statusFor(err error) int {
	// TODO: используйте errors.Is для каждой доменной ошибки.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_status_mapping ---")

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
