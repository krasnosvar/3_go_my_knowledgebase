package main

import (
	"fmt"
)

// sum возвращает сумму двух целых чисел.
func sum(a int, b int) int {
	return a + b
}

// divide делит a на b (целочисленное деление).
// Второе значение ok сообщает, удалось ли деление: при b == 0 верните 0, false.
// (В Go для этого обычно возвращают error — это тема 12_errors; здесь достаточно bool.)
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	fmt.Println("--- Проверка 01_numeric (Решение) ---")

	// Тест 1: Сложение
	if sum(2, 3) == 5 {
		fmt.Println("Тест 1 (sum): PASS")
	} else {
		fmt.Printf("Тест 1 (sum): FAIL (ожидалось 5, получено %d)\n", sum(2, 3))
	}

	// Тест 2: Деление без ошибки
	val, ok := divide(10, 2)
	if ok && val == 5 {
		fmt.Println("Тест 2 (divide): PASS")
	} else {
		fmt.Printf("Тест 2 (divide): FAIL (ожидалось 5, true; получено %d, %t)\n", val, ok)
	}

	// Тест 3: Деление на ноль
	_, ok = divide(10, 0)
	if !ok {
		fmt.Println("Тест 3 (divide by zero): PASS")
	} else {
		fmt.Println("Тест 3 (divide by zero): FAIL (ожидалось ok == false)")
	}
}
