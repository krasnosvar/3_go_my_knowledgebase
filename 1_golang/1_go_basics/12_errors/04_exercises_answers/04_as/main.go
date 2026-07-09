package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field: %s", e.Field)
}

func getValidationField(err error) (string, bool) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		return valErr.Field, true
	}
	return "", false
}

func main() {
	fmt.Println("--- Проверка 04_as (Решение) ---")

	valErr := &ValidationError{Field: "email"}
	err1 := fmt.Errorf("request processing failed: %w", valErr)
	err2 := fmt.Errorf("connection lost")

	f1, ok1 := getValidationField(err1)
	f2, ok2 := getValidationField(err2)

	t1 := ok1 && f1 == "email"
	t2 := !ok2 && f2 == ""

	if t1 && t2 {
		fmt.Println("Тест errors.As: PASS")
	} else {
		fmt.Printf("Тест errors.As: FAIL (f1=%q, ok1=%t, f2=%q, ok2=%t)\n", f1, ok1, f2, ok2)
	}
}
