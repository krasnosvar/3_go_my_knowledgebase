package main

import (
	"fmt"
)

// runCalculation должен вернуть 100, если input > 10, и 200 в противном случае.
// Вся логика должна работать с внешней переменной result.
func runCalculation(input int) int {
	result := 0

	if input > 10 {
		// TODO: обновите значение result на 100
		// ВНИМАНИЕ: используйте оператор присвоения =, а не :=
		result := 100 // Умышленный баг (shadowing)
		_ = result
	} else {
		// TODO: обновите значение result на 200
		result := 200 // Умышленный баг (shadowing)
		_ = result
	}

	return result
}

func main() {
	fmt.Println("--- Проверка 06_scope ---")

	// Тесты
	res1 := runCalculation(15)
	if res1 == 100 {
		fmt.Println("Тест input > 10: PASS")
	} else {
		fmt.Printf("Тест input > 10: FAIL (ожидалось 100, получено %d - возможно, сработал shadowing!)\n", res1)
	}

	res2 := runCalculation(5)
	if res2 == 200 {
		fmt.Println("Тест input <= 10: PASS")
	} else {
		fmt.Printf("Тест input <= 10: FAIL (ожидалось 200, получено %d - возможно, сработал shadowing!)\n", res2)
	}
}
