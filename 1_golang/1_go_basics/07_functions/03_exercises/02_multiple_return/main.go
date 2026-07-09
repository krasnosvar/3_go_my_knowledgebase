package main

import (
	"errors"
	"fmt"
)

var _ = errors.New // Держим импорт активным для компиляции заготовки

// divide возвращает частное (quotient) и остаток от деления (remainder).
// Если b == 0, возвращает ошибку "деление на ноль".
// Используйте именованные возвращаемые переменные в сигнатуре.
func divide(a, b int) (quotient int, remainder int, err error) {
	// TODO: Реализуйте функцию.
	return 0, 0, nil
}

func main() {
	fmt.Println("--- Проверка 02_multiple_return ---")

	q1, r1, err1 := divide(10, 3) // 3, 1, nil
	_, _, err2 := divide(10, 0)   // 0, 0, error

	t1 := q1 == 3 && r1 == 1 && err1 == nil
	t2 := err2 != nil

	if t1 && t2 {
		fmt.Println("Тест divide: PASS")
	} else {
		fmt.Printf("Тест divide: FAIL (q1=%d, r1=%d, err1=%v, err2=%v)\n", q1, r1, err1, err2)
	}
}
