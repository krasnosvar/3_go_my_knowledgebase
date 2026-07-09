package main

import (
	"errors"
	"fmt"
)

// sum возвращает сумму двух целых чисел.
var _ = errors.New

func sum(a int, b int) int {
	// TODO: реализуйте функцию
	return 0
}

// divide делит a на b. Если b == 0, возвращает ошибку.
func divide(a, b int) (int, error) {
	// TODO: реализуйте функцию, учитывая деление на ноль
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 01_numeric ---")

	// Тест 1: Сложение
	if sum(2, 3) == 5 {
		fmt.Println("Тест 1 (sum): PASS")
	} else {
		fmt.Printf("Тест 1 (sum): FAIL (ожидалось 5, получено %d)\n", sum(2, 3))
	}

	// Тест 2: Деление без ошибки
	val, err := divide(10, 2)
	if err == nil && val == 5 {
		fmt.Println("Тест 2 (divide): PASS")
	} else {
		fmt.Printf("Тест 2 (divide): FAIL (ожидалось 5, nil; получено %d, %v)\n", val, err)
	}

	// Тест 3: Деление на ноль
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Тест 3 (divide by zero): PASS")
	} else {
		fmt.Println("Тест 3 (divide by zero): FAIL (ожидалась ошибка)")
	}
}
