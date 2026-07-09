package main

import (
	"fmt"
)

// runUntilSentinel симулирует бесконечный цикл, который получает числа из генератора nextVal.
// Цикл должен суммировать числа. Если nextVal возвращает -1 (sentinel value),
// цикл должен немедленно завершиться через break, а накопленная сумма — возвратиться.
func runUntilSentinel(nextVal func() int) int {
	sum := 0

	// TODO: Реализуйте бесконечный цикл с проверкой на возвращаемое значение -1.

	return sum
}

func main() {
	fmt.Println("--- Проверка 07_infinite_loop ---")

	// Тестовый генератор, возвращающий 1, 2, 3, затем -1
	counter := 0
	generator := func() int {
		counter++
		if counter > 3 {
			return -1
		}
		return counter
	}

	sum := runUntilSentinel(generator)

	if sum == 6 { // 1 + 2 + 3 = 6
		fmt.Println("Тест бесконечного цикла: PASS")
	} else {
		fmt.Printf("Тест бесконечного цикла: FAIL (получено %d)\n", sum)
	}
}
