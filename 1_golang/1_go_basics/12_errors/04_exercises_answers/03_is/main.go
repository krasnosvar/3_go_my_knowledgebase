package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func isNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func main() {
	fmt.Println("--- Проверка 03_is (Решение) ---")

	err1 := fmt.Errorf("repository load failed: %w", ErrNotFound)
	err2 := fmt.Errorf("some other error")

	t1 := isNotFoundError(err1) == true
	t2 := isNotFoundError(err2) == false

	if t1 && t2 {
		fmt.Println("Тест errors.Is: PASS")
	} else {
		fmt.Printf("Тест errors.Is: FAIL (t1=%t, t2=%t)\n", t1, t2)
	}
}
