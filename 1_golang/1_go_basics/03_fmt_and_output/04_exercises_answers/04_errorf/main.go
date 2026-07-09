package main

import (
	"errors"
	"fmt"
)

var ErrTimeout = errors.New("timeout error")

// wrapTimeoutError оборачивает ошибку ErrTimeout через %w
func wrapTimeoutError() error {
	return fmt.Errorf("http request failed: %w", ErrTimeout)
}

func main() {
	fmt.Println("--- Проверка 04_errorf (Решение) ---")

	err := wrapTimeoutError()
	if err == nil {
		fmt.Println("Тест wrapTimeoutError: FAIL (ошибка равна nil)")
		return
	}

	expectedText := "http request failed: timeout error"
	if err.Error() == expectedText && errors.Is(err, ErrTimeout) {
		fmt.Println("Тест wrapTimeoutError: PASS")
	} else {
		fmt.Printf("Тест wrapTimeoutError: FAIL (получен текст: %q, errors.Is: %t)\n", err.Error(), errors.Is(err, ErrTimeout))
	}
}
