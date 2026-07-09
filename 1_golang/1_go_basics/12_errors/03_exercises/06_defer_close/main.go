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

// ProcessData читает данные из rc и закрывает его через defer.
// Если Close() возвращает ошибку, а при обработке ошибок не возникло (err == nil),
// обновите именованный возвращаемый аргумент err ошибкой из Close().
func ProcessData(rc io.ReadCloser) (err error) {
	// TODO: Реализуйте чтение и обработку закрытия в defer.
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_defer_close ---")

	closerWithErr := &MockCloser{CloseErr: errors.New("failed to close resource")}
	err := ProcessData(closerWithErr)

	t1 := err != nil && err.Error() == "failed to close resource"

	if t1 {
		fmt.Println("Тест defer close error: PASS")
	} else {
		fmt.Printf("Тест defer close error: FAIL (получена ошибка: %v)\n", err)
	}
}
