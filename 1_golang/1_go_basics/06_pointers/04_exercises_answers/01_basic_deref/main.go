package main

import (
	"fmt"
)

// getPointerValue возвращает значение по указателю или -1, если указатель nil
func getPointerValue(ptr *int) int {
	if ptr == nil {
		return -1
	}
	return *ptr
}

func main() {
	fmt.Println("--- Проверка 01_basic_deref (Решение) ---")

	x := 42
	v1 := getPointerValue(&x)
	v2 := getPointerValue(nil)

	if v1 == 42 && v2 == -1 {
		fmt.Println("Тест getPointerValue: PASS")
	} else {
		fmt.Printf("Тест getPointerValue: FAIL (v1=%d, v2=%d)\n", v1, v2)
	}
}
