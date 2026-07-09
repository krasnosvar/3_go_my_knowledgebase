package main

import (
	"fmt"
)

// getPointerValue принимает указатель и возвращает значение, на которое он указывает.
// Если указатель равен nil, вернуть -1.
func getPointerValue(ptr *int) int {
	// TODO: Реализуйте функцию разыменования с проверкой на nil.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_basic_deref ---")

	x := 42
	v1 := getPointerValue(&x)
	v2 := getPointerValue(nil)

	if v1 == 42 && v2 == -1 {
		fmt.Println("Тест getPointerValue: PASS")
	} else {
		fmt.Printf("Тест getPointerValue: FAIL (v1=%d, v2=%d)\n", v1, v2)
	}
}
