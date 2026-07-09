package main

import (
	"fmt"
)

// compare возвращает true, если a больше b, и a не равно 0
func compare(a, b int) bool {
	// TODO: Верните условие (a > b) И (a != 0)
	return false
}

func main() {
	fmt.Println("--- Проверка 02_comparison ---")

	if compare(5, 3) == true && compare(0, -1) == false && compare(2, 4) == false {
		fmt.Println("Тест compare: PASS")
	} else {
		fmt.Println("Тест compare: FAIL")
	}
}
