package main

import (
	"errors"
	"fmt"
	"io"
)

type MockCloser struct {
	io.Reader
	CloseErr error
}

func (m *MockCloser) Close() error {
	return m.CloseErr
}

func ProcessData(rc io.ReadCloser) (err error) {
	defer func() {
		closeErr := rc.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	// Успешный процесс, возвращает err = nil
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_defer_close (Решение) ---")

	closerWithErr := &MockCloser{CloseErr: errors.New("failed to close resource")}
	err := ProcessData(closerWithErr)

	t1 := err != nil && err.Error() == "failed to close resource"

	if t1 {
		fmt.Println("Тест defer close error: PASS")
	} else {
		fmt.Printf("Тест defer close error: FAIL (получена ошибка: %v)\n", err)
	}
}
