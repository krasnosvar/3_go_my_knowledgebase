package main

import (
	"fmt"
)

// runCalculation обновляет внешнюю переменную result без shadowing
func runCalculation(input int) int {
	result := 0

	if input > 10 {
		// Используем =, чтобы обновить внешнюю переменную, а не :=
		result = 100
	} else {
		result = 200
	}

	return result
}

func main() {
	fmt.Println("--- Проверка 06_scope (Решение) ---")

	// Тесты
	res1 := runCalculation(15)
	if res1 == 100 {
		fmt.Println("Тест input > 10: PASS")
	} else {
		fmt.Printf("Тест input > 10: FAIL (ожидалось 100, получено %d)\n", res1)
	}

	res2 := runCalculation(5)
	if res2 == 200 {
		fmt.Println("Тест input <= 10: PASS")
	} else {
		fmt.Printf("Тест input <= 10: FAIL (ожидалось 200, получено %d)\n", res2)
	}
}
