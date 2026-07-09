package main

import (
	"fmt"
)

// getZeroValues должна вернуть дефолтные нулевые значения для:
// 1. int
// 2. string
// 3. bool
// 4. *int
func getZeroValues() (int, string, bool, *int) {
	// TODO: верните нулевые значения для указанных типов
	return 1, "test", true, new(int)
}

func main() {
	fmt.Println("--- Проверка 05_zero_values ---")

	i, s, b, ptr := getZeroValues()

	if i == 0 {
		fmt.Println("Тест int: PASS")
	} else {
		fmt.Printf("Тест int: FAIL (получено %d)\n", i)
	}

	if s == "" {
		fmt.Println("Тест string: PASS")
	} else {
		fmt.Printf("Тест string: FAIL (получено %q)\n", s)
	}

	if b == false {
		fmt.Println("Тест bool: PASS")
	} else {
		fmt.Printf("Тест bool: FAIL (получено %t)\n", b)
	}

	if ptr == nil {
		fmt.Println("Тест *int: PASS")
	} else {
		fmt.Printf("Тест *int: FAIL (получено %v)\n", ptr)
	}
}
