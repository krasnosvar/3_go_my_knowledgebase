package main

import (
	"fmt"
)

// swap меняет местами значения
func swap(a, b *int) {
	if a != nil && b != nil {
		*a, *b = *b, *a
	}
}

func main() {
	fmt.Println("--- Проверка 05_swap (Решение) ---")

	x, y := 10, 20
	swap(&x, &y)

	if x == 20 && y == 10 {
		fmt.Println("Тест swap: PASS")
	} else {
		fmt.Printf("Тест swap: FAIL (получено x=%d, y=%d)\n", x, y)
	}
}
