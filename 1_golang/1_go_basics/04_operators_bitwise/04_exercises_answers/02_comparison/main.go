package main

import (
	"fmt"
)

// compare возвращает true, если a больше b, и a не равно 0
func compare(a, b int) bool {
	return a > b && a != 0
}

func main() {
	fmt.Println("--- Проверка 02_comparison (Решение) ---")

	if compare(5, 3) == true && compare(0, -1) == false && compare(2, 4) == false {
		fmt.Println("Тест compare: PASS")
	} else {
		fmt.Println("Тест compare: FAIL")
	}
}
