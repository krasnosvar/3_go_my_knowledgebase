package main

import (
	"fmt"
)

// runUntilSentinel симулирует бесконечный цикл
func runUntilSentinel(nextVal func() int) int {
	sum := 0
	for {
		val := nextVal()
		if val == -1 {
			break
		}
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("--- Проверка 07_infinite_loop (Решение) ---")

	counter := 0
	generator := func() int {
		counter++
		if counter > 3 {
			return -1
		}
		return counter
	}

	sum := runUntilSentinel(generator)

	if sum == 6 {
		fmt.Println("Тест бесконечного цикла: PASS")
	} else {
		fmt.Printf("Тест бесконечного цикла: FAIL (получено %d)\n", sum)
	}
}
