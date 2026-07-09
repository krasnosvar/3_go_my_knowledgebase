package main

import (
	"fmt"
)

// calculate возвращает сумму, разность, произведение и частное двух чисел
func calculate(a, b int) (int, int, int, int) {
	// TODO: Верните a+b, a-b, a*b, a/b.
	return 0, 0, 0, 0
}

func main() {
	fmt.Println("--- Проверка 01_arithmetic ---")

	sum, diff, prod, div := calculate(10, 2)

	if sum == 12 && diff == 8 && prod == 20 && div == 5 {
		fmt.Println("Тест calculate: PASS")
	} else {
		fmt.Printf("Тест calculate: FAIL (получено: sum=%d, diff=%d, prod=%d, div=%d)\n", sum, diff, prod, div)
	}
}
