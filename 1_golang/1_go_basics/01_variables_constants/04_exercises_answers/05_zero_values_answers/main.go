package main

import (
	"fmt"
)

// getZeroValues возвращает дефолтные нулевые значения
func getZeroValues() (int, string, bool, *int) {
	var i int
	var s string
	var b bool
	var ptr *int
	return i, s, b, ptr
}

func main() {
	fmt.Println("--- Проверка 05_zero_values (Решение) ---")

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
