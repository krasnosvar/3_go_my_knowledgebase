package main

import (
	"fmt"
)

func wrapError(err error) error {
	return fmt.Errorf("service failure: %w", err)
}

func main() {
	fmt.Println("--- Проверка 02_wrap (Решение) ---")

	orig := fmt.Errorf("database query error")
	wrapped := wrapError(orig)

	t1 := wrapped != nil
	t2 := false
	if wrapped != nil {
		t2 = wrapped.Error() == "service failure: database query error"
	}

	if t1 && t2 {
		fmt.Println("Тест wrapError: PASS")
	} else {
		fmt.Printf("Тест wrapError: FAIL (wrapped=%v)\n", wrapped)
	}
}
