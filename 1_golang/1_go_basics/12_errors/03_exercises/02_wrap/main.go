package main

import (
	"fmt"
)

// wrapError принимает ошибку err и возвращает ее обернутой
// с добавлением контекста "service failure". Используйте глагол %w.
func wrapError(err error) error {
	// TODO: Реализуйте обертывание ошибки.
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_wrap ---")

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
