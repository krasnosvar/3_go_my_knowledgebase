package main

import (
	"fmt"
)

// swap меняет местами значения, на которые указывают a и b
func swap(a, b *int) {
	// TODO: Реализуйте обмен значениями по указателям.
}

func main() {
	fmt.Println("--- Проверка 05_swap ---")

	x, y := 10, 20
	swap(&x, &y)

	if x == 20 && y == 10 {
		fmt.Println("Тест swap: PASS")
	} else {
		fmt.Printf("Тест swap: FAIL (получено x=%d, y=%d)\n", x, y)
	}
}
