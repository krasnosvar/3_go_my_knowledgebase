package main

import (
	"fmt"
	"math"
)

// avg возвращает среднее арифметическое двух целых чисел.
// Например, avg(3, 4) должно вернуть 3.5, а не 3.0.
func avg(a int, b int) float64 {
	// TODO: реализуйте функцию с использованием явного преобразования типов
	return 0
}

func main() {
	fmt.Println("--- Проверка 05_type_conversion ---")

	// Для сравнения float используем epsilon-разницу (math.Abs)
	allPassed := true

	// Тест 1
	got := avg(3, 4)
	if math.Abs(got-3.5) < 1e-9 {
		fmt.Printf("Тест 1 (avg от %d и %d): PASS\n", 3, 4)
	} else {
		fmt.Printf("Тест 1 (avg от %d и %d): FAIL (ожидалось %.2f, получено %.2f)\n", 3, 4, 3.5, got)
		allPassed = false
	}

	// Тест 2
	got = avg(10, 20)
	if math.Abs(got-15.0) < 1e-9 {
		fmt.Printf("Тест 2 (avg от %d и %d): PASS\n", 10, 20)
	} else {
		fmt.Printf("Тест 2 (avg от %d и %d): FAIL (ожидалось %.2f, получено %.2f)\n", 10, 20, 15.0, got)
		allPassed = false
	}

	// Тест 3
	got = avg(0, 0)
	if math.Abs(got-0.0) < 1e-9 {
		fmt.Printf("Тест 3 (avg от %d и %d): PASS\n", 0, 0)
	} else {
		fmt.Printf("Тест 3 (avg от %d и %d): FAIL (ожидалось %.2f, получено %.2f)\n", 0, 0, 0.0, got)
		allPassed = false
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
