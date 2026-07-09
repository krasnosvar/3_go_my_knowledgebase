package main

import (
	"fmt"
)

// runSafely перехватывает панику и возвращает ее как error
func runSafely(task func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Восстанавливаемся и преобразуем panic в error
			err = fmt.Errorf("%v", r)
		}
	}()
	task()
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_error_wrapper (Решение) ---")

	err1 := runSafely(func() {})
	err2 := runSafely(func() {
		panic("что-то пошло не так")
	})

	t1 := err1 == nil
	t2 := err2 != nil && err2.Error() == "что-то пошло не так"

	if t1 && t2 {
		fmt.Println("Тест runSafely: PASS")
	} else {
		fmt.Printf("Тест runSafely: FAIL (err1=%v, err2=%v)\n", err1, err2)
	}
}
