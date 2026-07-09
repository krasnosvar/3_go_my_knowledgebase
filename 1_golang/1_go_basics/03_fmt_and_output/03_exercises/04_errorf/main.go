package main

import (
	"errors"
	"fmt"
)

var ErrTimeout = errors.New("timeout error")

// wrapTimeoutError должна обернуть ошибку ErrTimeout через %w и добавить текст "http request failed: "
func wrapTimeoutError() error {
	// TODO: Верните обернутую ошибку с помощью fmt.Errorf
	return nil
}

func main() {
	fmt.Println("--- Проверка 04_errorf ---")

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
